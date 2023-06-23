package main

import (
	grpcServer "gobase/grpc"
	grpcHandler "gobase/internal/user/handler/grpc"
	"gobase/internal/user/handler/grpc/protoc"
	restHandler "gobase/internal/user/handler/rest"
	"gobase/logger"
	restServer "gobase/rest"
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
	l.Infof("hello %v %v", "world", "hi")
	l.Debug("hello world")
	l.Debugf("hello ", "world")
	l.Warn("hello world")
	l.Warnf("hello ", "world")
	l.Error("hello world")
	l.Errorf("hello ", "world")

	// Grpc server
	userGrpcHandler := grpcHandler.NewUserHandler()
	grpcSV := grpcServer.NewServer()
	protoc.RegisterUserServer(grpcSV.Instance(), userGrpcHandler)
	grpcSV.Serve()

	// Rest api server
	userRestHandler := restHandler.NewUserHandler()
	restSV := restServer.NewServer(userRestHandler, "", 8002)
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
