package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	corsCfg := cors.Config{
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodOptions},
		AllowHeaders:     []string{"Authorization"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// Only allow Shitcamp client domains or localhost
			switch origin {
			case "https://shitcamp.github.io",
				"https://shitcamp-unofficial.github.io":
				return true

			default:
				return strings.HasPrefix(origin, "http://localhost")
			}
		},
		MaxAge: 12 * time.Hour,
	}

	return cors.New(corsCfg)
}
