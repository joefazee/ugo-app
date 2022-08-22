package main

import "net/http"

func (a *application) get(path string, h http.HandlerFunc) {
	a.App.Routes.Get(path, h)
}

func (a *application) post(path string, h http.HandlerFunc) {
	a.App.Routes.Post(path, h)
}

func (a *application) use(m ...func(handler http.Handler) http.Handler) {
	a.App.Routes.Use(m...)
}
