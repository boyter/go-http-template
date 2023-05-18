package main

import (
	"changeme/common"
	"changeme/handlers"
	"changeme/service"
	"embed"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

//go:embed assets/static
var staticFiles embed.FS

func main() {
	config := common.NewConfig()
	log.Info().Str(common.UniqueCode, "").Interface("config", config).Msg("config")

	ser := service.NewService(config)

	switch config.LogLevel {
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	app, err := handlers.NewApplication(ser, staticFiles)
	if err != nil {
		log.Error().Str(common.UniqueCode, "5576633b").Err(err).Msg("error creating application")
		return
	}

	srv := &http.Server{
		Addr:    ":" + strconv.Itoa(config.HttpPort),
		Handler: app.Routes(),
	}

	log.Log().Str(common.UniqueCode, "3812c7e").Msg("starting server on :" + strconv.Itoa(config.HttpPort))
	err = srv.ListenAndServe()
	log.Error().Str(common.UniqueCode, "42aa9c1").Err(err).Msg("exiting server")
}
