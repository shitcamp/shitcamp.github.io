package handlers

import (
	"net/http"

	logger "github.com/sirupsen/logrus"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"

	"github.com/gin-gonic/gin"
)

func GetStreamerNames(c *gin.Context) {
	names := shitcamp.GetStreamerNames()

	resp := &GetStreamerNamesResp{Users: names}
	respSuccess(c, resp)
}

func GetStreamers(c *gin.Context) {
	users := shitcamp.GetStreamers()

	resp := &GetStreamersResp{Users: users}
	respSuccess(c, resp)
}

func GetSchedule(c *gin.Context) {
	dates := shitcamp.GetSchedule()

	resp := &GetScheduleResp{Dates: dates}
	respSuccess(c, resp)
}

func SetSchedule(c *gin.Context) {
	req := new(SetScheduleReq)
	if err := c.ShouldBindJSON(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	err := shitcamp.SetSchedule()
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetSchedule_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}

func SetFeaturedStreamersForVod(c *gin.Context) {
	req := new(SetFeaturedStreamersForVodReq)
	if err := c.ShouldBindJSON(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	err := shitcamp.SetFeaturedStreamersForVod()
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetFeaturedStreamersForVod_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}
