package hero

import (
	"context"
	"pie76bot/pkg/logger"
)

type Service interface {
	CreateHero(ctx context.Context, heroDTO ToCreateDTO) (string, error)
	GetHero(ctx context.Context, userID string) ([]Hero, error)
}

type service struct {
	logs    *logger.Logger
	storage Storage
}

func (s *service) CreateHero(ctx context.Context, heroDTO ToCreateDTO) (string, error) {
	hero := CreateHeroDTO(heroDTO)

	return s.storage.CreateHero(ctx, hero)
}

func (s *service) GetHero(ctx context.Context, userID string) ([]Hero, error) {
	return s.storage.GetHero(ctx, userID)
}

func NewService(logs *logger.Logger, storage Storage) Service {
	return &service{
		logs:    logs,
		storage: storage,
	}
}
