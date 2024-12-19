package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"github.com/hashicorp/consul/api"
	"os"
)


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
	registration.Name = "user-service"
	registration.ID = "user-service-" + address
	registration.Address = address
	registration.Port = 80

	err = consul.Agent().ServiceRegister(registration)
	if err != nil {
		panic(err)
	}
}


func main() {
	registerServiceConsul()

	r := gin.Default()
	r.GET("/user/health-check", func(c *gin.Context) {
		c.String(200, "User service is running")
	})

	r.POST("/user/register", func(c *gin.Context) {
		var request struct {
			Username string `json:"username"`
			Email    string `json:"email"`
		}

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
			return
		}

		log.Printf("New user registered: %+v\n", request)

		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "User registered successfully",
		})
	})

	log.Println("User-service is running on port 80")
	r.Run(":80")
}
