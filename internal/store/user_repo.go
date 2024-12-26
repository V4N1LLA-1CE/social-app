package store

import (
	"context"
	"database/sql"
)

type userRepo struct {
	db *sql.DB
}

func (r *userRepo) Create(ctx context.Context) error {
	return nil
}
