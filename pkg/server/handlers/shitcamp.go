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

	err := shitcamp.SetSchedule(req.Dates)
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetSchedule_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}

func SetFeaturedUsersForVod(c *gin.Context) {
	req := new(SetFeaturedUsersForVodReq)
	if err := c.ShouldBindJSON(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	err := shitcamp.SetFeaturedUsersForVod(req.VideoID, req.FeaturedUsers)
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetFeaturedUsersForVod_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}
