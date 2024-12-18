package main

import (
	"context"
	"errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	"log"

	// Importing OpenTelemetry packages for HTTP instrumentation and tracing
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.26.0"
	trace2 "go.opentelemetry.io/otel/trace"
	"net/http"
)

var users = map[string]string{
	"1": "Alice",
	"2": "Bob",
	"3": "Charlie",
}

func main() {
	mux := http.NewServeMux()
	traceProvider, err := initOpenTelemetry()
	if err != nil {
		panic(err)
	}
	defer traceProvider.Shutdown(context.Background())

	wrappedMux := otelMid(mux)
	mux.HandleFunc("/user", GetUser)

	panic(http.ListenAndServe(":8086", wrappedMux))
}

func initOpenTelemetry() (*trace.TracerProvider, error) {
	traceExporter, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithInsecure(),
		// The endpoint where Jaeger or another collector is running
		otlptracehttp.WithEndpoint("localhost:4318"))

	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(
		trace.WithSampler(trace.AlwaysSample()),
		//0.0 would mean sampling 0% of requests (never sampling).
		//0.5 would mean sampling 50% of requests (half of them).
		//1.0 would mean sampling 100% of requests (all of them).
		//
		//trace.WithSampler(trace.TraceIDRatioBased(0.1)), // 10% of traces
		trace.WithBatcher(traceExporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("user-micro"),
		)),
	)
	// Set the global TracerProvider for the application
	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.TraceContext{})
	//TODO: for distributed tracing add TEXT MAP Propagator
	return traceProvider, nil
}

func GetUser(w http.ResponseWriter, r *http.Request) {

	tracer := otel.Tracer("user-micro")

	ctx, span := tracer.Start(r.Context(), "GetUser")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()

	userId := r.URL.Query().Get("id") // Get the user ID from the query parameters
	user, err := GetUserById(ctx, userId, tracer)
	if err != nil {
		span.SetAttributes(semconv.HTTPResponseStatusCodeKey.Int(http.StatusBadRequest))
		span.SetAttributes(attribute.String("user_id", userId))
		span.SetAttributes(attribute.String("traceId", traceId))
		span.SetStatus(codes.Error, err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	span.SetAttributes(semconv.HTTPResponseStatusCodeKey.Int(http.StatusOK))
	w.Write([]byte(user))
}

func otelMid(next http.Handler) http.Handler {
	return otelhttp.NewHandler(next, "user-micro")
}

func GetUserById(ctx context.Context, userId string, tracer trace2.Tracer) (string, error) {
	_, span := tracer.Start(ctx, "GetUserById")
	defer span.End()

	traceId := span.SpanContext().TraceID().String()
	log.Printf("traceId: %s hitting db\n", traceId)

	if user, exists := users[userId]; exists {
		span.SetStatus(codes.Ok, "user found in database")
		return user, nil
	}
	span.SetStatus(codes.Error, "user not found in database")

	return "", errors.New("user not found")
}
