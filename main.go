package main

import (
	"net/http"

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
	return router
}

func main() {
	CreateServer().Run()
}
