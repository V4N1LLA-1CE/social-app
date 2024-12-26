package store

import (
	"context"
	"database/sql"
)

type UserRepo struct {
	db *sql.DB
}

func (r *UserRepo) Create(ctx context.Context) error {
	return nil
}
