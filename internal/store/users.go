package store

import (
	"context"
	"database/sql"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  Password  `json:"-"`
	Activated bool      `json:"activation_status"`
	CreatedAt time.Time `json:"created_at"`
}

type Password struct {
	PasswordHash      []byte
	PasswordPlaintext string
}

func (p *Password) Set(plaintext string) error {
	// hash password with bcrypt
	h, err := bcrypt.GenerateFromPassword([]byte(plaintext), 12)
	if err != nil {
		return err
	}

	p.PasswordPlaintext = plaintext
	p.PasswordHash = h

	return nil
}

type UserStore struct {
	db *sql.DB
}

func (s *UserStore) Create(ctx context.Context, user *User) error {
	query := `
  INSERT INTO users (email, password, activated) 
  VALUES ($1, $2, $3)
  RETURNING id, created_at
  `

	args := []any{
		user.Email,
		user.Password.PasswordHash,
		false,
	}

	err := s.db.QueryRowContext(ctx, query, args...).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "users_email_key"):
			return ErrDuplicateEmail
		default:
			return err
		}
	}

	return nil
}
