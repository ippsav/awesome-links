package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/generated"
	"github.com/ippsav/awesome-links/api/graph/model"
)

func (r *linkResolver) Owner(ctx context.Context, obj *ent.Link) (*ent.User, error) {
	return obj.QueryOwner().First(ctx)
}

func (r *mutationResolver) CreateLink(ctx context.Context, createLinkInput model.CreateLinkInput) (*ent.Link, error) {
	// placeholder
	id := uuid.New()
	return r.controller.Link.Create(ctx, createLinkInput, id)
}

func (r *queryResolver) Link(ctx context.Context, id uuid.UUID) (*ent.Link, error) {
	return r.controller.Link.GetByID(ctx, id)
}

// Link returns generated.LinkResolver implementation.
func (r *Resolver) Link() generated.LinkResolver { return &linkResolver{r} }

type linkResolver struct{ *Resolver }
