package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mongodb-crud/configs"
	"github.com/go-mongodb-crud/routes"
)

func init() {
	configs.ConnectDB()
}

func main() {
	app := gin.Default()

	routes.UserRoutes(app)

	app.Run()
}
