package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/ippsav/awesome-links/api/cmd/api/resolver"
	"github.com/ippsav/awesome-links/api/ent"
	"github.com/ippsav/awesome-links/api/internal/adapter/controller"
)

func GraphqlHandler(client *ent.Client, controller *controller.Controller) gin.HandlerFunc {

	return func(c *gin.Context) {
		h := handler.NewDefaultServer(resolver.NewSchema(client, controller, c.Request))
		h.ServeHTTP(c.Writer, c.Request)
	}
}
