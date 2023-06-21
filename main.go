package main

import (
	grpcServer "gobase/grpc"
	grpcHandler "gobase/internal/user/handler/grpc"
	"gobase/internal/user/handler/grpc/protoc"
)

func main() {

	userGrpcHandler := grpcHandler.NewUserHandler()

	grpcServer := grpcServer.NewServer()

	protoc.RegisterUserServer(grpcServer.Instance(), userGrpcHandler)

	grpcServer.Serve()

}
