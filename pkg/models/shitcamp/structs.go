package shitcamp

import (
	"time"
)

type User struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
}

type Video struct {
	ID                string    `json:"id"`
	UserName          string    `json:"user_name"`
	Title             string    `json:"title"`
	CreatedAt         time.Time `json:"created_at"`
	URL               string    `json:"url"`
	ThumbnailURL      string    `json:"thumbnail_url"`
	ViewCount         int       `json:"view_count"`
	FeaturedStreamers []string  `json:"featured_users"`
}

type Event struct {
	Title         string    `json:"title"`
	StartTime     time.Time `json:"start_time"`
	UserName      string    `json:"user_name"`
	FeaturedUsers []string  `json:"featured_users"`
	VideoID       string    `json:"video_id"`
	Video         Video     `json:"video"`
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
//			featured_users: ["Nmplol"],
//			video_id: "",
//          video: {}
//		}
//	],
//},
