package main

import "github.com/gin-gonic/gin"

func main3() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "zet",
		})
	})
	r.Run() // 0.0.0.0:8080 でサーバーを立てます。
}
