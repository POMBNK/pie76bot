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
	updates, err := b.initUpdateConfig()
	if err != nil {
		return fmt.Errorf("failed to start bot due error:%w", err)
	}

	b.handleUpdates(updates)

	return nil
}

func (b *Bot) handleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		if update.Message.Text != "" {
			err := b.handleButton(update.Message)
			if err != nil {
				b.logs.Errorf("failed to handle update:%s", err)
				continue
			}
		}
	}
}

func (b *Bot) initUpdateConfig() (tgbotapi.UpdatesChannel, error) {
	updatesConfig := tgbotapi.NewUpdate(0)
	updatesConfig.Timeout = 60

	updates, err := b.bot.GetUpdatesChan(updatesConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to get updates channel due error:%w", err)
	}
	return updates, err
}
