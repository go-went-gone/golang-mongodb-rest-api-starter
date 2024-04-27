package routes

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/controllers"
	"github.com/gin-gonic/gin"
)

func PingRoute(ginRouterGroup *gin.RouterGroup) {
	ping := ginRouterGroup.Group("/ping")
	{
		ping.GET(
			"",
			controllers.Ping,
		)
	}
}
