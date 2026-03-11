package database

import (
	"context"
	"log"

	"github.com/Prayas-35/ragkit/engine/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() {
	cfg := config.LoadConfig()
	dsn := cfg.DatabaseUri
	if dsn == "" {
		log.Fatal("DATABASE_URL not set")
	}

	dbpool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	DB = dbpool
	log.Println("✅ Connected to Neon PostgreSQL")

	if cfg.DB_SYNC {
		log.Println("🔄 Running database migrations...")
		if err := migrate(dbpool); err != nil {
			log.Fatal("Failed to run migrations:", err)
		}
	}
}

func migrate(pool *pgxpool.Pool) error {

	ctx := context.Background()

	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id UUID PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE,
		password TEXT,
		created_at TIMESTAMP DEFAULT now()
	);
	`

	bookTable := `
	CREATE TABLE IF NOT EXISTS books (
		id UUID PRIMARY KEY,
		title TEXT,
		author TEXT,
		available BOOLEAN DEFAULT true
	);
	`

	issueTable := `
	CREATE TABLE IF NOT EXISTS issues (
		id UUID PRIMARY KEY,
		user_id UUID REFERENCES users(id),
		book_id UUID REFERENCES books(id),
		issued_at TIMESTAMP DEFAULT now(),
		returned_at TIMESTAMP
	);
	`

	queries := []string{userTable, bookTable, issueTable}

	for _, q := range queries {
		_, err := pool.Exec(ctx, q)
		if err != nil {
			return err
		}
	}

	return nil
}
