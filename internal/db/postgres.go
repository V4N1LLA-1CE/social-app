package db

import (
	"context"
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"time"
)

func New(dsn string, maxOpenConns, maxIdleConns int, maxIdleTime time.Duration) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	// configure db with settings
	conn.SetConnMaxIdleTime(maxIdleTime)
	conn.SetMaxIdleConns(maxIdleConns)
	conn.SetMaxOpenConns(maxOpenConns)

	// 5 second timeout if connection to db cannot be established
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = conn.PingContext(ctx)
	if err != nil {
		return nil, err
	}

	return conn, nil
}
