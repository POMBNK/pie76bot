package bot

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

const (
	commandStart = "/start"
	buttonHeroes = "Heroes"
	buttonPies   = "Pies"
	buttonCreate = "Create"
	buttonShow   = "Show exist"
)

var commands = []string{commandStart, buttonHeroes, buttonPies, buttonCreate, buttonShow}

func (b *Bot) handleButton(message *tgbotapi.Message) error {
	switch message.Text {
	case commandStart:
		return b.handleStartButton(message)
	case buttonHeroes:
		return b.handleHeroesButton(message)
	case buttonCreate:
		return b.handleCreateHeroButton(context.Background(), message)
	default:
		return b.handleHeroStats(context.Background(), message)
	}
}

func (b *Bot) handleStartButton(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(buttonHeroes),
			tgbotapi.NewKeyboardButton(buttonPies),
		),
	)

	msg.ReplyMarkup = keyboard
	msg.Text = "Choose an action"
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "Wrong command or invalid input")
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return fmt.Errorf("invalid input")
}

func existCommand(input string, validCommands []string) bool {
	for _, c := range validCommands {
		if c == input {
			return true
		}
	}
	return false
}
