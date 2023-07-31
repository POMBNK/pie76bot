package bot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"pie76bot/internal/user"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	userService user.Service
}

func New(bot *tgbotapi.BotAPI, service user.Service) *Bot {
	return &Bot{
		bot:         bot,
		userService: service,
	}
}

func (b *Bot) Start() error {
	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		return err
	}
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.IsCommand() {
			err = b.handleCommand(update.Message)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
