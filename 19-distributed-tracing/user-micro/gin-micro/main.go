package main

import (
	"context"
	"errors"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	trace2 "go.opentelemetry.io/otel/trace"
)

// In-memory mock database for user data
var users = map[string]string{
	"1": "Alice",
	"2": "Bob",
	"3": "Charlie",
}

func main() {
	// Step 1: Initialize OpenTelemetry
	traceProvider, err := initOpenTelemetry()
	if err != nil {
		panic(err)
	}
	defer traceProvider.Shutdown(context.Background())

	// Step 2: Create a Gin router
	r := gin.Default()

	// Step 3: Add OpenTelemetry middleware to Gin
	// This will automatically trace all incoming HTTP requests handled by Gin.
	r.Use(otelgin.Middleware("user-micro"))

	// Step 4: Define the `/user` endpoint
	r.GET("/user/:id", GetUser)
	r.GET("/call-order-service", CallOrderService)

	// Step 5: Start the server on port 8086
	if err := r.Run(":8086"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
func CallOrderService(c *gin.Context) {
	// Step 1: Start a new distributed tracing span using OpenTelemetry.
	// This span will track the operation of calling the order service,
	ctx, span := otel.Tracer("user-micro").Start(c.Request.Context(), "CallOrderService")
	defer span.End() // Ensures that the span ends when the function is done (clean-up step).

	// Step 2: Create a new HTTP client that uses OpenTelemetry instrumentation.
	// The `otelhttp.NewTransport` wraps the default HTTP transport with tracing instrumentation.
	// This allows OpenTelemetry to automatically create spans for outgoing HTTP calls.
	client := &http.Client{Transport: otelhttp.NewTransport(http.DefaultTransport)}

	// Step 3: Create an HTTP GET request to the external service (Order Service) endpoint.
	// The context `ctx` contains tracing metadata that will be sent with the request.
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:8089/order", nil)
	if err != nil {
		// If constructing the HTTP request fails (e.g., incorrect URL or method), log the error
		// and respond to the client with a 500 Internal Server Error.
		log.Printf("Failed to construct request for the order service: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 4: Inject trace context metadata into the HTTP request headers.
	// OpenTelemetry uses this to propagate trace data to the downstream Order Service.
	// This ensures distributed tracing works seamlessly across multiple services.
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(req.Header))

	// Step 5: Execute the HTTP request to the Order Service using the instrumented client.
	resp, err := client.Do(req)
	if err != nil {
		// If there's an error communicating with the Order Service (e.g., server down or network issue),
		// log the error and respond to the client with a 500 Internal Server Error.
		log.Printf("Failed to call the order service: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 6: Read the response body returned by the Order Service.
	// The `io.ReadAll` function reads the entire body into memory as a byte slice.
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		// If reading the response body fails, log the error and respond with a 500 error to the client.
		log.Printf("Failed to read response from order service: %v", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Step 7: Set the status of the tracing span to "Ok" to indicate that
	// the request to the Order Service was successful and the response was processed correctly.
	span.SetStatus(codes.Ok, "order service response received")

	// Step 8: Log the response received from the external service for debugging or monitoring purposes.
	// This provides visibility into what the downstream service returned.
	log.Printf("Order service response: %s", string(b))

	// Step 9: Send the response from the Order Service back to the client.
	// Respond with HTTP status 200 (OK) if everything is successful.
	// Use `c.String` to return the response as a plain-text string.
	c.String(http.StatusOK, string(b))
}

// Initialize OpenTelemetry for distributed tracing
func initOpenTelemetry() (*trace.TracerProvider, error) {
	// Set up the OTLP trace exporter to send tracing data to the OpenTelemetry Collector
	traceExporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithInsecure(),                 // No TLS for local development
		otlptracehttp.WithEndpoint("localhost:4318"), // Collector/Jaeger endpoint
	)
	if err != nil {
		return nil, err
	}

	// Configure a TracerProvider
	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()), // Sample all traces
		trace.WithBatcher(traceExporter),        // Batch traces in export
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("user-micro"), // Set the service name for tracing
		)),
	)

	// Register the global TracerProvider and propagators
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	return traceProvider, nil
}

// Handler for the `/user` endpoint
// Retrieves a user by ID based on the query parameter `id`
func GetUser(c *gin.Context) {
	// Step 6: Start a new span for the handler
	tracer := otel.Tracer("user-micro")
	ctx, span := tracer.Start(c.Request.Context(), "GetUser")
	defer span.End()

	// Extract the user ID from the query parameters
	userId := c.Param("id")
	traceId := span.SpanContext().TraceID().String()

	// Retrieve user data from the mock database
	user, err := GetUserById(ctx, userId, tracer)
	if err != nil {
		// Handle and record any errors in the span
		span.SetAttributes(semconv.HTTPResponseStatusCodeKey.Int(http.StatusBadRequest)) // HTTP 400
		span.SetAttributes(attribute.String("user_id", userId))                          // Attach user ID
		span.SetAttributes(attribute.String("traceId", traceId))
		span.AddEvent("USER NOT FOUND")                            // Record event in tracing span// Attach trace ID
		span.SetStatus(codes.Error, err.Error())                   // Mark span as error
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}) // Respond with 400 error
		return
	}

	// Mark the span as successful
	span.SetAttributes(semconv.HTTPResponseStatusCodeKey.Int(http.StatusOK)) // HTTP 200
	c.String(http.StatusOK, user)                                            // Respond with the user
}

// Simulate retrieving a user by ID (mock database operation)
func GetUserById(ctx context.Context, userId string, tracer trace2.Tracer) (string, error) {
	// Step 7: Create a child span for this database operation
	_, span := tracer.Start(ctx, "GetUserById")
	defer span.End()

	// Log the trace ID for debugging
	traceId := span.SpanContext().TraceID().String()
	log.Printf("traceID: %s - fetching user from database", traceId)

	// Check the in-memory database for the user ID
	if user, exists := users[userId]; exists {
		// Record success in the span
		span.SetStatus(codes.Ok, "user found")
		return user, nil
	}

	// Record failure in the span
	span.SetStatus(codes.Error, "user not found")
	return "", errors.New("user not found")
}
