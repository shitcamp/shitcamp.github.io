package handlers

import (
	"net/http"

	logger "github.com/sirupsen/logrus"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"

	"github.com/gin-gonic/gin"
)

func GetLiveStreams(c *gin.Context) {
	req := new(GetLiveStreamsReq)
	if err := c.ShouldBindQuery(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	streams, err := twitch.GetStreams()
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("GetLiveStreams_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	resp := &GetLiveStreamsResp{Streams: streams}
	respSuccess(c, resp)
}

func GetVods(c *gin.Context) {
	req := new(GetVodsReq)
	if err := c.ShouldBindQuery(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	vods, err := twitch.GetVods()
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("GetVods_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	resp := &GetVodsResp{Vods: vods}
	respSuccess(c, resp)
}

func GetClips(c *gin.Context) {
	req := new(GetClipsReq)
	if err := c.ShouldBindQuery(req); err != nil {
		respError(c, http.StatusBadRequest, err)
		return
	}

	clips, err := twitch.GetClips()
	if err != nil {
		logger.WithField("req", req).WithError(err).Error("GetClips_error")
		respError(c, http.StatusInternalServerError, err)
		return
	}

	resp := &GetClipsResp{Clips: clips}
	respSuccess(c, resp)
}
