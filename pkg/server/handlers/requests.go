package handlers

import (
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/schedule"
)

type SetScheduleReq struct {
	Dates []*schedule.DateSchedule `json:"dates"`
}

type SetFeaturedUsersForVodReq struct {
	VideoID       string   `json:"id"`
	FeaturedUsers []string `json:"featured_users"`
}

type SetFeaturedUsersForStreamReq struct {
	StreamID      string   `json:"id"`
	FeaturedUsers []string `json:"featured_users"`
}

type GetLiveStreamsReq struct {
	Users []string `form:"user"`
}

type GetVodsReq struct {
	Users []string `form:"user"`
}

type GetClipsReq struct {
	Users []string `form:"user"`
}
