package handlers

import (
	"changeme/common"
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
	"time"
)

const (
	jsonContentType = "application/json; charset=utf-8"
	jsonIndent      = "    "
)

type Timing struct {
	TimeMillis int64  `json:"timeMillis"`
	Source     string `json:"source"`
}

type HealthCheckResult struct {
	Success  bool                `json:"success"`
	Messages []string            `json:"messages"`
	Time     time.Time           `json:"time"`
	Timing   []Timing            `json:"timing"`
	Response HealthCheckResponse `json:"response"`
}

type HealthCheckResponse struct {
	IpAddress string `json:"ipAddress"`
	MemUsage  string `json:"memUsage"`
}

func (app *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Info().Str(common.UniqueCode, "b10037c0").Str("ip", GetIP(r)).Msg("HealthCheck")
	start := common.MakeTimestampMilli()

	t, _ := json.MarshalIndent(HealthCheckResult{
		Success:  true,
		Messages: []string{},
		Time:     time.Now().UTC(),
		Timing: []Timing{
			{
				Source:     "HealthCheck",
				TimeMillis: common.MakeTimestampMilli() - start,
			},
		},
		Response: HealthCheckResponse{
			IpAddress: GetIP(r),
			MemUsage:  common.MemUsage(),
		},
	}, "", jsonIndent)

	w.Header().Set("Content-Type", jsonContentType)
	_, _ = fmt.Fprint(w, string(t))
}
