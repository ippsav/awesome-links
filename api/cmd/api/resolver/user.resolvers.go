package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/ippsav/awesome-links/api/cmd/api/authentication"
	"github.com/ippsav/awesome-links/api/cmd/api/middleware"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/model"
)

func (r *mutationResolver) Register(ctx context.Context, createUserInput model.CreateUserInput) (*ent.User, error) {
	ginContext, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	session := middleware.GetSession(ginContext)
	us, err := r.controller.User.Create(ctx, createUserInput)
	if err != nil {
		return nil, err
	}
	accessToken, err := authentication.CreateToken(us.ID.String(), r.config.JWT.ExpirationTime, r.config.JWT.Secret)
	if err != nil {
		return nil, err
	}
	session.Set(middleware.CookieKey, accessToken)
	session.Save()
	return us, nil
}

func (r *mutationResolver) Login(ctx context.Context, loginUserInput model.LoginUserInput) (*ent.User, error) {
	ginContext, err := middleware.GinContextFromContext(ctx)
	if err != nil {
		return nil, err
	}
	session := middleware.GetSession(ginContext)
	us, err := r.controller.User.GetByEmail(ctx, loginUserInput)
	if err != nil {
		return nil, err
	}
	accessToken, err := authentication.CreateToken(us.ID.String(), r.config.JWT.ExpirationTime, r.config.JWT.Secret)
	if err != nil {
		return nil, err
	}
	session.Set(middleware.CookieKey, accessToken)
	session.Save()
	return us, nil
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
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
	return r.controller.User.GetByID(ctx, id)
}
