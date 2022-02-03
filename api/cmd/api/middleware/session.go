package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

var sessionCtxKey = "session"

func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		c.Set(sessionCtxKey, &session)
	}
}

func GetSession(c *gin.Context) *sessions.Session {
	session, _ := c.Get(sessionCtxKey)
	return session.(*sessions.Session)
}
