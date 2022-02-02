package service

import (
	"context"
	"io"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type User interface {
	Create(ctx context.Context, email, password, imageURL string) (*entity.User, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error)
}

type Cloudinary interface {
	UploadImage(ctx context.Context, file io.Reader) (string, error)
}

type userService struct {
	User
	Cloudinary
}

func NewUserService(ur User, c Cloudinary) *userService {
	return &userService{ur, c}
}

func (us *userService) Create(ctx context.Context, inp entity.CreateUserInput) (*entity.User, error) {
	imageUrl, err := us.Cloudinary.UploadImage(ctx, inp.Image.File)
	if err != nil {
		return nil, err
	}
	return us.User.Create(ctx, inp.Email, inp.Password, imageUrl)
}

func (us *userService) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	return us.User.GetByID(ctx, id)
}
