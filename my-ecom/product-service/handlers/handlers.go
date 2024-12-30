package handlers

import (
	"product-service/internal/products"
	"github.com/gin-gonic/gin"
	"os"
	"net/http"
)

type Handler struct {
	p *products.Conf
}

func NewHandler(p *products.Conf) *Handler {
	return &Handler{
		p: p,
	}
}

func API(p *products.Conf) *gin.Engine{
	r := gin.New()
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}
	// h := NewHandler(p)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ping successful"})
	})
	return r
}
