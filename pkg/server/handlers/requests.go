package handlers

import "github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"

type SetScheduleReq struct {
	Dates []*shitcamp.DateSchedule `json:"dates"`
}

type SetFeaturedUsersForVodReq struct {
	VideoID       string   `json:"id"`
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
