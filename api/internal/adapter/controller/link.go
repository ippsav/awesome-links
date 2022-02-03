package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type Link interface {
	Create(ctx context.Context, inp entity.CreateLinkInput, ownerID uuid.UUID) (*entity.Link, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error)
	GetUserLinks(ctx context.Context, ownerID uuid.UUID) ([]*entity.Link, error)
}

type linkController struct {
	Link
}

func NewLinkController(lr Link) *linkController {
	return &linkController{
		Link: lr,
	}
}

func (lc *linkController) Create(ctx context.Context, inp entity.CreateLinkInput, ownerID uuid.UUID) (*entity.Link, error) {
	return lc.Link.Create(ctx, inp, ownerID)
}

func (lc *linkController) GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error) {
	return lc.Link.GetByID(ctx, id)
}

func (lc *linkController) GetUserLinks(ctx context.Context, ownerID uuid.UUID) ([]*entity.Link, error) {
	return lc.Link.GetUserLinks(ctx, ownerID)
}
