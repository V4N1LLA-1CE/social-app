package store

import (
	"context"
	"database/sql"
	"errors"
)

var (
	ErrDuplicateEmail = errors.New("account with this email already exists")
	ErrRecordNotFound = errors.New("record cannot be found")
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
