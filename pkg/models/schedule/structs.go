package schedule

import (
	"time"

	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"
)

type Event struct {
	Title         string      `json:"title"`
	StartTime     time.Time   `json:"start_time"`
	UserName      string      `json:"user_name"`
	FeaturedUsers []string    `json:"featured_users"`
	VideoID       string      `json:"video_id"`
	Vod           *twitch.Vod `json:"vod"`
}

type DateSchedule struct {
	Date   string   `json:"date"`
	Events []*Event `json:"events"`
}
