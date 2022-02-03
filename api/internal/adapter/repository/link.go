package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/ent/link"
	"github.com/ippsav/awesome-links/api/ent/user"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type linkRepository struct {
	client *ent.Client
}

func NewLinkRepository(client *ent.Client) *linkRepository {
	return &linkRepository{client}
}

func (lr *linkRepository) Create(ctx context.Context, title, description, imageURL, URL string, ownerID uuid.UUID) (*entity.Link, error) {
	link, err := lr.client.Link.Create().SetTitle(title).SetDescription(description).SetImageURL(imageURL).SetURL(URL).SetOwnerID(ownerID).Save(ctx)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (lr *linkRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error) {
	link, err := lr.client.Link.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return link, nil
}

func (lr *linkRepository) GetUserLinks(ctx context.Context, ownerID uuid.UUID) ([]*entity.Link, error) {
	links, err := lr.client.Link.Query().Where(link.HasOwnerWith(user.ID(ownerID))).All(ctx)
	if err != nil {
		return nil, err
	}
	return links, nil
}
