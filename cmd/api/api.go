package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type application struct {
	config config
}

type config struct {
	addr string
}

func (app *application) mount() http.Handler {
	//This creates a new mux (short for "multiplexer"), which is responsible for routing incoming requests to the correct handler
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Recoverer) // to recover from panic
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthCheckHandler)
	})

	return r
}

func (app *application) run(mux http.Handler) error {

	//This creates an HTTP server. It uses the mux to handle requests and the addr (from the config struct) to know which address to listen to.
	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30, // This is the maximum time the server will wait for sending the response back to the client. If the server takes longer than 30 seconds to send the response (for example, if generating the response is a lengthy process), the server will close the connection and return an error.
		ReadTimeout:  time.Second * 10, //  maximum time the server will wait for the client to send the request. 		If the client takes longer than 10 seconds to send the entire request (e.g., a very large body or slow network), the server will close the connection and return an error.
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)

	//This starts the HTTP server and listens for incoming requests on the configured address
	return srv.ListenAndServe()
}
