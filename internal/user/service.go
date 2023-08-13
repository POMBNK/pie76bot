package user

import (
	"context"
	"fmt"
	"pie76bot/pkg/logger"
)

type Service interface {
	CreateUser(ctx context.Context, userDTO ToCreateDTO) error
	GetUser(ctx context.Context, userID string) (User, error)
	SignUP(ctx context.Context, userDTO ToCreateDTO) error
}

type service struct {
	logs    *logger.Logger
	storage Storage
}

func (s *service) CreateUser(ctx context.Context, userDTO ToCreateDTO) error {
	user := CreateUserDTO(userDTO)
	return s.storage.CreateUser(ctx, user)
}

func (s *service) GetUser(ctx context.Context, telegramID string) (User, error) {
	u, err := s.storage.GetUser(ctx, telegramID)
	if err != nil {
		return User{}, err
	}
	return u, nil
}

func (s *service) SignUP(ctx context.Context, userDTO ToCreateDTO) error {
	user := CreateUserDTO(userDTO)

	existedUser, err := s.storage.GetUser(ctx, user.TelegramID)
	if err != nil {
		return err
	}
	if existedUser.TelegramID == user.TelegramID {
		return fmt.Errorf("this user already exist")
	}
	err = s.storage.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	return nil
}

func NewService(logs *logger.Logger, storage Storage) Service {
	return &service{
		logs:    logs,
		storage: storage,
	}
}
