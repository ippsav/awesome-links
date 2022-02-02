package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, createLinkInput model.CreateLinkInput) (*ent.Link, error) {
	return r.controller.Link.Create(ctx, createLinkInput)
}

func (r *queryResolver) Link(ctx context.Context, id uuid.UUID) (*ent.Link, error) {
	return r.controller.Link.GetByID(ctx, id)
}
