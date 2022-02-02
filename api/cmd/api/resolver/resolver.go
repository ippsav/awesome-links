package resolver

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/graph/generated"
	"github.com/ippsav/awesome-links/api/internal/adapter/controller"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	client     *ent.Client
	controller *controller.Controller
}

func NewSchema(client *ent.Client, controller *controller.Controller) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{Resolvers: &Resolver{client, controller}})
}
