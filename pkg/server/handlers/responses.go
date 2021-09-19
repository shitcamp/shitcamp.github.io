package handlers

import (
	"net/http"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"

	"github.com/gin-gonic/gin"
	"github.com/shitcamp-unofficial/shitcamp/pkg/common"
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"
)

func respSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, common.Response{Data: data})
}

func respError(c *gin.Context, httpCode int, err error) {
	c.JSON(httpCode, common.Response{Error: err.Error()})
}

type GetStreamerNamesResp struct {
	Users []string `json:"users"`
}

type GetStreamersResp struct {
	Users []*shitcamp.User `json:"users"`
}

type GetScheduleResp struct {
	Dates []*shitcamp.DateSchedule `json:"dates"`
}

type GetLiveStreamsResp struct {
	Streams []*twitch.LiveStream `json:"streams"`
}

type GetVodsResp struct {
	Vods []*twitch.Vod `json:"vods"`
}

type GetClipsResp struct {
	Clips []*twitch.Clip `json:"clips"`
}
