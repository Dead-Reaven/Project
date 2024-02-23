package setup

import (
	"context"
	"net"
	"project/tools/log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func InitClient(port string) (*grpc.ClientConn, context.Context, context.CancelFunc) {
	conn, err := grpc.Dial(port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {

		log.Error("problem with the server:" + err.Error())
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)

	return conn, ctx, cancel
}

type ServiceRegistrar interface {
	Register(*grpc.Server)
}

type RegisterFunc func(*grpc.Server)

func InitServer(port string, register RegisterFunc) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Error("tcp connection failed:" + err.Error())
	}
	println("listening at" + lis.Addr().String())

	s := grpc.NewServer(grpc.UnaryInterceptor(loggingInterceptor))

	register(s)

	if err := s.Serve(lis); err != nil {
		log.Error("grpc server failed:" + err.Error())
	}
}

func loggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Out(info.FullMethod)

	return handler(ctx, req)
}
