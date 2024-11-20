package main

import (
	"context"
	"github.com/MirToykin/voice-transcriber-tg-bot/consumer/event_consumer"
	tgClient "github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/telegram"
	trbClient "github.com/MirToykin/voice-transcriber-tg-bot/internal/clients/transcriber"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/config"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/service/telegram_events"
	trbService "github.com/MirToykin/voice-transcriber-tg-bot/internal/service/transcriber"
	"github.com/MirToykin/voice-transcriber-tg-bot/internal/storage/sqlc"
	"github.com/gomscourse/common/pkg/closer"
	"log"
	"sync"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	cfg := config.MustLoad()
	ctx := context.TODO()

	trClient, err := trbClient.New(cfg.TranscriberHost)
	if err != nil {
		log.Fatalf("failed to get transcriber client: %s", err)
	}

	trService := trbService.New(trClient)
	eventsService := telegram_events.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		trService,
	)

	storage, err := sqlc.New(cfg.DbDSN)
	if err != nil {
		log.Fatalf("failed to get storage: %s", err)
	}

	consumer := event_consumer.New(eventsService, storage, batchSize)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := consumer.Start(ctx); err != nil {
			log.Fatalf("service is unable to start consuming events: %s", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := consumer.StartUnprocessed(ctx); err != nil {
			log.Fatalf("service is unable to start consuming unprocessed events: %s", err)
		}
	}()

	wg.Wait()
}
