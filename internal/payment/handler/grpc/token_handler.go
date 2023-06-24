package grpc

import (
	"context"
	"gobase/internal/payment/handler/grpc/protoc"
	"gobase/internal/payment/repository"
	"gobase/internal/payment/usecase"
)

func (h *handler) GetTokenByID(ctx context.Context, in *protoc.GetTokenRequest) (*protoc.GetTokenResponse, error) {

	repo := repository.NewInstance(repository.WithGorm(h.utils.GetDB()))
	uc := usecase.NewTokenUC(h.utils, repo)

	uc.GetTokenByID(ctx, int(in.GetId()))

	return &protoc.GetTokenResponse{Message: "success"}, nil
}
