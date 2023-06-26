// Filename: cmd/api/main.go

package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

// The application version number
const version string = "1.0.0"

// The configuration settings
type config struct {
	port int
	env  string
}

// Dependency injection
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	// read in the flags that are needed to populate our config
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development | staging | production)")
	flag.Parse()

	// Create the logger
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	// Create an instance of our application struct
	app := &application{
		config: cfg,
		logger: logger,
	}

	// Create the servemux
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr: fmt.Sprintf(":%d", cfg.port),
		// Use the custom route handler that we create in routes.go
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Start our server
	logger.Printf("Starting %s on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}
