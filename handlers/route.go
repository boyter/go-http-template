package handlers

import (
	"changeme/common"
	"changeme/service"
	"embed"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog/log"
	"io/fs"
	"net/http"
)

type Application struct {
	Service     *service.Service
	StaticFiles embed.FS
}

func NewApplication(service *service.Service, staticFiles embed.FS) (Application, error) {
	application := Application{
		Service:     service,
		StaticFiles: staticFiles,
	}
	err := application.ParseTemplates()
	return application, err
}

func (app *Application) Routes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// JSON routes
	router.HandleFunc("/health-check/", app.HealthCheck).Methods("GET")

	// Regular routes
	router.HandleFunc("/", app.Index).Methods("GET")

	// Setup to serve files from the supplied directory
	staticFS := fs.FS(app.StaticFiles)
	// strip off the location such that we route correctly
	staticContent, err := fs.Sub(staticFS, "assets/static")
	if err != nil {
		log.Fatal().Str(common.UniqueCode, "f8819e2a").Err(err).Msg("error with static content")
	}
	f := http.FileServer(http.FS(staticContent))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static", f))

	return router
}
