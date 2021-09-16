package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shitcamp-unofficial/shitcamp/pkg/common"
)

func respSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, common.Response{Data: data})
}

func respError(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, common.Response{Error: err.Error()})
}

type GetReferralCodeResp struct {
	ReferralCode string `json:"referralCode"`
}

type CreateCheckoutSessionResp struct {
	SessionID string `json:"sessionId"`
}

type CreateCustomerPortalResp struct {
	URL string `json:"url"`
}
