package grpc

import (
	"context"
	"gobase/internal/payment/handler/grpc/protoc"
	"gobase/internal/payment/repository"
	"gobase/internal/payment/usecase"
)

func (h *handler) Ping(ctx context.Context, in *protoc.Request) (*protoc.Response, error) {

	repo := repository.NewInstance(repository.WithGorm(h.utils.GetDB()))
	uc := usecase.NewUserUC(h.utils, repo)

	uc.GetUserByID(ctx, 1)

	return &protoc.Response{Message: "pong"}, nil
}
