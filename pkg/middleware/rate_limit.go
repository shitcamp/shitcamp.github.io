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

func RateLimit(tokens uint64, interval time.Duration) gin.HandlerFunc {
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
	
		// TODO: key does not really represent "users", so rate limits should not be configured for the user level currently.

		// {"addr":"10.244.11.194:53554","c_ip":"10.244.11.194","hdr":"110.54.181.116,162.158.178.74","host":"shitcamp-api-2xvip.ondigitalocean.app","level":"info","msg":"req_ip","time":"2021-09-29T13:42:42Z"}
		// {"addr":"10.244.11.194:52406","c_ip":"10.244.11.194","hdr":"84.155.82.89,162.158.92.161","host":"shitcamp-api-2xvip.ondigitalocean.app","level":"info","msg":"req_ip","time":"2021-09-29T13:42:42Z"}
		// {"addr":"10.244.11.194:52486","c_ip":"10.244.11.194","hdr":"2600:1702:1640:960:8ab:c2c2:3617:b1e,162.158.74.146","host":"shitcamp-api-2xvip.ondigitalocean.app","level":"info","msg":"req_ip","time":"2021-09-29T13:42:44Z"}
		
		// logger.WithFields(logger.Fields{
		// 	"c_ip": key,
		// 	"host": c.Request.Host,
		// 	"addr": c.Request.RemoteAddr,
		// 	"hdr":  c.Request.Header.Get("X-Forwarded-For"),
		// }).Info("req_ip")

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
