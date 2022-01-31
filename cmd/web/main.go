package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	infoLog := log.New(os.Stdout, colorGreen+"INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, colorRed+"ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
