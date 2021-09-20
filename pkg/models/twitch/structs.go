package twitch

import "time"

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

type LiveStream struct {
	Video
	//id: "43740004317",
	//user_name: "tonytigre",
	//title: "drawing doodles for new followers! | !insta !doodle",
	//created_at: "2021-09-18T08:17:04Z",
	//url: `https://www.twitch.tv/tonytigre`,
	//thumbnail_url:
	//"https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
	//view_count: 102,
}

type Vod struct {
	Video
	Duration string `json:"duration"`
	//id: "1151405309",
	//user_name: "ludwig",
	//title:
	//"CAN THREE GAMERS BEAT THE BEST SMASH ULTIMATE PLAYER IN THE WORLD?* (usa*)",
	//created_at: "2021-09-17T17:11:07Z",
	//url: "https://www.twitch.tv/videos/1148914392",
	//thumbnail_url:
	//"https://static-cdn.jtvnw.net/cf_vods/d1m7jfoe9zdc1j/4ec1d0730e7592010c4b_ludwig_43704072301_1631722525//thumb/thumb0-%{width}x%{height}.jpg",
	//view_count: 161803,
	//duration: "37m38s",
	//featured_users: ["Ludwig", "Stanz"],
}

type Clip struct {
	ID              string    `json:"id"`
	BroadcasterName string    `json:"broadcaster_name"`
	Title           string    `json:"title"`
	CreatedAt       time.Time `json:"created_at"`
	URL             string    `json:"url"`
	ThumbnailURL    string    `json:"thumbnail_url"`
	ViewCount       int       `json:"view_count"`
	Duration        string    `json:"duration"`

	//id: "UgliestFrailGarageNinjaGrumpy-2Vbp2Vo9tOhlPCUT",
	//broadcaster_name: "ludwig",
	//title: "Ludwig on Mizkif and Maya",
	//created_at: "2021-09-17T00:22:16Z",
	//url: "https://clips.twitch.ztv/UgliestFrailGarageNinjaGrumpy-2Vbp2Vo9tOhlPCUT",
	//thumbnail_url:
	//"https://clips-media-assets2.twitch.tv/AT-cm%7CJ3q8vdlW6dlaPbToZTwlEw-preview-480x272.jpg",
	//view_count: 75088,
	//duration: "30s",
}
