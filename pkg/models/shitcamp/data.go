package shitcamp

import (
	"encoding/json"

	logger "github.com/sirupsen/logrus"
)

var allStreamers = map[string]*User{
	"AdeptTheBest": {
		Name: "AdeptTheBest", Description: "", ProfileImageURL: "",
	},
	"AustinShow": {
		Name: "AustinShow", Description: "", ProfileImageURL: "",
	},
	"Cyr": {
		Name: "Cyr", Description: "", ProfileImageURL: "",
	},
	"EsfandTV": {
		Name: "EsfandTV", Description: "", ProfileImageURL: "",
	},
	"HasanAbi": {
		Name: "HasanAbi", Description: "", ProfileImageURL: "",
	},
	"JustaMinx": {
		Name: "JustaMinx", Description: "", ProfileImageURL: "",
	},
	"Jschlatt": {
		Name: "Jschlatt", Description: "", ProfileImageURL: "",
	},
	"Kaceytron": {
		Name: "Kaceytron", Description: "", ProfileImageURL: "",
	},
	"Ludwig": {
		Name: "Ludwig", Description: "", ProfileImageURL: "",
	},
	"Malena": {
		Name: "Malena", Description: "", ProfileImageURL: "",
	},
	"Myth": {
		Name: "Myth", Description: "", ProfileImageURL: "",
	},
	"Nmplol": {
		Name: "Nmplol", Description: "", ProfileImageURL: "",
	},
	"QTCinderella": {
		Name: "QTCinderella", Description: "", ProfileImageURL: "",
	},
	"Sodapoppin": {
		Name: "Sodapoppin", Description: "", ProfileImageURL: "",
	},
	"WillNeff": {
		Name: "WillNeff", Description: "", ProfileImageURL: "",
	},
	"xQcOW": {
		Name: "xQcOW", Description: "", ProfileImageURL: "",
	},
}

var schedule = make([]*DateSchedule, 0)

//"id": "43740004317",
//"user_name": "QTCinderella",
//"title": "drawing doodles for new followers! | !insta !doodle",
//"created_at": "2021-09-18T08:17:04Z",
//"url": "https://www.twitch.tv/QTCinderella",
//"thumbnail_url":"https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
//"view_count": 102

var vodIDFeaturedMap = map[string][]string{}

func init() {
	// TODO:
	vodIDFeaturedMapStr := `
{
    "1153659413": ["Ludwig", "QTCinderella"]
}
`

	scheduleStr := `
[
        {
            "date": "2021-09-26",
            "events": [
                {
                    "title": "Opening Ceremony",
                    "start_time": "2021-09-26T19:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "Adeptthebest",
                        "AustinShow",
                        "Cyr",
                        "EsfandTV",
                        "HasanAbi",
                        "Jschlatt ",
                        "JustaMinx",
                        "Kaceytron",
                        "Ludwig",
                        "Malena",
                        "Nmplol",
                        "Myth",
                        "QTCinderella",
                        "Sodapoppin",
                        "WillNeff",
                        "xQcOW"
                    ],
                    "vod_link": "",
                    "video": {}
                }
            ]
        },
        {
            "date": "2021-09-27",
            "events": [
                {
                    "title": "Pancake breakfast",
                    "start_time": "2021-09-26T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol"
                    ],
                    "vod_link": "",
                    "video": {}
                },
                {
                    "title": "Chopped",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ],
                    "video_id": "43740004317",
                    "vod_link": "",
                    "video": {}
                },
                {
                    "title": "Mogul Money",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "featured_users": [
                        "Ludwig"
                    ],
                    "vod_link": "",
                    "video": {}
                },
                {
                    "title": "IRL stream",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "featured_users": [
                        "Ludwig"
                    ],
                    "vod_link": "",
                    "video": {}
                },
                {
                    "title": "Ghost stories and Smores",
                    "start_time": "2021-09-26T19:00:00.00-07:00",
                    "user_name": "WillNeff",
                    "featured_users": [
                        "WillNeff"
                    ]
                }
            ]
        },
        {
            "date": "2021-09-28",
            "events": [
                {
                    "title": "French toast breakfast",
                    "start_time": "2021-09-26T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol"
                    ]
                },
                {
                    "title": "Without a Recipe contest",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ]
                },
                {
                    "title": "IRL Stream",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi"
                    ]
                },
                {
                    "title": "Tie Dye Shirts",
                    "start_time": "2021-09-26T15:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "Taco Dinner",
                    "start_time": "2021-09-26T7:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "PJ Party",
                    "start_time": "2021-09-26T20:00:00.00-07:00",
                    "user_name": "Sodapoppin",
                    "featured_users": [
                        "Sodapoppin"
                    ]
                }
            ]
        },
        {
            "date": "2021-09-29",
            "events": [
                {
                    "title": "Grand Slam breakfast",
                    "start_time": "2021-09-26T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol"
                    ]
                },
                {
                    "title": "Secret event",
                    "start_time": "2021-09-26T11:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "Closing Ceremony",
                    "start_time": "2021-09-26T18:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ]
                }
            ]
        }
    ]
`

	err := json.Unmarshal([]byte(vodIDFeaturedMapStr), &vodIDFeaturedMap)
	if err != nil {
		logger.WithError(err).Error("init_featured_map_error")
	}

	var newSchedule = make([]*DateSchedule, 0)
	err = json.Unmarshal([]byte(scheduleStr), &newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_unmarshal_schedule_error")
		return
	}

	err = SetSchedule(newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_SetSchedule_error")
	}
}
