package main

import (
	"github.com/joefazee/ladiwork/middleware"
	"log"
	"os"

	"github.com/joefazee/ladiwork/data"
	"github.com/joefazee/ladiwork/handlers"
	"github.com/joefazee/ugo"
)

func initApplication() *application {

	path, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	// init ugo
	core := &ugo.Ugo{AppName: "LadiWork"}
	err = core.New(path)
	if err != nil {
		log.Fatal(err)
	}

	appMiddleware := &middleware.Middleware{
		App: core,
	}

	httpHandlers := &handlers.Handler{
		App: core,
	}

	app := &application{
		App:        core,
		Handler:    httpHandlers,
		Middleware: appMiddleware,
	}

	app.App.Routes = app.routes()
	app.Models = data.New(app.App.DB.Pool)
	app.Middleware.Models = app.Models
	httpHandlers.Models = app.Models

	return app

}
