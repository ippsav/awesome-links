package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/generated"
	"github.com/ippsav/awesome-links/api/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, createUserInput model.CreateUserInput) (*ent.User, error) {
	us, err := r.controller.User.Create(ctx, createUserInput)
	fmt.Println(err)
	return us, nil
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Bookmarks(ctx context.Context, obj *ent.User) ([]*ent.Link, error) {
	panic(fmt.Errorf("not implemented"))
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
