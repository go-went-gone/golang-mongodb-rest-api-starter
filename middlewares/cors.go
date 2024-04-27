package middlewares

import (
	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		ginContext.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ginContext.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		ginContext.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
		ginContext.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")

		ginContext.Next()
	}
}
