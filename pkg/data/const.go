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
    "dates": [
        {
            "date": "2021-09-26",
            "events": [
                {
                    "title": "Opening Ceremony",
                    "start_time": "2021-09-26T19:00:00.00-07:00",
                    "user_name": "xQcOW",
                    "description": "Opening night with team selection, and games like trivia, roulette, hide 'n seek and mafia",
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
					"vod": {
						"id": "1160400711",
						"user_name": "xQcOW",
						"title": "Sh*t Camp 2021 Opening Ceremony",
						"created_at": "2021-09-27T01:56:02Z",
						"url": "https://www.youtube.com/watch?v=qA6SVvB3Ff8&ab_channel=QTCinderellaVODs",
						"thumbnail_url": "https://i.ytimg.com/vi/qA6SVvB3Ff8/hqdefault.jpg?sqp=-oaymwEcCNACELwBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLCKTISmtPNiAP4JfctRiZudcQalwQ",
						"view_count": 1518951,
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
							"xQcOW",
							"PhinTTV"
						],
						"duration": "3h49m56s"
					},
					"thumbnail_url": "./opening-ceremony.png"
                }
            ]
        },
        {
            "date": "2021-09-27",
            "events": [
                {
                    "title": "Pancake breakfast",
                    "start_time": "2021-09-27T09:00:00.00-07:00",
                    "user_name": "Nmplol",
                    "featured_users": [
                        "Nmplol",
                        "Malena"
                    ],
                    "video_id": "1160772412",
					"thumbnail_url": "./pancake-breakfast.png"
                },
                {
                    "title": "Chopped competition",
                    "start_time": "2021-09-27T12:00:00.00-07:00",
                    "user_name": "QTCinderella",
                    "featured_users": [
                        "QTCinderella",
                        "Ludwig",
                        "xQcOW",
                        "JustaMinx",
                        "Cyr",
                        "Sodapoppin",
                        "Adeptthebest",
                        "EsfandTV"
                    ],
                    "video_id": "1160962922",
					"thumbnail_url": "./chopped-competition.png"
                },
                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T17:00:00.00-07:00",
                    "user_name": "HasanAbi",
                    "featured_users": [
                        "HasanAbi"
                    ],
                    "video_id": "",
					"thumbnail_url": "./scavenger-hunt.jpg"
                },
                {
                    "title": "IRL Stream- Scavenger Hunt",
                    "start_time": "2021-09-27T17:00:00.00-07:00",
                    "user_name": "xQcOW",
                    "featured_users": [
                        "xQcOW",
                        "QTCinderella"
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
                    "title": "Mogul Money (time unconfirmed)",
                    "start_time": "2021-09-28T12:00:00.00-07:00",
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
    ],
    "is_latest_schedule": true,
	"last_update_time": "2021-09-27T12:45:00.00-07:00"
}
`

const InitialTeamsInfo = `
{
	"teams": [
		{
			"name": "Team America",
			"captain": "HasanAbi",
			"users": [
				"WillNeff",
				"Myth",
				"Malena",
				"EsfandTV",
				"AustinShow",
				"AdeptTheBest",
				"Cyr"
			],
			"thumbnail_url": "./team-america.png"
		},
		{
			"name": "Team Cummunism",
			"captain": "Ludwig",
			"users": [
				"Jschlatt",
				"Nmplol",
				"QTCinderella",
				"Sodapoppin",
				"JustaMinx",
				"Kaceytron",
				"xQcOW"
			],
			"thumbnail_url": "./team-cummunism.png"
		}
	]
}
`

// TODO: struct with stream_id, video_id, featured_streamers to reduce duplication

// Featured users for live streams
const InitialStreamIDFeaturedMapStr = `
{
    "43862301165": [
        "QTCinderella",
        "Ludwig",
        "xQcOW",
        "WillNeff",
        "Cyr",
        "AustinShow",
        "EsfandTV",
        "Sodapoppin",
        "JustaMinx",
        "Adeptthebest"
    ]
    "40022962763": [
        "Kaceytron",
        "JustaMinx",
        "AustinShow",
        "Jschlatt",
        "Cyr",
        "Myth",
        "Malena",
        "QTCinderella"
    ],
    "40022423755": [
        "Nmplol",
        "Malena",
        "AustinShow",
        "EsfandTV",
        "Myth"
    ],
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
        "xQcOW",
        "PhinTTV"
    ],
    "40018159803": ["Nmplol", "Sodapoppin", "JustaMinx", "Kaceytron", "Malena"],
    "40013385979": ["Nmplol", "Malena", "Sodapoppin", "QTCinderella", "Ludwig", "Cyr"]
}
`

// 1160772412: QTCinderella, JustaMinx, Ludwig, Kaceytron, PhinTTV
// "1160901842": ["Kaceytron", "JustaMinx", "AustinShow", "Jschlatt", "Cyr", "Myth", "Malena", "QTCinderella"]
// "1160962922": Nmplol, Jschlatt

// Featured users for vods
const InitialVodIDFeaturedMapStr = `
{
    "1160962922": [
        "QTCinderella",
        "Ludwig",
        "xQcOW",
        "WillNeff",
        "Cyr",
        "AustinShow",
        "EsfandTV",
        "Sodapoppin",
        "JustaMinx",
        "Adeptthebest",
        "PhinTTV"
    ]
    "1160901842": [
        "Kaceytron",
        "JustaMinx",
        "AustinShow",
        "Jschlatt",
        "Myth",
        "Malena"
    ],
    "1160772412": [
        "Nmplol",
        "EsfandTV",
        "Myth",
        "AustinShow",
        "Malena"
    ],
    "1160549541": [
        "EsfandTV",
        "AustinShow"
    ],
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
        "xQcOW",
        "PhinTTV"
    ],
    "1159913213": ["Nmplol", "Sodapoppin", "JustaMinx", "Kaceytron", "Malena"],
    "1159263358": ["Sodapoppin", "Ludwig", "Nmplol"],
    "1158909839": ["Nmplol", "Malena", "Sodapoppin", "QTCinderella", "Ludwig", "Cyr"],
    "1158523015": ["QTCinderella", "JustaMinx"],
    "1158517155": ["QTCinderella", "JustaMinx"],
    "1158517135": ["JustaMinx", "QTCinderella"]
}
`
