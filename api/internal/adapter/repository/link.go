package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type linkRepository struct {
	client *ent.Client
}

func NewLinkRepository(client *ent.Client) *linkRepository {
	return &linkRepository{client}
}

func (lr *linkRepository) Create(ctx context.Context, inp entity.CreateLinkInput) (*entity.Link, error) {
	link, err := lr.client.Link.Create().SetTitle(inp.Title).SetDescription(inp.Description).SetImageURL(inp.ImageURL).SetURL(inp.URL).SetOwnerID(inp.ID).Save(ctx)
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
