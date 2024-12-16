package handlers

import (
	"book-store/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	c        *models.Conn
	validate *validator.Validate
}

func NewConf(c *models.Conn, validate *validator.Validate) *Handler {
	return &Handler{c: c, validate: validate}
}
func SetupGINRoutes(c *models.Conn) *gin.Engine {
	r := gin.Default()

	h := NewConf(c, validator.New())
	r.GET("/ping", h.Ping)
	r.POST("/create", h.CreateBook)
	return r
}
