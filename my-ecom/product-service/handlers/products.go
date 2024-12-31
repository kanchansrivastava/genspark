package handlers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log/slog"
	"net/http"
	"product-service/internal/products"
	"product-service/pkg/ctxmanage"
	"product-service/pkg/logkey"
	"time"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	fmt.Println("CreateProduct Handler called!!")
	traceId := ctxmanage.GetTraceIdOfRequest(c)

	if c.Request.ContentLength > 5*1024 {
		slog.Error("request body limit breached",
			slog.String(logkey.TraceID, traceId),
			slog.Int64("Size Received", c.Request.ContentLength),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": "payload exceeding size limit",
		})
		return
	}

	var newProduct products.NewProduct
	if err := c.ShouldBindJSON(&newProduct); err != nil {
		slog.Error("json validation error",
			slog.String(logkey.TraceID, traceId),
			slog.String(logkey.ERROR, err.Error()),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid JSON: %s", err.Error()),
		})
		return
	}

	if err := h.validate.Struct(newProduct); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make([]string, len(validationErrors))
		for i, e := range validationErrors {
			errorMessages[i] = fmt.Sprintf("%s %s", e.Field(), e.Tag())
		}
		slog.Error("validation failed",
			slog.String(logkey.TraceID, traceId),
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorMessages,
		})
		return
	}

	paisaPrice, err := products.ValidatePrice(newProduct.Price)
	if err != nil {
		slog.Error("price validation failed",
			slog.String(logkey.TraceID, traceId),
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Invalid price: %v", err),
		})
		return
	}

	ctx := c.Request.Context()

	// Insert product into the database
	product, err := h.p.InsertProduct(ctx, newProduct)
	if err != nil {
		slog.Error("error in creating the product",
			slog.String(logkey.TraceID, traceId),
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": fmt.Sprintf("Product Creation Failed: %s", err.Error()),
		})
		return
	}

	// Handle product creation on Stripe asynchronously
	go func() {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		err = h.p.CreateProductStripe(ctx, product.ID, product.Name, product.Description, paisaPrice)
		if err != nil {
			slog.Error("Error creating product on Stripe",
				slog.String(logkey.TraceID, traceId),
				slog.String(logkey.ERROR, err.Error()),
			)
			return
		}

		slog.Info("Product created successfully on Stripe",
			slog.String(logkey.TraceID, traceId),
		)
	}()

	// Respond with the created product information
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Product created successfully",
		"product": product,
	})
}
