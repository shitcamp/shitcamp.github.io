package server

import (
	"github.com/gin-gonic/gin"
)

const (
	gPayment = "/twitch"
)

func newRouter() *gin.Engine {
	router := gin.Default()

	//router.Use(middleware.AuthenticationMiddleware(gWebhook))

	// Routes
	paymentRouter := router.Group(gPayment)
	{
		//paymentRouter.GET("/get_payment_setup", handlers.GetPaymentSetup)
		//paymentRouter.POST("/create_checkout_session", handlers.CreateCheckoutSession)
	}

	return router
}
