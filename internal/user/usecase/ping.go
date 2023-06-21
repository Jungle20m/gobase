package usecase

import "fmt"

type IUseCase interface {
	Ping()
}

type userCase struct {
}

func NewUseCase() *userCase {
	return &userCase{}
}

func (u *userCase) Ping() {
	fmt.Println("Ping...")
}
