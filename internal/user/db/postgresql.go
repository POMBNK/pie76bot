package db

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"pie76bot/internal/user"
	"pie76bot/pkg/client/postgresql"
	"pie76bot/pkg/logger"
	"time"
)

var ErrNotFound = errors.New("no user found")

type storage struct {
	logs   *logger.Logger
	client postgresql.Client
}

func (s *storage) CreateUser(ctx context.Context, user user.User) error {
	q := `INSERT INTO users (tg_id) VALUES ($1) RETURNING id`
	err := s.client.QueryRow(ctx, q, user.TelegramID).Scan(&user.Uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		return fmt.Errorf("can't create user due error:%w", err)
	}

	return nil
}

func (s *storage) GetUser(ctx context.Context, telegramID string) (user.User, error) {
	var unitUser user.User
	q := `SELECT id FROM users WHERE tg_id = $1`
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()
	err := s.client.QueryRow(ctx, q, telegramID).Scan(&unitUser.Uuid)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.User{}, ErrNotFound //fmt.Errorf("no user found: %w", err)
		}
		return user.User{}, fmt.Errorf("can't get user due error:%w", err)
	}

	return unitUser, nil
}

func NewStorage(logs *logger.Logger, client postgresql.Client) user.Storage {
	return &storage{
		logs:   logs,
		client: client,
	}
}
