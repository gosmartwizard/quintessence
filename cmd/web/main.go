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
	}

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", app.HomeHandler)

	mux.HandleFunc("/quint/create", app.QuintCreateHandler)
	mux.HandleFunc("/quint/update", app.QuintUpdateHandler)
	mux.HandleFunc("/quint/delete", app.QuintDeleteHandler)
	mux.HandleFunc("/quint/list", app.QuintListHandler)
	mux.HandleFunc("/quint/get", app.QuintGetHandler)

	server := &http.Server{
		Addr:     cfg.address,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Server listening on port %s", cfg.address)

	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
