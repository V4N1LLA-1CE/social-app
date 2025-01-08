package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrNotFound = errors.New("Record not found")
)

type Store struct {
	Users UserRepository
	Posts PostRepository
}

func NewStore(db *sql.DB) Store {
	return Store{
		Users: &UserStore{db: db},
		Posts: &PostStore{db: db},
	}
}

// all repositories
type UserRepository interface {
	Create(context.Context, *User) error
}

type PostRepository interface {
	Create(context.Context, *Post) error
	GetById(context.Context, int64) (*Post, error)
}
