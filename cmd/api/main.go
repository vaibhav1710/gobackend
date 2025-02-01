package main

import (
	"log"

	"github.com/vaibhav1710/gobackend/internal/env"
)

func main() {
	// loading from env and giving a fallback value
	cfg := config{
		addr: env.GetString("ADDR", ":3000"),
	}

	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
