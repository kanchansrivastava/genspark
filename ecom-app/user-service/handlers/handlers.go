package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"os"
	"user-service/internal/auth"
	"user-service/internal/stores/kafka"
	"user-service/internal/users"
	"user-service/middleware"
)

type Handler struct {
	u        *users.Conf
	validate *validator.Validate
	k        *kafka.Conf
	a        *auth.Keys
}

func NewHandler(u *users.Conf, a *auth.Keys, k *kafka.Conf) *Handler {
	return &Handler{
		u:        u,
		k:        k,
		a:        a,
		validate: validator.New(),
	}
}

func API(u *users.Conf, a *auth.Keys, k *kafka.Conf) *gin.Engine {
	r := gin.New()
	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	h := NewHandler(u, a, k)

	prefix := os.Getenv("SERVICE_ENDPOINT_PREFIX")
	if prefix == "" {
		panic("SERVICE_ENDPOINT_PREFIX is not set")
	}
	v1 := r.Group(prefix)
	v1.Use(middleware.Logger())
	{
		v1.Use(gin.Logger(), gin.Recovery())
		v1.POST("/signup", h.Signup)
		v1.POST("/login", h.Login)

		// this middleware would be applied to the handler functions which are after it
		// it would not apply to the previous one
		v1.Use(middleware.Authentication())
		v1.GET("/check", func(c *gin.Context) {
			c.JSON(200, gin.H{"Auth Check": "You are authenticated"})
		})
	}
	return r
}
