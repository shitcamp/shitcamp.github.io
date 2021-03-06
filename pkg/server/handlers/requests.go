package handlers

import (
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/schedule"
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"
)

type HealthCheckReq struct {
	Users []string `form:"user"`
}

type SetScheduleReq struct {
	*schedule.Schedule
}

type SetTeamsInfoReq struct {
	*shitcamp.TeamsInfo
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
