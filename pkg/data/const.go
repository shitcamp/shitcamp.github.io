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
    "dates":[
       {
          "date":"2021-09-26",
          "events":[
             {
                "title":"Opening Ceremony",
                "start_time":"2021-09-26T19:00:00-07:00",
                "description":"Shitcamp is officially here! The captains pick their teams and then compete in games of trivia, roulette, hide 'n seek and mafia.",
                "thumbnail_url":"./opening-ceremony.png",
                "user_name":"xQcOW",
                "featured_users":[
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
                "video_id":"",
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
				}
             }
          ]
       },
       {
          "date":"2021-09-27",
          "events":[
             {
                "title":"Pancake breakfast",
                "start_time":"2021-09-27T09:00:00-07:00",
                "description":"The campers prepare a scuffed breakfast and recount Malena's drunken behavior from the previous night.",
                "thumbnail_url":"./pancake-breakfast.png",
                "user_name":"Nmplol",
                "featured_users":[
                   "Nmplol",
                   "Malena"
                ],
                "video_id":"1160772412"
             },
             {
                "title":"Chopped competition",
                "start_time":"2021-09-27T12:45:00-07:00",
                "description":"Extremely talented hardcore gamers whip up Michelin star dishes from random ingredients.",
                "thumbnail_url":"./chopped-competition.png",
                "user_name":"QTCinderella",
                "featured_users":[
                   "QTCinderella",
                   "Ludwig",
                   "xQcOW",
                   "JustaMinx",
                   "Cyr",
                   "Sodapoppin",
                   "Adeptthebest",
                   "EsfandTV"
                ],
                "video_id":"1160962922"
             },
             {
                "title":"IRL Stream- Scavenger Hunt",
                "start_time":"2021-09-27T17:00:00-07:00",
                "description":"Hasan is forced to go outside and touch grass while leading his team in a scavenger hunt.",
                "thumbnail_url":"./scavenger-hunt.jpg",
                "user_name":"HasanAbi",
                "featured_users":[
                   "HasanAbi"
                ],
                "video_id":"1161184256"
             },
             {
                "title":"IRL Stream- Scavenger Hunt",
                "start_time":"2021-09-27T17:00:00-07:00",
                "description":"The juicer and his team try to speedrun shooting polaroids of items for the scavenger hunt.",
                "thumbnail_url":"./scavenger-hunt.jpg",
                "user_name":"xQcOW",
                "featured_users":[
                   "xQcOW",
                   "QTCinderella"
                ],
                "video_id":"1161192739"
             },
             {
                "title":"Boat racing",
                "start_time":"2021-09-27T22:30:00-07:00",
                "description":"Ready, Set, Sail! The two teams compete in a cardboard boat making competition while doing drunk challenges to gain an edge.",
                "thumbnail_url":"./boat-racing.jpg",
                "user_name":"WillNeff",
                "featured_users":[
                   "WillNeff"
                ],
                "video_id":"1161394966"
             }
          ]
       },
       {
          "date":"2021-09-28",
          "events":[
             {
                "title":"French toast breakfast",
                "start_time":"2021-09-28T08:00:00-07:00",
                "description":"",
                "thumbnail_url":"./french-toast-breakfast.png",
                "user_name":"EsfandTV",
                "featured_users":[
                   "EsfandTV"
                ],
                "video_id":""
             },
             {
                "title":"IRL stream- Morning Hike",
                "start_time":"2021-09-28T09:00:00-07:00",
                "description":"",
                "thumbnail_url":"./hike.png",
                "user_name":"Myth",
                "featured_users":[
                   "Myth"
                ],
                "video_id":""
             },
             {
                "title":"IRL Stream- Gun range",
                "start_time":"2021-09-28T11:30:00-07:00",
                "description":"",
                "thumbnail_url":"./gun-range.jpg",
                "user_name":"HasanAbi",
                "featured_users":[
                   "HasanAbi",
                   "xQcOW",
                   "Ludwig",
                   "Nmplol"
                ],
                "video_id":""
             },
             {
                "title":"Mogul Money (time unconfirmed)",
                "start_time":"2021-09-28T15:00:00-07:00",
                "description":"",
                "thumbnail_url":"./mogul-money.jpg",
                "user_name":"Ludwig",
                "featured_users":[
                   "Ludwig",
                   "Sodapoppin",
                   "xQcOW"
                ],
                "video_id":""
             },
             {
                "title":"PJ Party",
                "start_time":"2021-09-28T19:00:00-07:00",
                "description":"Face masks, Nail painting, Never have I ever, Who's most likely to.",
                "thumbnail_url":"./pj-party.jpg",
                "user_name":"Sodapoppin",
                "featured_users":[
                   "Sodapoppin",
                   "QTCinderella"
                ],
                "video_id":""
             }
          ]
       },
       {
          "date":"2021-09-29",
          "events":[
             {
                "title":"Grand Slam breakfast",
                "start_time":"2021-09-29T08:00:00-07:00",
                "description":"",
                "thumbnail_url":"./grand-slam-breakfast.png",
                "user_name":"Nmplol",
                "featured_users":[
                   "Nmplol",
                   "Malena"
                ],
                "video_id":""
             },
             {
                "title":"Kickball tournament ft. 100 Thieves, OfflineTV",
                "start_time":"2021-09-29T11:00:00-07:00",
                "description":"32 streamers. Kickball. The Most Ambitious Crossover Event in History.",
                "thumbnail_url":"./kickball-tournament.jfif",
                "user_name":"Ludwig",
                "featured_users":[
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
                "video_id":""
             },
             {
                "title":"Closing Ceremony",
                "start_time":"2021-09-29T19:00:00-07:00",
                "description":"",
                "thumbnail_url":"./closing-ceremony.jpg",
                "user_name":"QTCinderella",
                "featured_users":[
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
                "video_id":""
             }
          ]
       }
    ],
    "is_latest_schedule":true,
    "last_update_time":"2021-09-27T20:00:00.2-07:00"
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
	],
   "team_results": [
      {
         "team_name": "Team America",
         "total_points": 270,
         "results": [
            {
               "contest_name": "Hide 'n seek",
               "points": 100,
               "description": "Malena was the last one to be found.",
               "vod_url": ""
            },
            {
               "contest_name": "Chopped",
               "points": 50,
               "description": "WillNeff came in 2nd place.",
               "vod_url": ""
            },
            {
               "contest_name": "Boat race",
               "points": 120,
               "description": "",
               "vod_url": ""
            }
         ]
      },
      {
         "team_name": "Team Cummunism",
         "total_points": 400,
         "results": [
            {
               "contest_name": "Trivia",
               "points": 100,
               "description": "",
               "vod_url": ""
            },
            {
               "contest_name": "Chopped",
               "points": 100,
               "description": "xQcOW won the cooking contest.",
               "vod_url": ""
            },
            {
               "contest_name": "Scavenger Hunt",
               "points": 200,
               "description": "",
               "vod_url": ""
            }
         ]
      }
   ]
}
`

// TODO: struct with stream_id, video_id, featured_streamers to reduce duplication

// Featured users for live streams
const InitialStreamIDFeaturedMapStr = `
{
   "43868406781":[
      "WillNeff",
      "xQcOW",
      "Sodapoppin",
      "QTCinderella",
      "JustaMinx",
      "Kaceytron",
      "HasanAbi",
      "AustinShow",
      "Adeptthebest",
      "Myth",
      "Cyr",
      "EsfandTV"
   ],
   "43384674940":[
      "xQcOW",
      "Ludwig",
      "Jschlatt",
      "Nmplol",
      "QTCinderella",
      "Sodapoppin",
      "JustaMinx",
      "Kaceytron"
   ],
   "40024100859":[
      "HasanAbi",
      "WillNeff",
      "Myth",
      "Malena",
      "EsfandTV",
      "AustinShow",
      "AdeptTheBest",
      "Cyr"
   ],
   "43862301165":[
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
   ],
   "40022962763":[
      "Kaceytron",
      "JustaMinx",
      "AustinShow",
      "Jschlatt",
      "Cyr",
      "Myth",
      "Malena",
      "QTCinderella"
   ],
   "40022423755":[
      "Nmplol",
      "Malena",
      "AustinShow",
      "EsfandTV",
      "Myth"
   ],
   "40020209259":[
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
   "40018159803":[
      "Nmplol",
      "Sodapoppin",
      "JustaMinx",
      "Kaceytron",
      "Malena"
   ],
   "40013385979":[
      "Nmplol",
      "Malena",
      "Sodapoppin",
      "QTCinderella",
      "Ludwig",
      "Cyr"
   ]
}
`

// All featured
// 1160772412: QTCinderella, JustaMinx, Ludwig, Kaceytron, PhinTTV
// "1160901842": ["Kaceytron", "JustaMinx", "AustinShow", "Jschlatt", "Cyr", "Myth", "Malena", "QTCinderella"]
// "1160962922": Nmplol, Jschlatt

// Featured users for vods
const InitialVodIDFeaturedMapStr = `
{
   "1161394966":[
      "WillNeff",
      "xQcOW",
      "Sodapoppin",
      "QTCinderella",
      "JustaMinx",
      "Kaceytron",
      "HasanAbi",
      "AustinShow",
      "Adeptthebest",
      "Myth",
      "Cyr",
      "EsfandTV",
      "PhinTTV"
   ],
   "1161192739":[
      "xQcOW",
      "Ludwig",
      "EsfandTV",
      "Sodapoppin",
      "JustaMinx",
      "Kaceytron",
      "PhinTTV"
   ],
   "1161184256":[
      "HasanAbi",
      "QTCinderella",
      "AustinShow",
      "AdeptTheBest",
      "Cyr",
      "Myth",
      "WillNeff"
   ],
   "1160962922":[
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
   ],
   "1160901842":[
      "Kaceytron",
      "JustaMinx",
      "AustinShow",
      "Jschlatt",
      "Myth",
      "Malena"
   ],
   "1160772412":[
      "Nmplol",
      "EsfandTV",
      "Myth",
      "AustinShow",
      "Malena"
   ],
   "1160549541":[
      "EsfandTV",
      "AustinShow"
   ],
   "1160400711":[
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
   "1159913213":[
      "Nmplol",
      "Sodapoppin",
      "JustaMinx",
      "Kaceytron",
      "Malena"
   ],
   "1159263358":[
      "Sodapoppin",
      "Ludwig",
      "Nmplol"
   ],
   "1158909839":[
      "Nmplol",
      "Malena",
      "Sodapoppin",
      "QTCinderella",
      "Ludwig",
      "Cyr"
   ],
   "1158523015":[
      "QTCinderella",
      "JustaMinx"
   ],
   "1158517155":[
      "QTCinderella",
      "JustaMinx"
   ],
   "1158517135":[
      "JustaMinx",
      "QTCinderella"
   ]
}
`
