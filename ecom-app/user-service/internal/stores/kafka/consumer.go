package kafka

import (
	"context"
	"github.com/twmb/franz-go/pkg/kgo"
	"log/slog"
	"time"
)

type ConsumeResult struct {
	Record *kgo.Record
}

func (c *Conf) ConsumeMessage(ctx context.Context, ch chan ConsumeResult) {

	for {
		fetches := c.consumer.PollFetches(ctx)
		if errs := fetches.Errors(); len(errs) > 0 {

			//maybe kafka is down
			slog.Error("ERROR: ", errs)
			time.Sleep(5 * time.Second)
			continue
		}

		iter := fetches.RecordIter()
		for !iter.Done() {
			rec := iter.Next()
			ch <- ConsumeResult{
				Record: rec,
			}
		}

	}
}
