package main

import (
	"fmt"
	"log"
	"time"

	"github.com/V4N1LLA-1CE/social-app/internal/database"
	"github.com/V4N1LLA-1CE/social-app/internal/env"
	"github.com/V4N1LLA-1CE/social-app/internal/store"
)

func main() {
	cfg := config{
		addr: fmt.Sprintf(":%s", env.GetString("PORT", "8080")),
		db: dbConfig{
			dsn:          env.GetString("POSTGRES_DSN", "postgres://root:toor@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetInt("POSTGRES_MAXOPENCONNS", 30),
			maxIdleConns: env.GetInt("POSTGRES_MAXIDLECONNS", 30),
			maxIdleTime:  time.Duration(env.GetInt("POSTGRES_MAXIDLETIME", 15)) * time.Minute,
		},
	}

	conn, err := database.New(cfg.db.dsn, cfg.db.maxOpenConns, cfg.db.maxIdleConns, cfg.db.maxIdleTime)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	log.Println("database connection pool established")

	store := store.NewStore(conn)

	app := &application{
		config: cfg,
		store:  store,
	}

	mux := app.mount()

	log.Fatal(app.serve(mux))
}
