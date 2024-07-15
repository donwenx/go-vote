package router

import (
	"vote/controllers"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 用户
	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
	}
	return router
}
