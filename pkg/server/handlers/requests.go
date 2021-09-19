package handlers

import "github.com/shitcamp-unofficial/shitcamp/pkg/models/shitcamp"

type SetScheduleReq struct {
	Dates []*shitcamp.DateSchedule `json:"dates"`
}

type SetFeaturedStreamersForVodReq struct {
	VideoID           string   `json:"id"`
	FeaturedStreamers []string `json:"featured_streamers"`
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
