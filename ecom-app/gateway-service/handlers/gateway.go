package handlers

import (
	"fmt"
	"gateway-service/internal/consul"
	"github.com/gin-gonic/gin"
	consulapi "github.com/hashicorp/consul/api"
	"io"
	"net/http"
	"strings"
)

type Handler struct {
	client *consulapi.Client
}

func NewHandler(client *consulapi.Client) *Handler {
	return &Handler{
		client: client,
	}
}

func (h *Handler) APIGateway(c *gin.Context) {
	//fmt.Println("Received request for " + c.Request.URL.Path)
	fullPath := c.Param("path") // give full path /users/create/123
	//fmt.Println(fullPath)
	segments := strings.Split(fullPath, "/")
	//fmt.Printf("%#v\n", segments)
	var serviceEndpoint string
	if len(segments) > 1 && segments[1] != "" {
		serviceEndpoint = segments[1]
	} else {
		return
	}

	fmt.Println(serviceEndpoint)
	pair, _, err := h.client.KV().Get(serviceEndpoint, nil)

	// make sure to check for pari == nil
	//because err indicates if connection to kv store was successful or not
	if err != nil || pair == nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{"error": "Service not found"},
		)
		fmt.Println("Service not found for " + c.Request.URL.Path)
		return
	}

	serviceName := string(pair.Value)
	fmt.Println("Service name is " + serviceName)

	serviceAddress, servicePort, err := consul.GetService(h.client, serviceName)
	if err != nil {
		// Respond with 503 if the service is unavailable in Consul
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Failed to reach service"})
		fmt.Println(err)
		return
	}

	ctx := c.Request.Context()
	httpQuery := fmt.Sprintf("http://%s:%d%s", serviceAddress, servicePort, fullPath)
	fmt.Println(httpQuery)

	req, err := http.NewRequestWithContext(ctx, c.Request.Method, httpQuery, c.Request.Body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to create request"})
		return
	}

	req.Header = c.Request.Header

	// this calls the service
	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to reach service"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response"})
		return
	}
	// Forward the service's response back to the client
	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)

}
