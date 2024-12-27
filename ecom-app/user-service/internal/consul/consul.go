package consul

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"os"
	"strconv"
)

func RegisterWithConsul() (*consulapi.Client, string, error) {
	//By default, in Docker, the value of HOSTNAME is set to the Ip
	// docker container address
	hostName := os.Getenv("HOSTNAME")
	svcName := os.Getenv("SERVICE_NAME")
	portString := os.Getenv("APP_PORT")
	consulAddress := os.Getenv("CONSUL_HTTP_ADDRESS")
	svcEndpointPrefix := os.Getenv("SERVICE_ENDPOINT_PREFIX")

	if hostName == "" || svcName == "" || portString == "" ||
		consulAddress == "" || svcEndpointPrefix == "" {
		return nil, "", errors.New(
			`env variables not set for hostName, 
                 svcName, port, consulAddress,svcEndpointPrefix`)
	}

	port, err := strconv.Atoi(portString)
	if err != nil {
		return nil, "", fmt.Errorf("port is not a number: %w", err)
	}
	config := consulapi.DefaultConfig()

	//Address is the address of the Consul server
	config.Address = consulAddress

	//creating a connection with the address
	client, err := consulapi.NewClient(config)
	if err != nil {
		return nil, "", fmt.Errorf("create consul client: %w", err)
	}

	registration := consulapi.AgentServiceRegistration{}
	regId := svcName + "-" + hostName
	// Set the unique ID for the service. This ID differentiates this service instance
	// from others, even when they have the same service name.
	registration.ID = regId // hostname is always unique, so no clashes

	// Set the name of the service. This is a logical identifier for the service
	// and is how other clients will find or query this service in Consul.
	registration.Name = svcName

	registration.Port = port

	// Set the address (hostname) where the service can be accessed. This is the IP/hostname
	// of the machine where the service is running.
	registration.Address = hostName

	// Define the health check for this service. Consul will periodically check the health of this service
	// using the details provided here. If the service fails this check, Consul may mark it as unavailable.
	registration.Check = &consulapi.AgentServiceCheck{
		// HTTP health check endpoint that Consul will query periodically to determine if the service is healthy.
		HTTP: fmt.Sprintf("http://%s:%d/ping", hostName, port),

		// The interval at which the health checks will be performed (e.g. 30s in this case)
		Interval: "30s",

		// Timeout specifies how long Consul will wait for the HTTP health check to respond.
		// If the service does not respond within this time, it is considered unhealthy.
		Timeout: "10s",

		// DeregisterCriticalServiceAfter specifies the time after which Consul will automatically
		// deregister the service if it remains in a critical health state. In this case, 30 seconds.
		DeregisterCriticalServiceAfter: "30s",
	}

	// Log a message to indicate that the service registration process is starting.
	fmt.Println("registering service with consul")

	// registering with consul
	err = client.Agent().ServiceRegister(&registration)
	if err != nil {
		return nil, "", fmt.Errorf("register service with consul: %w", err)
	}
	kv := client.KV()
	pair := consulapi.KVPair{
		Key:   svcEndpointPrefix,
		Value: []byte(svcName)}

	_, err = kv.Put(&pair, nil)
	if err != nil {
		return nil, "", fmt.Errorf("register service with consul: %w", err)
	}

	// If everything completes successfully, return nil to indicate there was no error.
	return client, regId, nil
}
