package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

// Declare a string containing the application version number. Later in the book
// we'll generate this automatically at the build time, but for now we'll just
// store the version number as a hard-coded global constant.
const version = "1.0.0"

// Define a config struct to hold all the configuration settings for our application.
// For now, the only configuration setting will be the network port that we
// want the server to listen on, and the name of the current operating env for
// the application (development, staging, production etc.). We will read in
// these configuration settings from command-line flags when the application starts.
type config struct {
	port int
	env  string
}

// Define the application struct to hold the dependencies for our HTTP handlers,
// helpers and middleware. At the moment, this only contains a copy of the config
// struct and a logger, but it will grow to include a lot more as our build progresses.
type application struct {
	config config
	logger *slog.Logger
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "Api server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development |staging|production)")
	flag.Parse()

	logger := slog.Default()

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Info(fmt.Sprintf("Starting %s server on %s", cfg.env, srv.Addr))
	err := srv.ListenAndServe()
	logger.Error(err.Error())
}
