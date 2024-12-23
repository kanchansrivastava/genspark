package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

// connect kafka using kgo

// create a record using kgo.Record{key:value}

// produceRecord:=client.ProduceSync(ctx,record)

// produceRecord.First()

func main() {
	seeds := []string{"kafka-service:9092"}

	var client *kgo.Client
	var err error
	for i := 0; i < 5; i++ {
		client, err = kgo.NewClient(
			kgo.SeedBrokers(seeds...),
			kgo.AllowAutoTopicCreation(),
		)
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
		err = client.Ping(context.Background())
		if err != nil {
			time.Sleep(2 * time.Second)
			continue
		}
	}
	if err != nil {
		panic(err)
	}
	if client == nil {
		panic("client is nil")
	}
	defer client.Close()

	record := &kgo.Record{Topic: "hello", Value: []byte("barbiz")}
	pr := client.ProduceSync(context.Background(), record)
	rec, err := pr.First()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(rec.Value))
}
