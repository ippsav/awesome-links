package controller

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type Link interface {
	Create(ctx context.Context, inp entity.CreateLinkInput) (*entity.Link, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error)
}

type linkController struct {
	Link
}

func NewLinkController(lr Link) *linkController {
	return &linkController{
		Link: lr,
	}
}

func (lc *linkController) Create(ctx context.Context, inp entity.CreateLinkInput) (*entity.Link, error) {
	return lc.Link.Create(ctx, inp)
}

func (lc *linkController) GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error) {
	return lc.Link.GetByID(ctx, id)
}
