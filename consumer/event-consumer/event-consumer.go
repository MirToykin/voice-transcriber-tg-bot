package event_consumer

import (
	"context"
	"github.com/MirToykin/voice-transcriber-tg-bot/events"
	"github.com/MirToykin/voice-transcriber-tg-bot/storage"
	"github.com/pkg/errors"
	"log"
	"sync"
	"time"
)

type Consumer struct {
	fetcher   events.Fetcher
	processor events.Processor
	storage   storage.Storage
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, st storage.Storage, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		storage:   st,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	failsCount := 0
	for {
		if failsCount >= 15 {
			return errors.New("unable to start consumer")
		}
		gotEvents, err := c.fetcher.Fetch(context.Background(), c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s", err.Error())
			failsCount++
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if len(gotEvents) == 0 {
			time.Sleep(1 * time.Second)
			continue
		}

		c.handleEvents(ctx, gotEvents)
	}
}

func (c *Consumer) handleEvents(ctx context.Context, eventsList []events.Event) {
	wg := sync.WaitGroup{}
	wg.Add(len(eventsList))

	for _, e := range eventsList {
		go c.handleEvent(ctx, &wg, &e)
	}

	wg.Wait()
}

func (c *Consumer) handleEvent(ctx context.Context, wg *sync.WaitGroup, event *events.Event) {
	defer wg.Done()
	task := withRetry(
		func() error {
			return c.processor.Process(ctx, event)
		}, 3,
	)
	err := task()

	if err != nil {
		log.Printf("ERROR: failed to process event %d: %s\n", event.ID, err)
		err = c.storage.SaveUnprocessed(ctx, event)
		if err != nil {
			log.Printf("ERROR: failed to save unprocessed event %d: %s\n", event.ID, err)
		}
	}
}

func withRetry(task func() error, maxAttempts int) func() error {
	return func() error {
		attempt := 0

		for {
			attempt++
			err := task()
			if err == nil {
				return nil
			} else if attempt < maxAttempts {
				time.Sleep(100 * time.Millisecond)
				continue
			} else {
				return err
			}
		}
	}

}
