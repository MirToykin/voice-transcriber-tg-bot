package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	tgClient "voice_transcriber_bot/clients/telegram"
	"voice_transcriber_bot/config"
	"voice_transcriber_bot/events/telegram"
	"voice_transcriber_bot/transcribtion/whisper"
	"voice_transcriber_bot/transcribtion/whisper/generated"
)

const (
	tgBotHost       = "api.telegram.org"
	transcriberHost = "transcriber:50051"
	storagePath     = "files_storage"
	batchSize       = 100
)

func main() {
	cfg := config.MustLoad()

	conn, err := grpc.NewClient(transcriberHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to transcriber service: %s", err)
	}

	defer func() { _ = conn.Close() }()

	whisperTranscriber := whisper.New(generated.NewTranscriptionServiceClient(conn))
	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		whisperTranscriber,
	)

}
