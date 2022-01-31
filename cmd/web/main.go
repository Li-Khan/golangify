package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Li-Khan/golangify/pkg/models/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	dsn := flag.String("dsn", "snippets.db", "Название базы данных")
	flag.Parse()

	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	infoLog := log.New(os.Stdout, colorGreen+"INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, colorRed+"ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Println(err)
		return
	}
	defer db.Close()

	templateCache, err := newTemplateCache("./ui/html")
	if err != nil {
		errorLog.Println(err)
		return
	}

	app := Application{
		InfoLog:       infoLog,
		ErrorLog:      errorLog,
		snippets:      &sqlite.SnippetModel{DB: db},
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
