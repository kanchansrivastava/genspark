package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"io"
	"net/http"
)

// routeMap is a map that defines routes and maps them to corresponding service names.
// Here, the route `/user/ping` is mapped to the service `micro-1`.
var routeMap = map[string]string{
	"/user/ping": "micro-1",
}

// gatewayHandler is the main HTTP handler function for all requests coming into the gateway.
func gatewayHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("request received for: %s\n", r.URL.Path)

	// Check if the requested path exists in the routeMap.
	serviceName, ok := routeMap[r.URL.Path]
	if !ok {
		// If the path does not exist, return a 404 error (Not Found).
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}

	// If the service exists, proxy the request to the corresponding service.
	proxyToService(serviceName, w, r)
}

// proxyToService proxies the HTTP request to the appropriate service based on the service name.
func proxyToService(serviceName string, w http.ResponseWriter, r *http.Request) {
	// Create a default configuration for Consul.
	config := api.DefaultConfig()

	// Setting the address where Consul is running. Change this to point to your actual Consul server.
	config.Address = "http://consul.diwakarv1:8500"

	// Create a new client to interact with the Consul API.
	consul, err := api.NewClient(config)
	if err != nil {
		// If an error occurs while creating the Consul client, return a 500 error (Internal Server Error).
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Query Consul for the service with the given name.
	services, _, err := consul.Catalog().Service(serviceName, "", nil)
	if err != nil {
		// If an error occurs while querying Consul, return a 500 error.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// If no services are found, return a 404 error.
	if len(services) == 0 {
		http.Error(w, "service not found", http.StatusNotFound)
		return
	}

	// Pick the first available service instance (can be enhanced later for load balancing).
	service := services[0]

	// Construct the URL to forward the request to the service.
	// `ServiceAddress` and `ServicePort` are the address and port of the service found in Consul.
	serviceAddress := fmt.Sprintf("http://%s:%d%s", service.ServiceAddress, service.ServicePort, r.URL.Path)

	// Make an HTTP GET request to the constructed service address.
	res, err := http.Get(serviceAddress)
	if err != nil {
		// If an error occurs while forwarding the request, return a 500 error.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Read the response body from the service response.
	b, err := io.ReadAll(res.Body)
	if err != nil {
		// If an error occurs while reading the response body, return a 500 error.
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set the same response status code as that of the service response.
	w.WriteHeader(res.StatusCode)

	// Forward the service response back to the requester (client).
	w.Write(b)
}
func main() {
	http.HandleFunc("/", gatewayHandler)
	http.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("pong"))
	})
	//http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
	//	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, "http://micro-1.app/user/ping", nil)
	//
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	// calling the remote service
	//	res, err := http.DefaultClient.Do(req)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//	defer res.Body.Close()
	//
	//	data, err := io.ReadAll(res.Body)
	//	if err != nil {
	//		w.WriteHeader(http.StatusInternalServerError)
	//		http.Error(w, err.Error(), http.StatusInternalServerError)
	//		return
	//	}
	//
	//	w.WriteHeader(res.StatusCode)
	//	w.Write(data)
	//})
	panic(http.ListenAndServe(":80", nil))

}
