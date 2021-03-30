package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PingEndpoint(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
