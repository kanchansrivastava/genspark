package handlers

import (
	"book-store/models"
	"github.com/gin-gonic/gin"
	"log"
)

func (h *Handler) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func (h *Handler) CreateBook(c *gin.Context) {
	var nb models.NewBook
	err := c.ShouldBindJSON(&nb)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	err = h.validate.Struct(nb)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"error": "please provide values in correct format"})
		return
	}
	book, err := h.c.InsertBook(c.Request.Context(), nb)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"error": "problem inserting book"})
		return
	}
	c.JSON(201, book)
}
