package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type config struct {
	address      string
	staticDir    string
	dbConnString string
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
	flag.StringVar(&cfg.dbConnString, "db-conn-string", "quint:quint4949@tcp(0.0.0.0:3306)/essencebox?parseTime=true", "MySQL DB connection string")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(cfg.dbConnString)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

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

	err = server.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(connString string) (*sql.DB, error) {
	db, err := sql.Open("mysql", connString)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
