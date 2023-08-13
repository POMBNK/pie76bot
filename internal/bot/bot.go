package bot

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"pie76bot/internal/hero"
	"pie76bot/internal/user"
	"pie76bot/pkg/logger"
)

type Bot struct {
	bot         *tgbotapi.BotAPI
	logs        *logger.Logger
	userService user.Service
	heroService hero.Service
	//pieService
}

func New(bot *tgbotapi.BotAPI, logs *logger.Logger, userService user.Service, heroService hero.Service) *Bot {
	return &Bot{
		bot:         bot,
		logs:        logs,
		userService: userService,
		heroService: heroService,
	}
}

func (b *Bot) Start() error {
	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		return err
	}

	//for keyboard
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			err = b.handleButton(update.Message)
			if err != nil {
				return fmt.Errorf("failed start tg bot due error:%w", err)
			}
		}
	}
	return nil
}
