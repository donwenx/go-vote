package router

import (
	"vote/config"
	"vote/controllers"

	"github.com/gin-contrib/sessions"
	sessions_redis "github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	store, _ := sessions_redis.NewStore(10, "tcp", config.RedisAddress, "", []byte("secret"))
	router.Use(sessions.Sessions("mysession", store))

	// 用户
	user := router.Group("/user")
	{
		user.POST("/register", controllers.UserController{}.Register)
		user.POST("login", controllers.UserController{}.Login)
	}

	player := router.Group("/player")
	{
		player.POST("/list", controllers.PlayerController{}.GetPlayers)
		player.POST("/create", controllers.PlayerController{}.CreatePlayers)
	}
	return router
}
