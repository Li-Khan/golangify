package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Li-Khan/golangify/config"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func main() {
	addr := flag.String("addr", ":4000", "HTTP Network address")
	flag.Parse()

	colorRed := "\033[31m"
	colorGreen := "\033[32m"

	infoLog := log.New(os.Stdout, colorGreen+"INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, colorRed+"ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &config.Application{
		InfoLog:  infoLog,
		ErrorLog: errorLog,
	}

	mux := http.NewServeMux()
	// mux.HandleFunc("/", app.home)
	mux.Handle("/", Home(app))
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileServer := http.FileServer(neuteredFileSystem{http.Dir("./ui/static")})
	mux.Handle("/static", http.NotFoundHandler())
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}

func (nfs neuteredFileSystem) Open(path string) (http.File, error) {
	f, err := nfs.fs.Open(path)
	if err != nil {
		return nil, err
	}

	s, err := f.Stat()
	if s.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := nfs.fs.Open(index); err != nil {
			closeErr := f.Close()
			if closeErr != nil {
				return nil, closeErr
			}
			return nil, err
		}
	}
	return f, nil
}
