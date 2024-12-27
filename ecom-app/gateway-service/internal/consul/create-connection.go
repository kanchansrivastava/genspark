package consul

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"time"
)

func CreateConnection() (*consulapi.Client, error) {
	config := consulapi.DefaultConfig()
	config.Address = "http://consul.diwakar:8500"
	t := time.Now()
	var err error
	var client *consulapi.Client
	for {
		client, err = consulapi.NewClient(config)
		fmt.Println("consul New Client status ", err)
		if err != nil {
			time.Sleep(5 * time.Second)
			continue
		}
		var s string
		s, err = client.Status().Leader()
		if err != nil {
			fmt.Println("consul connection status ", err)
			time.Sleep(5 * time.Second)
			continue
		}
		fmt.Println(s)

		if time.Since(t) > 10*time.Minute {
			return nil, fmt.Errorf("consul connection timeout")
		}
		break

	}
	if err != nil {
		return nil, err
	}
	return client, nil

}
