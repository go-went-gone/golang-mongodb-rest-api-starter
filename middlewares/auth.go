package middlewares

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/models"
	db "github.com/ebubekiryigit/golang-mongodb-rest-api-starter/models/db"
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

// authrization, authentication
func JWTMiddleware() gin.HandlerFunc {
	return func(ginContext *gin.Context) {
		token := ginContext.GetHeader("Authroization")
		tokenModel, err := services.VerifyToken(token, db.TokenTypeAccess)
		if err != nil {
			models.SendErrorResponse(ginContext, http.StatusUnauthorized, err.Error())
			return
		}

		ginContext.Set("userIdHex", tokenModel.User.Hex())
		ginContext.Set("userId", tokenModel.User)

		ginContext.Next()
	}
}
