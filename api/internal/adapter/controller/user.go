package controller

import (
	"context"

	"github.com/ippsav/awesome-links/api/internal/entity"
)

type User interface {
	Create(ctx context.Context, inp entity.CreateUserInput) (*entity.User, error)
}

type userController struct {
	User
}

func NewUserController(ur User) *userController {
	return &userController{User: ur}
}

func (uc *userController) Create(ctx context.Context, inp entity.CreateUserInput) (*entity.User, error) {
	return uc.User.Create(ctx, inp)
}
