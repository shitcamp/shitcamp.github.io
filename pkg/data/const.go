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

const InitialScheduleStr = `
{
    "is_latest_schedule": true,
	"last_update_time": "2021-09-26T13:45:00.00-07:00",
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
                    "video_id": "1160400711",
					"thumbnail_url": "./opening-ceremony.png"
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
                        "Nmplol",
                        "Malena"
                    ],
                    "video_id": "",
					"thumbnail_url": "./pancake-breakfast.png"
                },
                {
                    "title": "Chopped competition",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella",
                        "JustaMinx"
                    ],
                    "video_id": "",
					"thumbnail_url": "./chopped-competition.png"
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
                    "video_id": "",
					"thumbnail_url": "./mogul-money.jpg"
                },
                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T17:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi",
                        "QTCinderella"
                    ],
                    "video_id": "",
					"thumbnail_url": "./scavenger-hunt.jpg"
                },
                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T17:00:00.00-07:00",
                    "user_name": "xQcOW",
                    "featured_users": [
                        "xQcOW"
                    ],
                    "video_id": "",
					"thumbnail_url": "./scavenger-hunt.jpg"
                },
                {
                    "title": "Ghost stories and Smores",
                    "start_time": "2021-09-27T20:30:00.00-07:00",
                    "user_name": "WillNeff",
                    "featured_users": [
                        "WillNeff"
                    ],
                    "video_id": "",
					"thumbnail_url": "./ghost-stories.png"
                }
            ]
        },
        {
            "date": "2021-09-28",
            "events": [
                {
                    "title": "French toast breakfast",
                    "start_time": "2021-09-28T08:00:00.00-07:00",
                    "user_name": "EsfandTV",
                    "featured_users": [
                        "EsfandTV"
                    ],
                    "video_id": "",
					"thumbnail_url": "./french-toast-breakfast.png"
                },
                {
                    "title": "IRL stream- Morning Hike",
                    "start_time": "2021-09-28T09:00:00.00-07:00",
                    "user_name": "Myth",
                    "featured_users": [
                        "Myth"
                    ],
                    "video_id": "",
					"thumbnail_url": "./hike.png"
                },
                {
                    "title": "IRL Stream- Gun range",
                    "start_time": "2021-09-28T13:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi", "xQcOW"
                    ],
                    "video_id": "",
					"thumbnail_url": "./gun-range.jpg"
                },
                {
                    "title": "PJ Party",
                    "start_time": "2021-09-28T19:00:00.00-07:00",
                    "user_name": "Sodapoppin",
                    "description": "Face masks, Nail painting, Never have I ever, Who's most likely to.",
                    "featured_users": [
                        "Sodapoppin",
                        "QTCinderella"
                    ],
                    "video_id": "",
					"thumbnail_url": "./pj-party.jpg"
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
                        "Nmplol",
                        "Malena"
                    ],
                    "video_id": "",
					"thumbnail_url": "./grand-slam-breakfast.png"
                },
                {
                    "title": "Kickball tournament ft. 100 Thieves, OfflineTV",
                    "start_time": "2021-09-29T11:00:00.00-07:00",
                    "user_name": "Ludwig",
                    "description": "32 streamers. Kickball. The Most Ambitious Crossover Event in History.",
                    "featured_users": [
                        "Adeptthebest",
                        "AustinShow",
                        "ConnorEatsPants",
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
					"thumbnail_url": "./kickball-tournament.jfif"
                },
                {
                    "title": "Closing Ceremony",
                    "start_time": "2021-09-29T19:00:00.00-07:00",
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
					"thumbnail_url": "./closing-ceremony.jpg"
                }
            ]
        }
    ]
}
`

const InitialTeamsInfo = `
{
    "teams": [
    ]
}
`

// Featured users for live streams
const InitialStreamIDFeaturedMapStr = `
{
    "40020209259": [
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
    "40018159803": ["Nmplol", "Sodapoppin", "JustaMinx", "Kaceytron", "Malena"],
    "40013385979": ["Nmplol", "Malena", "Sodapoppin", "QTCinderella", "Ludwig", "Cyr"]
}
`

// Featured users for vods
const InitialVodIDFeaturedMapStr = `
{
    "1160400711": [
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
    "1159913213": ["Nmplol", "Sodapoppin", "JustaMinx", "Kaceytron", "Malena"],
    "1159263358": ["Sodapoppin", "Ludwig", "Nmplol"],
    "1158909839": ["Nmplol", "Malena", "Sodapoppin", "QTCinderella", "Ludwig", "Cyr"],
    "1158523015": ["QTCinderella", "JustaMinx"],
    "1158517155": ["QTCinderella", "JustaMinx"],
    "1158517135": ["JustaMinx", "QTCinderella"]
}
`
