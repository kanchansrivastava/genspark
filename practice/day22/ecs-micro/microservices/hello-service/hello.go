package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/hashicorp/consul/api"
	"os"
)


func main() {
	registerServiceConsul()

	r := gin.Default()
	r.GET("/greet/health-check", func(c *gin.Context) {
		c.String(200, "Hello Service health check working !!")
	})

	r.GET("/greet/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		if name == "" {
			c.String(http.StatusBadRequest, "Name is required")
			return
		}
		c.String(200, "Hello"+" "+name+"!")
	})
	log.Println("Listening on port 80")
	r.Run(":80")
}


func registerServiceConsul() {
	config := api.DefaultConfig()
	config.Address = "http://consul.app:8500" 

	consul, err := api.NewClient(config)
	if err != nil {
		panic(err)
	}

	address := os.Getenv("HOSTNAME")
	if address == "" {
		panic("HOSTNAME not set")
	}

	registration := &api.AgentServiceRegistration{}
	registration.Name = "hello-service"
	registration.ID = "hello-service-" + address
	registration.Address = address
	registration.Port = 80

	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}
