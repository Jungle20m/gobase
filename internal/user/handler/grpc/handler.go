package grpc

import (
	"context"
	"gobase/common"
	"gobase/internal/user/handler/grpc/protoc"
)

type userHandler struct {
	protoc.UnimplementedUserServer
	common.IAppContext
}

func NewUserHandler() *userHandler {
	return &userHandler{
		UnimplementedUserServer: protoc.UnimplementedUserServer{},
		IAppContext:             nil,
	}
}

func (h *userHandler) Ping(ctx context.Context, in *protoc.Request) (*protoc.Response, error) {
	return &protoc.Response{Message: "pong"}, nil
}
