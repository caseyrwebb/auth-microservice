package app

import (
	"github.com/caseyrwebb/auth-microservice/app/data"
	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     data.GoDB
}

func New() *App {
	a := &App{
		Router: mux.NewRouter(),
	}

	return a
}
