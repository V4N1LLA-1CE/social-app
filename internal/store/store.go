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
	Users    UserRepository
	Posts    PostRepository
	Comments CommentRepository
}

func NewStore(db *sql.DB) Store {
	return Store{
		Users:    &UserStore{db},
		Posts:    &PostStore{db},
		Comments: &CommentStore{db},
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

type CommentRepository interface {
	GetByPostID(context.Context, int64) ([]Comment, error)
}
