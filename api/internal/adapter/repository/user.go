package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *userRepository {
	return &userRepository{client}
}

func (ur *userRepository) Create(ctx context.Context, email, password, imageURL string) (*entity.User, error) {
	user, err := ur.client.User.Create().SetEmail(email).SetPassword(password).SetImageURL(imageURL).Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (ur *userRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := ur.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
