package routes

import (
	controller "github.com/OBAIDULLAHKHANKHAIL/golang-jwt-project/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(incommingRoutes *gin.Engine) {
	incommingRoutes.POST("user/signup", controller.Signup())
	incommingRoutes.POST("user/login", controller.Login())
}
