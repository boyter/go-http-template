package handlers

import (
	"net/http"
)

// GetIP attempts to use the X-FORWARDED-FOR http header for code behind proxies and load balancers
// (such as on hosts like Heroku) while falling back to the RemoteAddr if the header isn't found.
// NB when behind a LB can return value like so "1.143.90.29, 127.0.0.1"
func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}
