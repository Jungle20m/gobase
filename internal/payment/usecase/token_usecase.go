package usecase

import (
	"context"
	"fmt"
	"gobase/internal/payment/model"
	mutils "gobase/utils"
)

type ITokenRepo interface {
	GetTokenByID(ctx context.Context, id int) (*model.Token, error)
}

type tokenUC struct {
	utils mutils.IUtils
	repo  ITokenRepo
}

func NewTokenUC(utils mutils.IUtils, repo ITokenRepo) *tokenUC {
	return &tokenUC{
		utils: utils,
		repo:  repo,
	}
}

func (u *tokenUC) GetTokenByID(ctx context.Context, id int) {
	user, err := u.repo.GetTokenByID(ctx, id)
	fmt.Println(user)
	fmt.Print(err)
}
