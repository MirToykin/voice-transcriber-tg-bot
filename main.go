package main

import (
	"context"
	tgClient "github.com/MirToykin/voice-transcriber-tg-bot/clients/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/config"
	event_consumer "github.com/MirToykin/voice-transcriber-tg-bot/consumer/event_consumer"
	"github.com/MirToykin/voice-transcriber-tg-bot/events/telegram"
	"github.com/MirToykin/voice-transcriber-tg-bot/storage/sqlc"
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion/whisper"
	"github.com/MirToykin/voice-transcriber-tg-bot/transcribtion/whisper/generated"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"sync"
)

const (
	tgBotHost = "api.telegram.org"
	batchSize = 100
)

func main() {
	cfg := config.MustLoad()
	ctx := context.TODO()

	conn, err := grpc.NewClient(cfg.TranscriberHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to transcriber service: %s", err)
	}

	defer func() { _ = conn.Close() }()

	whisperTranscriber := whisper.New(generated.NewTranscriptionServiceClient(conn))
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		whisperTranscriber,
	)

	storage, err := sqlc.New(cfg.DbDSN)
	if err != nil {
		log.Fatalf("failed to get storage: %s", err)
	}

	consumer := event_consumer.New(&eventsProcessor, &eventsProcessor, storage, batchSize)
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
