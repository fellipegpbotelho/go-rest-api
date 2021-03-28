package main

import (
	"net/http"

	endpoints "github.com/fellipegpbotelho/go-rest-api/endpoints/login"
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.GET("/ping", Ping)
	router.POST("/login", endpoints.LoginEndpoint)
	return router
}

func main() {
	CreateServer().Run()
}
