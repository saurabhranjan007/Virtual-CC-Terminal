package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Entry point of client interface

const version = "1.0.0"
const cssVersion = "1" // to downgrade/upgrade the css

// to hold info about the application
type config struct {
	port int
	env  string
	api  string
	db   struct {
		dsn string // Data Source Name
	}
	stripe struct {
		secret string // prrivate key
		key    string // publishable key
	}
}

// Receiver struct
type application struct {
	config        config
	infoLog       *log.Logger
	errorLog      *log.Logger
	templateCache map[string]*template.Template
	version       string
}

// Web server function
func (app *application) Serve() error {

	// Using the http package (creating web server configuratiton)
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%d", app.config.port),
		Handler:           app.routes(), // Route handler
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}

	app.infoLog.Printf("Starting HTTP Server in %s mode on Port %d", app.config.env, app.config.port)

	return srv.ListenAndServe()

}

func main() {

	// Create a var of type conifg
	var cfg config

	// Some values of the variable cfg will come from the command line flag
	// Default - port 4000 and environemnt - development
	flag.IntVar(&cfg.port, "port", 4000, "Server port to listen")
	flag.StringVar(&cfg.env, "env", "development", "Application environment {development/production}")
	flag.StringVar(&cfg.api, "api", "http://localhost:4001", "URL to api}")

	// parse the command flags
	flag.Parse()

	// Reading the strip data from .env file
	cfg.stripe.key = os.Getenv("STRIPE_KEY")
	cfg.stripe.secret = os.Getenv("STRIPE_SECRET")

	// Setting up logs
	infoLog := log.New(os.Stdout, "INFO/t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR/t", log.Ldate|log.Ltime|log.Lshortfile)

	// Creating a map for the values
	tc := make(map[string]*template.Template)

	// Application variable (ref to application)
	app := &application{
		config:        cfg,
		infoLog:       infoLog,
		errorLog:      errorLog,
		templateCache: tc,
		version:       version,
	}

	// Calling the Web Server
	err := app.Serve()
	if err != nil {
		app.errorLog.Println(err)
		log.Fatal(err)
	}
}
