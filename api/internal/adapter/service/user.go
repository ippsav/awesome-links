package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type User interface {
	Create(ctx context.Context, user *entity.User) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type userService struct {
	User
}

func NewUserService(ur User) *userService {
	return &userService{ur}
}
