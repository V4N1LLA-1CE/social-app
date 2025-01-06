package store

import (
	"context"
	"database/sql"
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
}
