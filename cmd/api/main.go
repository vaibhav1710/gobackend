package main

import (
	"log"

	"github.com/vaibhav1710/gobackend/internal/db"
	"github.com/vaibhav1710/gobackend/internal/env"
	"github.com/vaibhav1710/gobackend/internal/store"
)

func main() {
	// loading from env and giving a fallback value
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:password@localhost/udemy?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOpenConns,
		cfg.db.maxIdleConns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}

	defer db.Close()

	store.NewStorage(db)

	app := &application{
		config: cfg,
		store:  store.NewStorage(nil),
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
