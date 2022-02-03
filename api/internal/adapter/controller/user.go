package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type User interface {
	Create(ctx context.Context, inp entity.CreateUserInput) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
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

func (uc *userController) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return uc.User.GetByID(ctx, id)
}
