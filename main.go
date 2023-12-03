package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mongodb-crud/configs"
)

func init() {
	configs.ConnectDB()
}

func main() {
	app := gin.Default()

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})

	app.Run()
}
