package store

import (
	"context"
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Post struct {
	ID        int64     `json:"id"`
	Content   string    `json:"content"`
	Title     string    `json:"title"`
	UserID    int64     `json:"user_id"`
	Tags      []string  `json:"tags"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostRepo struct {
	db *sql.DB
}

func (r *PostRepo) Create(ctx context.Context, post *Post) error {
	query := `
  INSERT INTO posts (content, title, user_id, tags)
  VALUES ($1, $2, $3, $4)
  RETURNING id, created_at, updated_at
  `

	args := []any{
		post.Content,
		post.Title,
		post.UserID,
		pq.Array(post.Tags),
	}

	err := r.db.QueryRowContext(ctx, query, args...).Scan(
		&post.ID,
		&post.CreatedAt,
		&post.UpdatedAt,
	)

	if err != nil {
		return err
	}

	return nil
}