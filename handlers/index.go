package handlers

import (
	"changeme/common"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func (app *Application) Index(w http.ResponseWriter, r *http.Request) {
	log.Info().Str(common.UniqueCode, "93c9683d").Str("ip", GetIP(r)).Msg("Index")

	query := strings.TrimSpace(r.URL.Query().Get("q"))

	err := indexTemplate.Execute(w, templateData{
		SearchTerm: query,
	})
	if err != nil {
		log.Error().Str(common.UniqueCode, "f6fb63c").Str("ip", GetIP(r)).Err(err).Msg("error executing template")
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
