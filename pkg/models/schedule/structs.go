package schedule

import (
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"
)

type Event struct {
	Title         string      `json:"title"`
	StartTime     time.Time   `json:"start_time"`
	Description   string      `json:"description"`
	ThumbnailURL  string      `json:"thumbnail_url"`
	UserName      string      `json:"user_name"`      // TODO: remove?
	FeaturedUsers []string    `json:"featured_users"` // TODO: remove?
	VideoID       string      `json:"video_id"`       // TODO: remove?
	Vod           *twitch.Vod `json:"vod"`
}

type DateSchedule struct {
	Date   string   `json:"date"`
	Events []*Event `json:"events"`
}

type Schedule struct {
	Dates            []*DateSchedule `json:"dates"`
	IsLatestSchedule bool            `json:"is_latest_schedule"`
}
