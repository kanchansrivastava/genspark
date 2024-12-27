package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"strings"
	"user-service/pkg/logkey"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		// We get the current request context
		ctx := c.Request.Context()

		// Extract the traceId from the request context
		// We assert the type to string since context.Value returns an interface{}
		traceId, ok := ctx.Value(TraceIdKey).(string)

		if !ok {
			traceId = "unknown"
		}
		authHeader := c.Request.Header.Get("Authorization")

		// Splitting the Authorization header based on the space character.
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			// If the header format doesn't match required format, log and send an error
			err := errors.New("expected authorization header format: Bearer <token>")
			slog.Error("An error occurred",
				slog.Any(logkey.ERROR, err),
				slog.Any(logkey.TraceID, traceId),
			)
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		// Call the validate token from auth struct
		//put the validated claims in context
		// do the next thing in the chain

	}
}
