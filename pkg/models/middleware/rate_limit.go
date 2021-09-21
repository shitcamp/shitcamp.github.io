package middleware

import (
	"fmt"
	"net/http"
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/common"

	logger "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"

	"github.com/sethvargo/go-limiter/memorystore"
)

func NewRateLimiter(tokens uint64, interval time.Duration) gin.HandlerFunc {
	store, err := memorystore.New(&memorystore.Config{
		Tokens:   tokens,
		Interval: interval,
	})
	if err != nil {
		logger.WithFields(logger.Fields{
			"tokens":   tokens,
			"interval": interval,
		}).WithError(err).Fatal("middleware_NewRateLimiter_error")
	}

	return func(c *gin.Context) {
		key := c.ClientIP()

		_, _, _, ok, err := store.Take(c, key)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, common.Response{
				Error: fmt.Sprintf("rate_limiter_error: %v", err),
			})
			return
		}

		if !ok {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, common.Response{Error: "rate_limit_reached"})
			return
		}
	}
}
