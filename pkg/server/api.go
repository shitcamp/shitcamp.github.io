package server

import (
	"github.com/gin-gonic/gin"
	"github.com/shitcamp-unofficial/shitcamp/pkg/server/handlers"
)

const (
	gShitcamp = "/shitcamp"
	gTwitch   = "/twitch"
)

func newRouter(auth gin.Accounts) *gin.Engine {
	router := gin.Default()

	router.Use(gin.BasicAuth(auth))

	api := router.Group("/api")

	// Routes
	shitcampRouter := api.Group(gShitcamp)
	{
		shitcampRouter.GET("/get_streamer_names", handlers.GetStreamerNames)
		shitcampRouter.GET("/get_streamers", handlers.GetStreamers)
		shitcampRouter.GET("/get_schedule", handlers.GetSchedule)
		shitcampRouter.POST("/set_schedule", handlers.SetSchedule)
		shitcampRouter.POST("/set_featured_users_for_vod", handlers.SetFeaturedUsersForVod)
	}

	twitchRouter := api.Group(gTwitch)
	{
		twitchRouter.GET("/get_live_streams", handlers.GetLiveStreams)
		twitchRouter.GET("/get_vods", handlers.GetVods)
		twitchRouter.GET("/get_clips", handlers.GetClips)
	}

	return router
}
