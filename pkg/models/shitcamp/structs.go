package shitcamp

import (
	"github.com/shitcamp-unofficial/shitcamp/pkg/models/twitch"
)

type User struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Event struct {
	Title             string       `json:"title"`
	StartTime         string       `json:"start_time"` // TODO: try time.Time?
	UserName          string       `json:"user_name"`
	FeaturedStreamers []string     `json:"featured_streamers"`
	VideoID           string       `json:"video_id"`
	Video             twitch.Video `json:"video"`
}

type DateSchedule struct {
	Date   string   `json:"date"`
	Events []*Event `json:"events"`
}

//{
//	date: "2021-09-27",
//	events: [
//		{
//			title: "Pancake breakfast",
//			start_time: "08:00:00.00-07:00",
//			user_name: "Nmplol",
//			featured_streamers: ["Nmplol"],
//			video_id: "",
//		}
//	],
//},
