package hero

import "context"

type Storage interface {
	CreateHero(ctx context.Context, hero Hero) (string, error)
	GetHero(ctx context.Context, userID string) ([]Hero, error)
}
