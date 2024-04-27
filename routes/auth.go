package routes

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/controllers"
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/middlewares/validators"
	"github.com/gin-gonic/gin"
)

func AuthRoute(ginRouterGroup *gin.RouterGroup) {
	auth := ginRouterGroup.Group("/auth")
	{
		auth.POST(
			"/register",
			validators.RegisterValidator(),
			controllers.Register,
		)

		auth.POST(
			"/login",
			validators.LoginValidator(),
			controllers.Login,
		)

		auth.POST(
			"/refresh",
			validators.RefreshValidator(),
			controllers.Refresh,
		)
	}
}
