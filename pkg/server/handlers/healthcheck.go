package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	req := new(HealthCheckReq)
	if err := c.ShouldBindQuery(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	respSuccess(c, nil)
}
