package main

import (
	"net/http"
	pb "project/API"
	"project/tools/log"
	"project/tools/setup"

	"google.golang.org/grpc/health"

	"github.com/gorilla/mux"
)

var conn, Ctx, cancel = setup.InitClient(":50051")
var ServiceForecast = pb.NewForecastServiceClient(conn)

func main() {
	defer conn.Close()
	defer cancel()

	health.NewServer()

	r := mux.NewRouter()

	r.HandleFunc("/api/auth", Auth)
	r.HandleFunc("/api/forecast/now", ForecastNow)
	r.HandleFunc("/api/forecast/history", ForecastHistory)

	log.Out("RUN HTTP ROUTER: localhost:8000")
	http.ListenAndServe(":8000", r)

}
