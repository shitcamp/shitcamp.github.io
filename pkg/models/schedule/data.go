package schedule

import (
	"encoding/json"

	logger "github.com/sirupsen/logrus"
)

var schedule = make([]*DateSchedule, 0)

//"id": "43740004317",
//"user_name": "QTCinderella",
//"title": "drawing doodles for new followers! | !insta !doodle",
//"created_at": "2021-09-18T08:17:04Z",
//"url": "https://www.twitch.tv/QTCinderella",
//"thumbnail_url":"https://static-cdn.jtvnw.net/previews-ttv/live_user_tonytigre-{width}x{height}.jpg",
//"view_count": 102

func init() {
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
                    "video_id": "",
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
                    "video_id": "",
                },
                {
                    "title": "Chopped",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ],
                    "video_id": "43740004317",
                },
                {
                    "title": "Mogul Money",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "featured_users": [
                        "Ludwig"
                    ],
                    "video_id": "",
                },
                {
                    "title": "IRL stream",
                    "start_time": "2021-09-26T12:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "featured_users": [
                        "Ludwig"
                    ],
                    "video_id": "",
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

	var newSchedule = make([]*DateSchedule, 0)
	err := json.Unmarshal([]byte(scheduleStr), &newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_unmarshal_schedule_error")
		return
	}

	err = SetSchedule(newSchedule)
	if err != nil {
		logger.WithError(err).Error("init_SetSchedule_error")
	}
}
