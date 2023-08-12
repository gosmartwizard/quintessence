package main

import (
	"flag"
	"log"
	"net/http"
)

type config struct {
	address   string
	staticDir string
}

func main() {

	var cfg config

	flag.StringVar(&cfg.address, "address", ":4949", "HTTP address to connect to")
	flag.StringVar(&cfg.staticDir, "static-directory", "./ui/static", "Path to static directory")

	flag.Parse()

	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir(cfg.staticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	mux.HandleFunc("/", HomeHandler)

	mux.HandleFunc("/quint/create", QuintCreateHandler)
	mux.HandleFunc("/quint/update", QuintUpdateHandler)
	mux.HandleFunc("/quint/delete", QuintDeleteHandler)
	mux.HandleFunc("/quint/list", QuintListHandler)
	mux.HandleFunc("/quint/get", QuintGetHandler)

	log.Printf("Server listening on port %s", cfg.address)

	err := http.ListenAndServe(cfg.address, mux)
	log.Fatal(err)
}
