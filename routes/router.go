package routes

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/docs"
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/middlewares"
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/models"
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/services"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func New() *gin.Engine {
	ginNew := gin.New()
	initRoute(ginNew)

	ginNew.Use(gin.LoggerWithWriter(middlewares.LogWriter()))
	ginNew.Use(gin.CustomRecovery(middlewares.AppRecovery()))
	ginNew.Use(middlewares.CORSMiddleware())

	v1Group := ginNew.Group("/v1")
	{
		PingRoute(v1Group)
		AuthRoute(v1Group)
		NoteRoute(v1Group, middlewares.JWTMiddleware())
	}

	docs.SwaggerInfo.BasePath = v1Group.BasePath() // adds /v1 to swagger base path

	ginNew.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return ginNew
}

func initRoute(r *gin.Engine) {
	_ = r.SetTrustedProxies(nil)
	r.RedirectTrailingSlash = false
	r.HandleMethodNotAllowed = true

	r.NoRoute(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusNotFound, c.Request.RequestURI+" not found")
	})

	r.NoMethod(func(c *gin.Context) {
		models.SendErrorResponse(c, http.StatusMethodNotAllowed, c.Request.Method+" is not allowed here")
	})
}

func InitGin() {
	gin.DisableConsoleColor()
	gin.SetMode(services.Config.Mode)
	// do some other initialization staff
}
