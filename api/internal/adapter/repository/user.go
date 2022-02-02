package repository

import (
	"context"

	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type userRepository struct {
	client *ent.Client
}

func NewUserRepository(client *ent.Client) *userRepository {
	return &userRepository{client}
}

func (ur *userRepository) Create(ctx context.Context, inp entity.CreateUserInput) (*entity.User, error) {
	user, err := ur.client.User.Create().SetEmail(inp.Email).SetPassword(inp.Password).SetImageURL("").Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}
