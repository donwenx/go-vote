package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	// 用户
	user := router.Group("/user")
	{
		user.GET("/", func(c *gin.Context) {
			c.String(http.StatusOK, "test01")
		})
	}
	return router
}
