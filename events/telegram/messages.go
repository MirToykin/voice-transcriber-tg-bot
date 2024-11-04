package telegram

import "fmt"

func getHelpMsg(lang string) string {
	return fmt.Sprintf(
		`Я могу конвертировать голосовые сообщения в текстовый формат.
Для распознования исползуется язык, настроенный в боте, если язык не настроен, 
то по умолчанию будет испоьзован язык, установленный на исползуемом Telegram клиенте.
Текущий язык:  %s`, lang,
	)
}

func getHelloMsg(username, lang string) string {
	return fmt.Sprintf("Привет, @%s\n\n%s", username, getHelpMsg(lang))
}

const msgHelp = `Я могу конвертировать голосовые сообщения в текстовый формат. 
В настоящий момент поддерживается только русский язык.`

const msgHello = "Привет! 👾\n\n" + msgHelp

const (
	msgUnknownCommand = "Неизвестная команда 🤔"
)
