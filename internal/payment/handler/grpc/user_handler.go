package grpc

import (
	"context"
	"gobase/internal/payment/handler/grpc/protoc"
	"gobase/internal/payment/repository"
	"gobase/internal/payment/usecase"
)

func (h *handler) GetUserByID(ctx context.Context, in *protoc.GetUserRequest) (*protoc.GetUserResponse, error) {

	repo := repository.NewInstance(repository.WithGorm(h.utils.GetDB()))
	uc := usecase.NewUserUC(h.utils, repo)

	uc.GetUserByID(ctx, int(in.GetId()))

	return &protoc.GetUserResponse{Message: "success"}, nil
}
