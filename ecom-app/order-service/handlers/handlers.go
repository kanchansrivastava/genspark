package handlers

import (
	"github.com/gin-gonic/gin"
	"order-service/internal/auth"
	"order-service/middleware"
	"os"
)

type Handler struct {
}

func (h Handler) Checkout(context *gin.Context) {

}

func NewHandler() *Handler {
	return &Handler{}
}

func API(endpointPrefix string, k *auth.Keys) *gin.Engine {
	r := gin.New()
	mode := os.Getenv("GIN_MODE")
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
	m, err := middleware.NewMid(k)
	if err != nil {
		panic(err)
	}

	h := NewHandler()
	r.Use(middleware.Logger(), gin.Recovery())

	r.GET("/ping", HealthCheck)
	v1 := r.Group(endpointPrefix)
	{
		v1.Use(m.Authentication())
		v1.POST("/checkout/:productID", h.Checkout)
		v1.GET("/ping", HealthCheck)
	}

	return r
}

func HealthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
