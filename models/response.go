package models

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response Base response
type Response struct {
	StatusCode int            `json:"-"`
	Success    bool           `json:"success"`
	Message    string         `json:"message,omitempty"`
	Data       map[string]any `json:"data,omitempty"`
}

func (response *Response) SendResponse(ginContext *gin.Context) {
	ginContext.AbortWithStatusJSON(response.StatusCode, response)
}

func SendResponseData(ginContext *gin.Context, data gin.H) {
	response := &Response{
		StatusCode: http.StatusOK,
		Success:    true,
		Data:       data,
	}
	response.SendResponse(ginContext)
}

func SendErrorResponse(c *gin.Context, status int, message string) {
	response := &Response{
		StatusCode: status,
		Success:    false,
		Message:    message,
	}
	response.SendResponse(c)
}
