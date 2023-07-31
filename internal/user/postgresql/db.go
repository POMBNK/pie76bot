package postgresql

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"pie76bot/internal/user"
	"pie76bot/pkg/client/postgresql"
	"pie76bot/pkg/logger"
)

type storage struct {
	logs   logger.Logger
	client postgresql.Client
}

func (s *storage) CreateUser(ctx context.Context, user user.User) error {
	q := `INSERT INTO "user" (id) VALUES ($1)`
	_, err := s.client.Exec(ctx, q, user.TelegramID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return err
		}
		return fmt.Errorf("can't create user due error:%w", err)
	}

	return nil
}

func (s *storage) GetUser(ctx context.Context, userID string) (user.User, error) {
	var tgUser user.User
	q := `SELECT id FROM "user" WHERE id = id`
	err := s.client.QueryRow(ctx, q, userID).Scan(&tgUser.TelegramID)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return user.User{}, err
		}
		return user.User{}, fmt.Errorf("can't get user due error:%w", err)
	}

	return tgUser, nil
}

func NewStorage(logs logger.Logger, client postgresql.Client) user.Storage {
	return &storage{
		logs:   logs,
		client: client,
	}
}
