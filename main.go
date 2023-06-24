package main

import (
	grpcPayment "gobase/internal/payment/handler/grpc"
	restPayment "gobase/internal/payment/handler/rest"
	"gobase/logger"
	mgorm "gobase/packages/gorm"
	grpcServer "gobase/packages/grpc"
	restServer "gobase/packages/rest"
	mutils "gobase/utils"
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

	db, err := mgorm.NewInstance("anhnv:123456@tcp(localhost:3308)/healthnet?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}

	utils := mutils.NewInstance(db.GetDB(), l)

	// Grpc server
	grpcSV := grpcServer.NewInstance("", 9002)
	grpcPaymentHandler := grpcPayment.NewHandler(utils)
	grpcPaymentHandler.Attach(grpcSV.GetInstance())
	grpcSV.Serve()

	// Rest api server
	restPaymentHandler := restPayment.NewHandler(utils)
	restSV := restServer.NewInstance(restPaymentHandler, "", 8002)
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
