package grpc

import (
	"gobase/internal/payment/handler/grpc/protoc"
	mutils "gobase/utils"
	"google.golang.org/grpc"
)

type handler struct {
	protoc.UnimplementedUserServer
	protoc.UnimplementedTokenServer
	utils mutils.IUtils
}

func NewHandler(utils mutils.IUtils) *handler {
	return &handler{
		UnimplementedUserServer: protoc.UnimplementedUserServer{},
		utils:                   utils,
	}
}

func (h *handler) Attach(server *grpc.Server) {
	protoc.RegisterUserServer(server, h)
	protoc.RegisterTokenServer(server, h)
}
