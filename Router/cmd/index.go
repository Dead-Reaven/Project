package main

import (
	"net/http"
	"project/Router/handlers"
	log "project/tools/log"

	"github.com/gorilla/mux"
)

func main() {
	log.Out("RUN HTTP ROUTER: localhost:8000")

	r := mux.NewRouter()

	r.HandleFunc("/api/auth", handlers.Auth)
	r.HandleFunc("/api/forecast/now", handlers.ForecastNow)
	r.HandleFunc("/api/forecast/history", handlers.ForecastHistory)

	http.ListenAndServe(":8000", r)

}
