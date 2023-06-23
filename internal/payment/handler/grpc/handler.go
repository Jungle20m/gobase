package grpc

import (
	"gobase/internal/payment/handler/grpc/protoc"
	mutils "gobase/utils"
)

type handler struct {
	protoc.UnimplementedUserServer
	utils mutils.IUtils
}

func NewHandler(utils mutils.IUtils) *handler {
	return &handler{
		UnimplementedUserServer: protoc.UnimplementedUserServer{},
		utils:                   utils,
	}
}
