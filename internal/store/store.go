package store

import (
	"context"
	"database/sql"
)

type Store struct {
	Users UserRepository
}

type UserRepository interface {
	Create(context.Context, *User) error
}

func NewStore(db *sql.DB) Store {
	return Store{
		Users: &UserStore{db: db},
	}
}
