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
	storage   storage.EventStorage
	batchSize int
}

func New(fetcher events.Fetcher, processor events.Processor, st storage.EventStorage, batchSize int) *Consumer {
	return &Consumer{
		fetcher:   fetcher,
		processor: processor,
		storage:   st,
		batchSize: batchSize,
	}
}

func (c *Consumer) Start(ctx context.Context) error {
	log.Println("Start fetching events...")
	failsCount := 0
	for {
		if failsCount >= 15 {
			return errors.New("unable to start consumer")
		}
		gotEvents, err := c.fetcher.Fetch(context.Background(), c.batchSize)
		if err != nil {
			log.Printf("[ERR] consumer: %s\n", err.Error())
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

func (c *Consumer) StartUnprocessed(ctx context.Context) error {
	log.Println("Start fetching unprocessed events...")
	failsCount := 0
	for {
		if failsCount >= 15 {
			return errors.New("unable to start handling unprocessed events")
		}
		storageEvents, err := c.storage.FetchUnprocessed(ctx, c.batchSize)
		if err != nil {
			log.Printf("[ERR] handling unprocessed events: %s\n", err.Error())
			failsCount++
			time.Sleep(500 * time.Millisecond)
			continue
		}

		if len(storageEvents) == 0 {
			time.Sleep(30 * time.Second)
			continue
		}

		c.handleUnprocessedEvents(ctx, storageEvents)
		time.Sleep(30)
	}
}

func (c *Consumer) handleEvents(
	ctx context.Context,
	eventsList []*events.Event,
) {
	wg := sync.WaitGroup{}
	wg.Add(len(eventsList))

	for _, e := range eventsList {
		go c.handleEvent(ctx, &wg, e)
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
		var processingError *events.ProcessingError
		isCustomErr := errors.As(err, &processingError)

		if !isCustomErr || processingError.NeedRetry {
			log.Printf("ERROR: failed to process event (external id: %v): %s\n", event.ExternalID, err)

			err = c.storage.SaveUnprocessed(ctx, event)
			if err != nil {
				log.Printf("ERROR: failed to save unprocessed event %d: %s\n", event.ID, err)
			}
		} else {
			log.Printf("ERROR processing event with no retry: %s", processingError)
		}
	}
}

func (c *Consumer) handleUnprocessedEvents(
	ctx context.Context,
	events []*events.Event,
) {
	wg := sync.WaitGroup{}
	wg.Add(len(events))

	for _, e := range events {
		go c.handleUnprocessedEvent(ctx, &wg, e)
	}

	wg.Wait()
}

func (c *Consumer) handleUnprocessedEvent(ctx context.Context, wg *sync.WaitGroup, event *events.Event) {
	defer wg.Done()

	task := withRetry(
		func() error {
			return c.processor.Process(ctx, event)
		}, 3,
	)
	err := task()

	if err != nil {
		var processingError *events.ProcessingError
		isCustomErr := errors.As(err, &processingError)
		if !isCustomErr || processingError.NeedRetry {
			log.Printf("ERROR: failed to process unprocessed event %d: %s\n", event.ID, err)
			return
		} else {
			log.Printf("ERROR processing unprocessed with no retry: %s", processingError.Error())
		}
	}

	err = c.storage.SetProcessed(ctx, event.ID)
	if err != nil {
		log.Printf("failed to set event %d as processed", event.ID)
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
