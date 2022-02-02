package repository

import (
	"context"
	"io"

	cloudinary "github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/cmd/api/config"
)

type cloudinaryRepository struct {
	storage *cloudinary.Cloudinary
}

func NewCloudinaryRepository(cfg *config.Config) (*cloudinaryRepository, error) {
	cld, err := cloudinary.NewFromParams(cfg.Cloudinary.CloudName, cfg.Cloudinary.ApiKey, cfg.Cloudinary.ApiSecret)
	if err != nil {
		return nil, err
	}
	return &cloudinaryRepository{storage: cld}, nil
}

func (cr *cloudinaryRepository) UploadImage(ctx context.Context, file io.Reader, imageID uuid.UUID) (string, error) {
	res, err := cr.storage.Upload.Upload(ctx, file, uploader.UploadParams{PublicID: imageID.String()})
	if err != nil {
		return "", err
	}
	return res.SecureURL, nil
}
