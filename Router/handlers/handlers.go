package handlers

import (
	"fmt"
	"net/http"
	log "project/tools/log"
)

// api/auth
func Auth(w http.ResponseWriter, r *http.Request) {
	log.Out("Auth")
	// Тут буде ваша логіка обробки /api/auth
	fmt.Fprintf(w, "Auth API")
}

// api/forecast/now
func ForecastNow(w http.ResponseWriter, r *http.Request) {
	log.Out("ForecastNow")
	// Тут буде ваша логіка обробки /api/forecast/now
	fmt.Fprintf(w, "Forecast Now API")
}

// api/forecast/history
func ForecastHistory(w http.ResponseWriter, r *http.Request) {
	log.Out("ForecastHistory")
	// Тут буде ваша логіка обробки /api/forecast/history
	fmt.Fprintf(w, "Forecast History API")
}
