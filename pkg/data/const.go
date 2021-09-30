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
                   "start_time": "2021-09-26T19:00:00-07:00",
                   "description": "Shitcamp is officially here! The captains pick their teams and then compete in games of trivia, roulette, hide 'n seek and mafia.",
                   "thumbnail_url": "./opening-ceremony.png",
                   "user_name": "xQcOW",
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
                   }
               }
           ]
       },
       {
           "date": "2021-09-27",
           "events": [
               {
                   "title": "Pancake breakfast",
                   "start_time": "2021-09-27T09:00:00-07:00",
                   "description": "The campers prepare a scuffed breakfast and recount Malena's drunken behavior from the previous night.",
                   "thumbnail_url": "./pancake-breakfast.png",
                   "user_name": "Nmplol",
                   "featured_users": [
                       "Nmplol",
                       "Malena"
                   ],
                   "video_id": "1160772412"
               },
               {
                   "title": "Chopped competition",
                   "start_time": "2021-09-27T12:45:00-07:00",
                   "description": "Extremely talented hardcore gamers whip up Michelin star dishes from random ingredients.",
                   "thumbnail_url": "./chopped-competition.png",
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
                   "video_id": "",
                   "vod": {
                       "id": "1160962922",
                       "user_name": "QTCinderella",
                       "title": "Running a little late- Chopped Next! | | !Subtember !po !clips !youtube !instagram",
                       "created_at": "2021-09-27T19:44:02Z",
                       "url": "https://www.youtube.com/watch?v=Hd3iUtuFj3Y&ab_channel=QTCinderellaVODs",
                       "thumbnail_url": "https://i.ytimg.com/vi/Hd3iUtuFj3Y/hqdefault.jpg?sqp=-oaymwEcCNACELwBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLBdhspIwatvM8oSO0nOD0mYfpg7oQ",
                       "view_count": 420417,
                       "featured_users": [
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
                       "duration": "2h16m57s"
                   }
               },
               {
                   "title": "IRL Stream- Scavenger Hunt",
                   "start_time": "2021-09-27T17:00:00-07:00",
                   "description": "Hasan is forced to go outside and touch grass while leading his team in a scavenger hunt.",
                   "thumbnail_url": "./scavenger-hunt.jpg",
                   "user_name": "HasanAbi",
                   "featured_users": [
                       "HasanAbi"
                   ],
                   "video_id": "",
                   "vod": {
                       "id": "1161184256",
                       "user_name": "HasanAbi",
                       "title": "REGULAR STREAM THEN IRL SCAVENGER HUNT AT SHITCAMP - SCHEDULE: !shitcamp",
                       "created_at": "2021-09-28T00:15:46Z",
                       "url": "https://www.youtube.com/watch?v=we-3i-DfyWs&ab_channel=QTCinderellaVODs",
                       "thumbnail_url": "https://i.ytimg.com/vi/we-3i-DfyWs/hqdefault.jpg?sqp=-oaymwEcCNACELwBSFXyq4qpAw4IARUAAIhCGAFwAcABBg==&rs=AOn4CLD9hOZUkkshMyL92zS39ZLGC5oGqA",
                       "view_count": 448840,
                       "featured_users": [
                           "HasanAbi",
                           "QTCinderella",
                           "AustinShow",
                           "AdeptTheBest",
                           "Cyr",
                           "Myth",
                           "WillNeff"
                       ],
                       "duration": "2h24m12s"
                   }
               },
               {
                   "title": "IRL Stream- Scavenger Hunt",
                   "start_time": "2021-09-27T17:00:00-07:00",
                   "description": "The juicer and his team try to speedrun shooting polaroids of items for the scavenger hunt.",
                   "thumbnail_url": "./scavenger-hunt.jpg",
                   "user_name": "xQcOW",
                   "featured_users": [
                       "xQcOW",
                       "QTCinderella"
                   ],
                   "video_id": "1161192739"
               },
               {
                   "title": "Boat racing",
                   "start_time": "2021-09-27T22:30:00-07:00",
                   "description": "Ready, Set, Sail! The two teams compete in a cardboard boat making competition while doing drunk challenges to gain an edge.",
                   "thumbnail_url": "./boat-racing.jpg",
                   "user_name": "WillNeff",
                   "featured_users": [
                       "WillNeff"
                   ],
                   "video_id": "1161394966"
               }
           ]
       },
       {
           "date": "2021-09-28",
           "events": [
               {
                   "title": "French toast breakfast",
                   "start_time": "2021-09-28T09:00:00-07:00",
                   "description": "EsfandTV and QTCinderella contribute equally towards feeding the campers on a breakfast stream that definitely started on time",
                   "thumbnail_url": "./french-toast-breakfast.png",
                   "user_name": "EsfandTV",
                   "featured_users": [
                       "EsfandTV"
                   ],
                   "video_id": "1161751195"
               },
               {
                   "title": "IRL Stream- Gun range",
                   "start_time": "2021-09-28T11:30:00-07:00",
                   "description": "Noobs with guns try to not shoot the cameraman. The two teams then compete in a shoot-off.",
                   "thumbnail_url": "./gun-range.jpg",
                   "user_name": "HasanAbi",
                   "featured_users": [
                       "HasanAbi",
                       "xQcOW",
                       "Ludwig",
                       "Nmplol"
                   ],
                   "video_id": "1161828036"
               },
               {
                   "title": "Workout stream",
                   "start_time": "2021-09-28T16:30:00-07:00",
                   "description": "Relatable streamers stall by hanging out, throwing around the old pig skin, having a taco party and then go to pump iron.",
                   "thumbnail_url": "./workout.jpg",
                   "user_name": "EsfandTV",
                   "featured_users": [
                       "EsfandTV",
                       "WillNeff",
                       "AustinShow",
                       "Myth"
                   ],
                   "video_id": "1162074919"
               },
               {
                   "title": "Mogul Money",
                   "start_time": "2021-09-28T17:00:00-07:00",
                   "description": "Sigma campers compete for money, a bidet and batteries on a completely original game show that is definitely not inspired by Jeopardy!",
                   "thumbnail_url": "./mogul-money.jpg",
                   "user_name": "Ludwig",
                   "featured_users": [
                       "Ludwig",
                       "Sodapoppin",
                       "xQcOW"
                   ],
                   "video_id": "1162103704"
               },
               {
                   "title": "PJ Party",
                   "start_time": "2021-09-28T20:30:00-07:00",
                   "description": "The PJ party takes a turn for the worse when the boys decide to take over \"girls night\". Tune in for a night of Do or Drink, face masks, makeovers and Never Have I Ever.",
                   "thumbnail_url": "./pj-party.jpg",
                   "user_name": "Sodapoppin",
                   "featured_users": [
                       "Sodapoppin",
                       "QTCinderella"
                   ],
                   "video_id": "1162277299"
               }
           ]
       },
       {
           "date": "2021-09-29",
           "events": [
               {
                   "title": "Grand Slam breakfast",
                   "start_time": "2021-09-29T08:00:00-07:00",
                   "description": "The campers wake up early for the final breakfast stream and cook pancakes with homemade syrup.",
                   "thumbnail_url": "./grand-slam-breakfast.png",
                   "user_name": "Myth",
                   "featured_users": [
                       "Myth"
                   ],
                   "video_id": "1162576914"
               },
               {
                   "title": "Kickball tournament ft. 100 Thieves, OfflineTV",
                   "start_time": "2021-09-29T12:30:00-07:00",
                   "description": "32 streamers from 4 teams compete in a Kickball tournament to win $10000 for their charities. Will the Shitcamp teams be able to beat the orgs? Come find out.",
                   "thumbnail_url": "./kickball-tournament.jfif",
                   "user_name": "Ludwig",
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
                   "video_id": ""
               },
               {
                   "title": "Closing Ceremony",
                   "start_time": "2021-09-29T19:00:00-07:00",
                   "description": "Team American and Team Cummunism go up against one last time in games of Charades and Family Fued.",
                   "thumbnail_url": "./closing-ceremony.jpg",
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
                   "video_id": "1163160816"
               }
           ]
       }
   ],
   "is_latest_schedule": true,
   "last_update_time": "2021-09-29T19:20:00-07:00"
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
         "total_points": 570,
         "results": [
            {
               "contest_name": "Hide 'n seek",
               "points": 100,
               "description": "Malena was the last one to be found.",
               "vod_url": "https://www.youtube.com/watch?v=qA6SVvB3Ff8&ab_channel=QTCinderellaVODs"
            },
            {
               "contest_name": "Chopped",
               "points": 50,
               "description": "WillNeff came in 2nd place.",
               "vod_url": "https://www.twitch.tv/videos/1160962922"
            },
            {
               "contest_name": "Boat race",
               "points": 120,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1161394966"
            },
            {
               "contest_name": "Gun range shoot-off",
               "points": 200,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1161828036"
            },
            {
               "contest_name": "Charades",
               "points": 100,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1163160816"
            }
         ]
      },
      {
         "team_name": "Team Cummunism",
         "total_points": 669,
         "results": [
            {
               "contest_name": "Trivia",
               "points": 100,
               "description": "",
               "vod_url": "https://www.youtube.com/watch?v=qA6SVvB3Ff8&ab_channel=QTCinderellaVODs"
            },
            {
               "contest_name": "Chopped",
               "points": 100,
               "description": "xQcOW won the cooking contest.",
               "vod_url": "https://www.twitch.tv/videos/1160962922"
            },
            {
               "contest_name": "Scavenger Hunt",
               "points": 200,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1161192739"
            },
            {
               "contest_name": "Kickball",
               "points": 200,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1162803783"
            },
            {
               "contest_name": "Family Fued",
               "points": 69,
               "description": "",
               "vod_url": "https://www.twitch.tv/videos/1163160816"
            }
         ]
      }
   ]
}
`

// TODO: struct with stream_id, video_id, featured_streamers instead to reduce duplication

// Featured users for live streams
const InitialStreamIDFeaturedMapStr = `
{
   "40033356299": [
      "QTCinderella",
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
      "Sodapoppin",
      "WillNeff",
      "xQcOW",
      "PhinTTV"
   ],
   "43886694045": [
      "Ludwig",
      "Adeptthebest",
      "AustinShow",
      "ConnorEatsPants",
      "Cyr",
      "EsfandTV",
      "HasanAbi",
      "Jschlatt ",
      "JustaMinx",
      "Kaceytron",
      "Malena",
      "Nmplol",
      "Myth",
      "QTCinderella",
      "Sodapoppin",
      "WillNeff",
      "xQcOW"
   ],
   "43884198765": [
       "Myth",
       "QTCinderella",
       "EsfandTV"
   ],
   "40029215963": [
      "Sodapoppin",
      "QTCinderella",
      "Kaceytron",
      "Adeptthebest",
      "JustaMinx",
      "Malena",
      "Ludwig",
      "HasanAbi",
      "Myth",
      "Cyr",
      "PhinTTV"
   ],
   "40028465115": [
      "Ludwig",
      "xQcOW",
      "Sodapoppin",
      "Cyr"
   ],
   "43393194316": [
      "EsfandTV",
      "Nmplol",
      "WillNeff",
      "AustinShow",
      "Myth",
      "Adeptthebest",
      "PhinTTV"
   ],
   "43875063085": [
      "Kaceytron",
      "JustaMinx"
   ],
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

// Featured users for vods
const InitialVodIDFeaturedMapStr = `
{
   "1163160816": [
      "QTCinderella",
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
      "Sodapoppin",
      "WillNeff",
      "xQcOW",
      "PhinTTV"
   ],
   "1162803783": [
      "Ludwig",
      "Adeptthebest",
      "AustinShow",
      "ConnorEatsPants",
      "Cyr",
      "EsfandTV",
      "HasanAbi",
      "Jschlatt ",
      "JustaMinx",
      "Kaceytron",
      "Malena",
      "Nmplol",
      "Myth",
      "QTCinderella",
      "Sodapoppin",
      "WillNeff",
      "xQcOW",
      "PhinTTV"
   ],
   "1162576914": [
      "Myth",
      "QTCinderella",
      "EsfandTV",
      "Nmplol",
      "Malena",
      "AustinShow",
      "PhinTTV"
   ],
   "1162277299": [
      "Sodapoppin",
      "QTCinderella",
      "Kaceytron",
      "Adeptthebest",
      "JustaMinx",
      "Malena",
      "Ludwig",
      "xQcOW",
      "Nmplol",
      "EsfandTV",
      "AustinShow",
      "Myth",
      "Cyr",
      "HasanAbi",
      "WillNeff",
      "PhinTTV"
   ],
   "1162103704": [
      "Ludwig",
      "xQcOW",
      "Sodapoppin",
      "Cyr"
   ],
   "1162074919": [
      "EsfandTV",
      "Nmplol",
      "Myth",
      "WillNeff",
      "AustinShow",
      "PhinTTV"
   ],
   "1161967843": [
      "Kaceytron",
      "JustaMinx",
      "QTCinderella"
   ],
   "1161931874": [
      "Kaceytron",
      "JustaMinx"
   ],
   "1161828036": [
      "HasanAbi",
      "xQcOW",
      "Ludwig",
      "Sodapoppin",
      "EsfandTV",
      "Myth",
      "Cyr",
      "WillNeff",
      "Adeptthebest",
      "AustinShow",
      "PhinTTV"
   ],
   "1161751195": [
      "EsfandTV",
      "QTCinderella",
      "AustinShow",
      "Ludwig",
      "Myth",
      "WillNeff",
      "PhinTTV"
   ],
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
