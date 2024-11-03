package telegram

import (
	"context"
	"github.com/pkg/errors"
	"log"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doTextCmd(ctx context.Context, text string, chatID int, username string) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s", text, username)

	switch text {
	case HelpCmd:
		return p.sendHelp(ctx, chatID)
	case StartCmd:
		return p.sendHello(ctx, chatID)
	default:
		return p.tgClient.SendMessage(ctx, chatID, msgUnknownCommand)
	}
}

func (p *Processor) sendHelp(ctx context.Context, chatID int) error {
	return p.tgClient.SendMessage(ctx, chatID, msgHelp)
}

func (p *Processor) sendHello(ctx context.Context, chatID int) error {
	return p.tgClient.SendMessage(ctx, chatID, msgHello)
}

func (p *Processor) sendTranscription(ctx context.Context, chatID int, filePath string) (err error) {
	transcribedText, err := p.transcriber.TranscribeByPath(ctx, filePath)
	if err != nil {
		return errors.Wrap(err, "failed to send transcription")
	}

	return p.tgClient.SendMessage(ctx, chatID, transcribedText)
}
