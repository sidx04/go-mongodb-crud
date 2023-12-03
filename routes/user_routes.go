package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mongodb-crud/controllers"
)

func UserRoutes(app *gin.Engine) {
	app.POST("/user", controllers.CreateUser)
	app.GET("/user/:id", controllers.GetUserById)
}
