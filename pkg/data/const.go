package data

type User struct {
	Name            string `json:"name"`
	Description     string `json:"description"`
	ProfileImageURL string `json:"profile_image_url"`
}

var AllStreamers = map[string]*User{
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

// TODO: update is_latest_schedules
const InitialScheduleStr = `
{
    "is_latest_schedule": false,
    "dates": [
        {
            "date": "2021-09-26",
            "events": [
                {
                    "title": "Opening Ceremony",
                    "start_time": "2021-09-26T19:00:00.00-07:00",
                    "user_name": "xQcOW",
					"description": "Opening night with trivia and mafia games",
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
                    "video_id": ""
                }
            ]
        },
        {
            "date": "2021-09-27",
            "events": [
                {
                    "title": "Pancake breakfast",
                    "start_time": "2021-09-27T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol", "Malena"
                    ],
                    "video_id": ""
                },
                {
                    "title": "Chopped competition",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella", "JustaMinx"
                    ],
                    "video_id": ""
                },
                {
                    "title": "Mogul Money",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "featured_users": [
                        "Ludwig",
                        "Sodapoppin",
                        "xQcOW"
                    ],
                    "video_id": ""
                },
                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi"
                    ],
                    "video_id": ""
                },
                                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "xQcOW",
                    "featured_users": [
                        "xQcOW"
                    ],
                    "video_id": ""
                },
                {
                    "title": "Ghost stories and Smores",
                    "start_time": "2021-09-27T19:00:00.00-07:00",
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
                    "start_time": "2021-09-28T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol", "Malena"
                    ]
                },
                {
                    "title": "IRL stream- Hike",
                    "start_time": "2021-09-28T08:00:00.00-07:00",
                    "user_name": "Myth",
                    "featured_users": [
                        "Myth"
                    ]
                },
                {
                    "title": "Without a Recipe contest",
                    "start_time": "2021-09-28T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ]
                },
                {
                    "title": "IRL Stream- Gun range",
                    "start_time": "2021-09-28T12:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi"
                    ]
                },
                {
                    "title": "Tie Dye Shirts",
                    "start_time": "2021-09-28T15:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "Taco Dinner",
                    "start_time": "2021-09-28T17:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "PJ Party",
                    "start_time": "2021-09-28T20:00:00.00-07:00",
                    "user_name": "Sodapoppin",
                    "description": "Face masks, Nail painting, Never have I ever, Who's most likely to",
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
                    "start_time": "2021-09-29T08:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol", "Malena"
                    ]
                },
                {
                    "title": "Secret event",
                    "start_time": "2021-09-29T11:30:00.00-07:00",
                    "user_name": "",
                    "featured_users": []
                },
                {
                    "title": "Closing Ceremony",
                    "start_time": "2021-09-29T18:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella"
                    ]
                }
            ]
        }
    ]
}
`

// TODO:
const InitialVodIDFeaturedMapStr = `
{
}
`
