package main

import (
	"github.com/joefazee/ladiwork/data"
	"github.com/joefazee/ladiwork/handlers"
	"github.com/joefazee/ladiwork/middleware"
	"github.com/joefazee/ugo"
)

type application struct {
	App        *ugo.Ugo
	Handler    *handlers.Handler
	Models     data.Models
	Middleware *middleware.Middleware
}

func main() {

	l := initApplication()

	l.App.ListenAndServe()

}
