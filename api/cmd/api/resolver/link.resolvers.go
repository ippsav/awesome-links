package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/cmd/api/middleware"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/model"
)

func (r *mutationResolver) CreateLink(ctx context.Context, createLinkInput model.CreateLinkInput) (*ent.Link, error) {
	ginContext, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	userID, ok := ginContext.Get("userID")
	if !ok {
		return nil, fmt.Errorf("something went wrong")
	}
	id, err := uuid.Parse(userID.(string))
	if err != nil {
		return nil, err
	}
	return r.controller.Link.Create(ctx, createLinkInput, id)
}

func (r *mutationResolver) BookmarkLink(ctx context.Context, id uuid.UUID) (*ent.Link, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Link(ctx context.Context, id uuid.UUID) (*ent.Link, error) {
	return r.controller.Link.GetByID(ctx, id)
}
