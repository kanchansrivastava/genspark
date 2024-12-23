package main

import (
	"context"
	"fmt"
	"github.com/twmb/franz-go/pkg/kgo"
	"time"
)

func main() {
	seeds := []string{"kafka-service:9092"}
	// One client can both produce and consume!
	// Consuming can either be direct (no consumer group), or through a group. Below, we use a group.
	cl, err := kgo.NewClient(
		kgo.SeedBrokers(seeds...),
		kgo.ConsumerGroup("my-group-identifier"),
		kgo.ConsumeTopics("hello"),
		kgo.FetchMinBytes(1),
		kgo.FetchMaxWait(10*time.Millisecond),

		//kgo.DisableAutoCommit(), // Disable auto-commit for manual control
	)
	if err != nil {
		panic(err)
	}
	defer cl.Close()
	ctx := context.Background()
	for {

		fetches := cl.PollFetches(ctx)
		fmt.Println("fetched")
		if errs := fetches.Errors(); len(errs) > 0 {
			// All errors are retried internally when fetching, but non-retriable errors are
			// returned from polls so that users can notice and take action.
			fmt.Println("errors", errs)
			fmt.Println("waiting for kafka to come up again")
			time.Sleep(5 * time.Second)
			continue
		}

		// We can iterate through a record iterator...
		iter := fetches.RecordIter()
		for !iter.Done() {
			record := iter.Next()
			fmt.Println(string(record.Value), "from an iterator!")
		}

	}
}
