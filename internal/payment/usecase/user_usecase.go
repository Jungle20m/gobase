package usecase

import (
	"context"
	"fmt"
	"gobase/internal/payment/model"
	mutils "gobase/utils"
)

type IUserRepo interface {
	GetUserByID(ctx context.Context, id int) (*model.User, error)
}

type userUC struct {
	utils mutils.IUtils
	repo  IUserRepo
}

func NewUserUC(utils mutils.IUtils, repo IUserRepo) *userUC {
	return &userUC{
		utils: utils,
		repo:  repo,
	}
}

func (u *userUC) GetUserByID(ctx context.Context, id int) {
	user, err := u.repo.GetUserByID(ctx, id)
	fmt.Println(user)
	fmt.Print(err)
}
