package server

import (
	"github.com/fellipegpbotelho/go-rest-api/endpoints/login"
	"github.com/fellipegpbotelho/go-rest-api/endpoints/ping"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", ping.PingEndpoint)
	router.POST("/login", login.LoginEndpoint)
	return router
}
