package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
)

var routeMap = map[string]string{
	"/greet/user/":  "hello-service",
	"/greet/health-check": "hello-service",
	"/user/register/": "user-service",
	"/user/health-check": "user-service",
}

func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received for: %s\n", r.URL.Path)
	serviceName, ok := routeMap[r.URL.Path]
	if !ok {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}

	proxyToService(serviceName, w, r)
}


func proxyToService(serviceName string, w http.ResponseWriter, r *http.Request) {
    config := api.DefaultConfig()
    config.Address = "http://consul.app:8500"

    consul, err := api.NewClient(config)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Query Consul for the service with the given name.
    services, _, err := consul.Health().Service(serviceName, "", true, nil)

	// 	if err != nil || len(services) == 0 {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    if len(services) == 0 {
        http.Error(w, "No healthy instances of service found", http.StatusInternalServerError)
        return
    }

    // Select the first available healthy service instance
    service := services[0]
    serviceAddress := fmt.Sprintf("http://%s:%d%s", service.Service.Address, service.Service.Port, r.URL.Path)

    // Make a proxy request to the selected service instance
    res, err := http.Get(serviceAddress)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    b, err := io.ReadAll(res.Body)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(res.StatusCode)
    w.Write(b)
}


func main() {
	http.HandleFunc("/", gatewayHandler)
	r := gin.Default()
	r.GET("/health-check", func(c *gin.Context) {
		c.String(200, "Hello Service health check working !!")
	})

	panic(http.ListenAndServe(":80", nil))
}
