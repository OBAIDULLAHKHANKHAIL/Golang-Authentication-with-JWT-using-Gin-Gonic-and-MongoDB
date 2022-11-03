package routes

import (
	controller "github.com/OBAIDULLAHKHANKHAIL/golang-jwt-project/controllers"
	"github.com/OBAIDULLAHKHANKHAIL/golang-jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.Use(middleware.Authenticate())
	incommingRoutes.GET("/users", controller.GetUsers())
	incommingRoutes.GET("/users/:user_id", controller.GetUser())
}
