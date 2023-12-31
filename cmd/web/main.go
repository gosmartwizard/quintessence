package main

import (
	"crypto/tls"
	"database/sql"
	"flag"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/mysqlstore"
	"github.com/alexedwards/scs/v2"
	"github.com/go-playground/form/v4"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gosmartwizard/quintessence/internal/models"
)

type config struct {
	address      string
	staticDir    string
	dbConnString string
	debug        bool
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	config
	quints         *models.QuintModel
	users          *models.UserModel
	templateCache  map[string]*template.Template
	formDecoder    *form.Decoder
	sessionManager *scs.SessionManager
}

func main() {

	var cfg config

	flag.StringVar(&cfg.address, "address", ":4949", "HTTP address to connect to")
	flag.StringVar(&cfg.staticDir, "static-directory", "./ui/static", "Path to static directory")
	flag.StringVar(&cfg.dbConnString, "db-conn-string", "quint:quint4949@tcp(0.0.0.0:3306)/essencebox?parseTime=true", "MySQL DB connection string")
	flag.BoolVar(&cfg.debug, "debug", false, "Debug Mode to print debug information")

	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)

	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(cfg.dbConnString)

	if err != nil {
		errorLog.Fatal(err)
	}

	defer db.Close()

	templateCache, err := newTemplateCache()
	if err != nil {
		errorLog.Fatal(err)
	}

	formDecoder := form.NewDecoder()

	sessionManager := scs.New()
	sessionManager.Store = mysqlstore.New(db)
	sessionManager.Lifetime = 12 * time.Hour
	sessionManager.Cookie.Secure = true

	app := &application{
		errorLog:       errorLog,
		infoLog:        infoLog,
		config:         cfg,
		quints:         &models.QuintModel{DB: db},
		users:          &models.UserModel{DB: db},
		templateCache:  templateCache,
		formDecoder:    formDecoder,
		sessionManager: sessionManager,
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	server := &http.Server{
		Addr:         app.address,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Server listening on port %s", cfg.address)

	//err = server.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")

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
