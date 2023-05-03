package route

import (
	controllers "web-api/src/controllers"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func SetupRoutes(router *gin.Engine) {

	/***BASEPATH OF AN API. NOTE:THIS SHOULDN'T BE CHANGED***/
	api := router.Group("/api/v1")

	/***ADD THE ROUTES HERE***/
	api.POST("/greetings", controllers.GreetUser)

	router.Run(viper.GetString("server.port"))
}
