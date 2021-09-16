package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"
)

func GetPaymentSetup(c *gin.Context) {
	resp := twitch.GetSetup()

	respSuccess(c, resp)
}

func GetCheckoutSession(c *gin.Context) {
	req := new(GetCheckoutSessionReq)
	if err := c.ShouldBindQuery(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	respSuccess(c, s)
}
