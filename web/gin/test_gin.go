package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 静态资源加载，本例为css,js以及资源图片
	router.StaticFS("/public", http.Dir("./web/"))
	//router.StaticFile("/favicon.ico", "./resources/favicon.ico")
	router.GET("/test/Ping", Ping)
	router.GET("/movies/get", GetMovie)
	router.Run(":8888")
}

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func GetMovie(c *gin.Context) {
	id := c.Query("id")
	c.JSON(200, gin.H{
		"message": "GetMovie",
		"id":      id,
	})
}
