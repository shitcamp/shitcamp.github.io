package handlers

import (
	"net/http"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/schedule"

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
	s, err := schedule.GetSchedule()
	if err != nil {
		logger.WithError(err).Error("GetSchedule_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	resp := &GetScheduleResp{Schedule: *s}
	respSuccess(c, resp)
}

func SetSchedule(c *gin.Context) {
	req := new(SetScheduleReq)
	if err := c.ShouldBindJSON(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	err := schedule.SetSchedule(req.Schedule)
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetSchedule_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}

func GetFeaturedUsersForStreams(c *gin.Context) {
	featuredStreamers := shitcamp.GetFeaturedStreamersForStreams()

	resp := &GetFeaturedUsersForStreamsResp{FeaturedStreamers: featuredStreamers}
	respSuccess(c, resp)
}

func SetFeaturedUsersForStream(c *gin.Context) {
	req := new(SetFeaturedUsersForStreamReq)
	if err := c.ShouldBindJSON(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	err := shitcamp.SetFeaturedUsersForVod(req.StreamID, req.FeaturedUsers)
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("SetFeaturedUsersForStream_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	respSuccess(c, nil)
}

func GetFeaturedUsersForVods(c *gin.Context) {
	featuredStreamers := shitcamp.GetFeaturedStreamersForVods()

	resp := &GetFeaturedUsersForVodsResp{FeaturedStreamers: featuredStreamers}
	respSuccess(c, resp)
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
