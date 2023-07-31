package user

import (
	"context"
	"pie76bot/pkg/logger"
)

type Service interface {
	CreateUser(ctx context.Context, userDTO ToCreateDTO) error
	GetUser(ctx context.Context, userID string) (User, error)
}

type service struct {
	logs    logger.Logger
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, userDTO ToCreateDTO) error {
	var dummyUser User // TODO: Put telegramID to field of USER struct
	return s.storage.CreateUser(ctx, dummyUser)
}

func (s *service) GetUser(ctx context.Context, userID string) (User, error) {
	return s.storage.GetUser(ctx, userID)
}
