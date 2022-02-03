package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/ippsav/awesome-links/api/cmd/api/authentication"
	"github.com/ippsav/awesome-links/api/cmd/api/config"
)

var CookieKey = "qid"

func AuthMiddleware(config *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		accessToken := session.Get(CookieKey)
		if accessToken == nil {
			return
		}
		claims, err := authentication.ParseToken(accessToken.(string), config.JWT.Secret)
		if err != nil {
			return
		}
		if claims.Valid() != nil {
			session.Delete(CookieKey)
			return
		}
		c.Set("userID", claims.UserID)
	}
}
