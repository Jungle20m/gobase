package main

import (
	grpcHandler "gobase/internal/user/handler/grpc"
	"gobase/internal/user/handler/grpc/protoc"
	restHandler "gobase/internal/user/handler/rest"
	"gobase/logger"
	grpcServer "gobase/packages/grpc"
	restServer "gobase/packages/rest"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	l, err := logger.NewInstance(
		logger.WithFormatter(logger.TextFormatter),
		//logger.WithFileOutput("log", "gobase.log"),
		logger.WithLevel(logger.TraceLevel),
	)
	if err != nil {
		log.Fatal(err)
	}

	l.Info("hello world")

	// Grpc server
	userGrpcHandler := grpcHandler.NewUserHandler()
	grpcSV := grpcServer.NewInstance("", 9002)
	protoc.RegisterUserServer(grpcSV.GetInstance(), userGrpcHandler)
	grpcSV.Serve()

	// Rest api server
	userRestHandler := restHandler.NewUserHandler()
	restSV := restServer.NewInstance(userRestHandler, "", 8002)
	restSV.Serve()

	// Graceful shutdown
	idleConnectionsClosed := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

		<-c

		grpcSV.Terminal()
		restSV.Terminal()

		close(idleConnectionsClosed)
	}()

	<-idleConnectionsClosed
}
