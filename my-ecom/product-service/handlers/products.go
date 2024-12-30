package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
	"product-service/internal/products"
	"product-service/pkg/logkey"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateProduct(c *gin.Context) {
	fmt.Println("CreateProduct Handler called!!")

	if c.Request.ContentLength > 5*1024 {
		slog.Error("request body limit breached",
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
			slog.String(logkey.ERROR, err.Error()),
		)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	if err := h.validate.Struct(newProduct); err != nil {
		slog.Error("validation failed",
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide values in correct format",
		})
		return
	}

	err := newProduct.ValidatePrice()
	if err != nil {
		slog.Error("price validation failed",
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "please provide valid price in INR",
		})
		return
	}

	ctx := c.Request.Context()

	product, err := h.p.InsertProduct(ctx, newProduct)
	if err != nil {
		slog.Error("error in creating the product",
			slog.String(logkey.ERROR, err.Error()),
		)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Product Creation Failed",
		})
		return
	}
	err = h.p.CreateProductStripe(ctx, product.ID, product.Name, product.Description, product.Price)
	if err != nil {
		slog.Error("Error creating product on Stripe",
			slog.String(logkey.ERROR, err.Error()),
		)
		return
	}

	slog.Info("Product created successfully on Stripe")

	// go func() {
		
	// }()

	c.JSON(http.StatusOK, product)
}
