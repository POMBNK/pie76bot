package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"pie76bot/internal/hero"
	"pie76bot/pkg/client/postgresql"
	"pie76bot/pkg/logger"
)

type storage struct {
	logs   *logger.Logger
	client postgresql.Client
}

func (s *storage) GetHero(ctx context.Context, userID string) ([]hero.Hero, error) {
	q := `SELECT h.id,h.name,h.luck FROM hero h JOIN users_heroes uh on h.id = uh.heroid WHERE uh.userid = $1`
	rows, err := s.client.Query(ctx, q, userID)
	if err != nil {
		return nil, err
	}
	heroes := make([]hero.Hero, 0)
	for rows.Next() {
		var heroUnit hero.Hero
		err = rows.Scan(&heroUnit.Id, heroUnit.Name, heroUnit.Luck)
		if err != nil {
			return nil, err
		}
		heroes = append(heroes, heroUnit)
	}

	if rows.Err() != nil {
		return nil, err
	}

	return heroes, nil
}

func (s *storage) CreateHero(ctx context.Context, hero hero.Hero) (string, error) {
	q := `INSERT INTO hero (name, luck) VALUES ($1,$2) RETURNING id`
	err := s.client.QueryRow(ctx, q, hero.Name, hero.Luck).Scan(&hero.Id)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", err
		}
		return "", fmt.Errorf("can't create user due error:%w", err)
	}

	return hero.Id, nil
}

func NewStorage(logs *logger.Logger, client postgresql.Client) hero.Storage {
	return &storage{
		logs:   logs,
		client: client,
	}
}
