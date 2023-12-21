package app

import (
	"lesson/internal/api/controller"
	"net/http"
)

func route(c controller.Controller) http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", c.Healthz)
	return mux
}
