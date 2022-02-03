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
		accessToken := session.Get(CookieKey).(string)
		claims, err := authentication.ParseToken(accessToken, config.JWT.Secret)
		if err != nil {
			return
		}
		if claims.Valid() != nil {
			session.Delete(CookieKey)
		}
		c.Set("userID", claims.UserID)
	}
}
