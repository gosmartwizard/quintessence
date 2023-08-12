package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type config struct {
	address   string
	staticDir string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config
}

func main() {

	var cfg config

	flag.StringVar(&cfg.address, "address", ":4949", "HTTP address to connect to")
	flag.StringVar(&cfg.staticDir, "static-directory", "./ui/static", "Path to static directory")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		config:   cfg,
	}

	server := &http.Server{
		Addr:     app.address,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Server listening on port %s", cfg.address)

	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
