package handlers

import (
	"book-store/models"
	"book-store/pkg/ctxmanage"
	"github.com/gin-gonic/gin"
	"log/slog"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func (h *Handler) CreateBook(c *gin.Context) {
	traceId := ctxmanage.GetTraceID(c.Request.Context())

	var nb models.NewBook
	err := c.ShouldBindJSON(&nb)
	if err != nil {
		//log.Println(err)
		slog.Error(
			"Invalid JSON request",
			slog.String("TRACE ID", traceId),
			slog.String("Error", err.Error()),
		)
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	err = h.validate.Struct(nb)
	if err != nil {
		//log.Println(err)
		slog.Error(
			"Validation failed",
			slog.String("TRACE ID", traceId),
			slog.String("Error", err.Error()),
		)

		c.JSON(400, gin.H{"error": "please provide values in correct format"})
		return
	}
	book, err := h.c.InsertBook(c.Request.Context(), nb)
	if err != nil {
		//log.Println(err)
		slog.Error(
			"Error inserting book",
			slog.String("TRACE ID", traceId),
			slog.String("Error", err.Error()),
		)
		c.JSON(500, gin.H{"error": "problem inserting book"})
		return
	}

	slog.Info("Book inserted", slog.String("TRACE ID", traceId))
	c.JSON(201, book)
}
