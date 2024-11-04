package telegram

import (
	"context"
	tgClient "github.com/MirToykin/voice-transcriber-tg-bot/clients/telegram"
	"github.com/pkg/errors"
	"log"
	"strings"
)

const (
	HelpCmd  = "/help"
	StartCmd = "/start"
)

func (p *Processor) doTextCmd(ctx context.Context, text string, chatID int, user *tgClient.From) error {
	text = strings.TrimSpace(text)

	log.Printf("got new command '%s' from '%s'", text, user.Username)

	switch text {
	case HelpCmd:
		return p.sendHelp(ctx, chatID, user)
	case StartCmd:
		return p.sendHello(ctx, chatID, user)
	default:
		return p.tgClient.SendMessage(ctx, chatID, msgUnknownCommand)
	}
}

func (p *Processor) sendHelp(ctx context.Context, chatID int, user *tgClient.From) error {
	return p.tgClient.SendMessage(ctx, chatID, getHelpMsg(user.LanguageCode))
}

func (p *Processor) sendHello(ctx context.Context, chatID int, user *tgClient.From) error {
	return p.tgClient.SendMessage(ctx, chatID, getHelloMsg(user.Username, user.LanguageCode))
}

func (p *Processor) sendTranscription(ctx context.Context, meta Meta, filePath string, lang *string) (err error) {
	transcribedText, err := p.transcriber.TranscribeByPath(ctx, filePath, lang)
	if err != nil {
		return errors.Wrap(err, "failed to send transcription")
	}

	if transcribedText == "" {
		return p.tgClient.SendMessage(ctx, meta.ChatID, "К сожалению, не удалось распознать сообщение")
	}

	return p.tgClient.SendReplyMessage(ctx, meta.ChatID, transcribedText, meta.MessageID)
}
