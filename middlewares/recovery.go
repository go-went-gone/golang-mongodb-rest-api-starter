package middlewares

import (
	"github.com/ebubekiryigit/golang-mongodb-rest-api-starter/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// The purpose of the AppRecovery middleware function is to handle panics that occur during the processing of 
// an HTTP request and recover from them gracefully. When a panic occurs in a route handler or another middleware 
// function, Gin will call the recovery function provided by AppRecovery.
func AppRecovery() func(c *gin.Context, recovered interface{}) {
	return func(c *gin.Context, recovered interface{}) {
		if err, ok := recovered.(string); ok {
			models.SendErrorResponse(c, http.StatusInternalServerError, err)
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"success": false}) // recovery failed
	}
}
