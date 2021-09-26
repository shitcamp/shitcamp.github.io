package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth(authAccounts gin.Accounts, exceptions ...string) gin.HandlerFunc {
	auth := gin.BasicAuth(authAccounts)

	return func(c *gin.Context) {
		path := c.Request.URL.Path
		for _, e := range exceptions {
			if path == e {
				return
			}
		}

		auth(c)
	}
}
