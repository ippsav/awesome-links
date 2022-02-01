package handler

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/gin-gonic/gin"
	"github.com/ippsav/awesome-links/api/cmd/api/resolver"
	"github.com/ippsav/awesome-links/api/ent"
)

func GraphqlHandler(client *ent.Client) gin.HandlerFunc {
	h := handler.NewDefaultServer(resolver.NewSchema(client))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}
