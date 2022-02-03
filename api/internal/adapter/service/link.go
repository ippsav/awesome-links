package service

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/internal/entity"
)

type Link interface {
	Create(ctx context.Context, title, description, imageURL, URL string, ownerID uuid.UUID) (*entity.Link, error)
	GetByID(ctx context.Context, id uuid.UUID) (*entity.Link, error)
	GetUserLinks(ctx context.Context, ownerID uuid.UUID) ([]*entity.Link, error)
}

type linkService struct {
	Link
	Cloudinary
}

func NewLinkService(lr Link, cr Cloudinary) *linkService {
	return &linkService{lr, cr}
}

func (ls *linkService) Create(ctx context.Context, inp entity.CreateLinkInput, ownerID uuid.UUID) (*entity.Link, error) {
	if inp.Image == nil {
		return ls.Link.Create(ctx, inp.Title, inp.Description, "", inp.URL, ownerID)
	}
	imageUrl, err := ls.Cloudinary.UploadImage(ctx, inp.Image.File)
	if err != nil {
		return nil, err
	}
	return ls.Link.Create(ctx, inp.Title, inp.Description, imageUrl, inp.URL, ownerID)
}
func (ls *linkService) GetUserLinks(ctx context.Context, ownerID uuid.UUID) ([]*entity.Link, error) {
	return ls.Link.GetUserLinks(ctx, ownerID)
}
