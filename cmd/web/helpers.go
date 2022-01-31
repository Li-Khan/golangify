package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *Application) methodNotAllowed(w http.ResponseWriter) {
	app.clientError(w, http.StatusMethodNotAllowed)
}

func (app *Application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	tmpl, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("шаблон %s не существует", name))
		return
	}

	err := tmpl.Execute(w, td)
	if err != nil {
		app.serverError(w, err)
	}
}
