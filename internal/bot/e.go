package bot

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

func (b *Bot) wrapErr(message *tgbotapi.Message) error {
	err := b.handleUnknownCommand(message)
	if err != nil {
		return err
	}
	return err
}
