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

	// Routes
	shitcampRouter := router.Group(gShitcamp)
	{
		shitcampRouter.GET("/get_streamer_names", handlers.GetStreamerNames)
		shitcampRouter.GET("/get_streamers", handlers.GetStreamers)
		shitcampRouter.GET("/get_schedule", handlers.GetSchedule)
		shitcampRouter.POST("/set_schedule", handlers.SetSchedule)
		shitcampRouter.GET("/set_featured_steamers_for_vod", handlers.SetFeaturedStreamersForVod)
	}

	twitchRouter := router.Group(gTwitch)
	{
		twitchRouter.GET("/get_live_streams", handlers.GetLiveStreams)
		twitchRouter.GET("/get_vods", handlers.GetVods)
		twitchRouter.GET("/get_clips", handlers.GetClips)
	}

	return router
}
