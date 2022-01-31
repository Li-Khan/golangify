package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/Li-Khan/golangify/config"
)

func Home(app *config.Application) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		files := []string{
			"./ui/html/home.page.tmpl",
			"./ui/html/base.layout.tmpl",
			"./ui/html/footer.partial.tmpl",
		}

		tmpl, err := template.ParseFiles(files...)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		err = tmpl.Execute(w, nil)
		if err != nil {
			app.ErrorLog.Println(err.Error())
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	}
}

// func (app *config.Application) home(w http.ResponseWriter, r *http.Request) {
// 	if r.URL.Path != "/" {
// 		http.NotFound(w, r)
// 		return
// 	}

// 	files := []string{
// 		"./ui/html/home.page.tmpl",
// 		"./ui/html/base.layout.tmpl",
// 		"./ui/html/footer.partial.tmpl",
// 	}

// 	tmpl, err := template.ParseFiles(files...)
// 	if err != nil {
// 		app.errorLog.Println(err.Error())
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		app.errorLog.Println(err.Error())
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 	}
// }

func showSnippet(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Отображение выбранной заметки с ID %d", id)
}

func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Метод запрещен!", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Форма для создания новой заметки..."))
}
