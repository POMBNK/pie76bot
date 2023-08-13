package bot

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"pie76bot/internal/hero"
	"pie76bot/internal/user"
	"strconv"
	"strings"
)

func (b *Bot) handleHeroesButton(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, "")
	keyboard := tgbotapi.NewReplyKeyboard(
		tgbotapi.NewKeyboardButtonRow(
			tgbotapi.NewKeyboardButton(buttonCreate),
			tgbotapi.NewKeyboardButton(buttonShow),
		),
	)
	msg.ReplyMarkup = keyboard
	msg.Text = "Heroes"
	_, err := b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

// need full refactor
func (b *Bot) handleCreateHeroButton(ctx context.Context, message *tgbotapi.Message) error {

	var userDTO user.ToCreateDTO

	id := strconv.Itoa(int(message.Chat.ID))
	userDTO.TelegramID = id

	err := b.userService.SignUP(ctx, userDTO)
	if err != nil {
		return err
	}

	alert := tgbotapi.NewMessage(message.Chat.ID, "Input hero name and luck points using space")
	_, err = b.bot.Send(alert)
	if err != nil {
		return err
	}

	return nil
}

func (b *Bot) handleGetHeroButton(ctx context.Context, message *tgbotapi.Message) error {
	//show heroes service method
	telegramID := strconv.Itoa(int(message.Chat.ID))
	heroes, err := b.heroService.GetHero(ctx, telegramID)
	if err != nil {
		return err
	}
	var text string
	for _, unitHero := range heroes {
		text += unitHero.Id + ", "
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, text)
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}
	return nil
}

func (b *Bot) handleHeroStats(ctx context.Context, message *tgbotapi.Message) error {
	var heroDTO hero.ToCreateDTO
	text := message.Text
	if existCommand(text, commands) {
		return nil
	}

	stats := strings.Split(text, " ")
	if len(stats) != 2 {
		err := b.handleUnknownCommand(message)
		if err != nil {
			return err
		}
	}

	name := stats[0]
	//validateHeroName()
	luck, err := strconv.Atoi(stats[1])
	if err != nil {
		return err
	}
	//validateHeroLuck()
	heroDTO.Name = name
	heroDTO.Luck = luck
	_, err = b.heroService.CreateHero(ctx, heroDTO)
	if err != nil {
		return err
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Hero created")
	_, err = b.bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
