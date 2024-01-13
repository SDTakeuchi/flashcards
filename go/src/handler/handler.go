package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func respSuccess(c *gin.Context, data any) {
	c.JSON(200, gin.H{
		"data": data,
	})
}

func respError(c *gin.Context, statusCode int, msg ...string) {
	var errMsg string
	if len(msg) >= 1 {
		errMsg = msg[0]
	} else {
		errMsg = http.StatusText(statusCode)
	}
	c.JSON(statusCode, gin.H{
		"message": errMsg,
	})
}
