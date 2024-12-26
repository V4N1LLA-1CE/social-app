package store

import (
	"context"
	"database/sql"
)

type postRepo struct {
	db *sql.DB
}

func (r *postRepo) Create(ctx context.Context) error {
	return nil
}
