package handlers

import "github.com/gin-gonic/gin"

func SetupGINRoutes() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Ping)
	return r
}
