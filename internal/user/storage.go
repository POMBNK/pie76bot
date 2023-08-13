package user

import "context"

type Storage interface {
	CreateUser(ctx context.Context, user User) error
	GetUser(ctx context.Context, telegramID string) (User, error)
}
