package main

import (
	"net/http"

	"github.com/beata-krasnopolska/bookingsGo/pkg/config"
	"github.com/beata-krasnopolska/bookingsGo/pkg/handlers"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/v5/middleware"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter() //create multiplex
	mux.Use(middleware.Recoverer)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
