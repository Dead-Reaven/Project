package main

import (
	"fmt"
	"net/http"
	pb "project/API"
	log "project/tools/log"
)

// api/auth
func Auth(w http.ResponseWriter, r *http.Request) {
	log.Out("Auth")
	fmt.Fprintf(w, "Auth API")
}

// api/forecast/now
func ForecastNow(w http.ResponseWriter, r *http.Request) {

	log.Out("ForecastNow")
	res, err := ServiceForecast.Now(Ctx, &pb.NowReq{})
	if err != nil {
		log.Error(err.Error())
		fmt.Fprintf(w, "Forecast Err:"+err.Error())
	} else {
		log.Success(res.Message)
		fmt.Fprintf(w, "Forecast Now API")
	}

}

// api/forecast/history
func ForecastHistory(w http.ResponseWriter, r *http.Request) {
	log.Out("ForecastHistory")
	fmt.Fprintf(w, "Forecast History API")
}
