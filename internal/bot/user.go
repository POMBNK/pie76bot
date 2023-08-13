package bot

import (
	"pie76bot/internal/user"
)

func (b *Bot) isUserExist(existUser user.User) bool {
	if existUser.TelegramID != "" {
		return true
	}
	return false
}
