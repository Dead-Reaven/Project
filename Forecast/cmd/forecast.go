package main

import (
	"context"
	"errors"
	pb "project/API"
	"project/tools/log"
	"project/tools/setup"

	"google.golang.org/genproto/googleapis/rpc/status"
	"google.golang.org/grpc"

	"google.golang.org/grpc/codes"
)

type ForecastService struct {
	pb.UnimplementedForecastServiceServer
}

func Register(s *grpc.Server) {
	pb.RegisterForecastServiceServer(s, &ForecastService{})
}

func main() {
	log.Out("Run Forecast micro-service")
	setup.InitServer(":50051", Register)
}

func (f *ForecastService) Now(ctx context.Context, req *pb.NowReq) (*status.Status, error) {
	return newResponse("forecast/now", codes.OK)
}

func newResponse(msg string, code codes.Code, e ...string) (*status.Status, error) {
	s := &status.Status{
		Message: msg,
		Code:    int32(code),
	}

	if len(e) > 0 {
		log.Error("Err:" + msg)
		return nil, errors.New(e[0]) // Використовуємо перший елемент зрізу
	}

	log.Success("Success: " + msg)
	return s, nil
}
