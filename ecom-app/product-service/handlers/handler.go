package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"product-service/internal/products"
	"product-service/middleware"
	"product-service/pkg/ctxmanage"
)

func API(p products.Conf, endpointPrefix string) *gin.Engine {

	r := gin.New()
	mode := os.Getenv("GIN_MODE")
	if mode == gin.ReleaseMode {
		gin.SetMode(mode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	//s := models.NewStore(&c)
	h := handler{Conf: p}
	//apply middleware to all the endpoints using r.Use
	r.Use(middleware.Logger(), gin.Recovery())
	r.GET("/ping", healthCheck)
	v1 := r.Group(endpointPrefix)
	{
		v1.Use(middleware.Logger())
		v1.POST("/create", h.CreateProduct)

	}

	return r
}

func healthCheck(c *gin.Context) {
	traceId := ctxmanage.GetTraceIdOfRequest(c)

	fmt.Println("healthCheck handler ", traceId)
	//JSON serializes the given struct as JSON into the response body. It also sets the Content-Type as "application/json".
	c.JSON(http.StatusOK, gin.H{"status": "ok"})

}
