package core

import (
	"github.com/gin-gonic/gin"
)

func MyServer() *gin.Engine {
	router := gin.Default()

	user := router.Group("/user")
	{
		user.GET("/", func(c *gin.Context) {})
		user.POST("/signIn", func(c *gin.Context) {})
		user.POST("/signUp", func(c *gin.Context) {})
		user.POST("/exit", func(c *gin.Context) {})
	}
	dish := router.Group("/dish")
	{
		//show all dishes
		dish.GET("/", func(c *gin.Context) {})
		dish.GET("/showone", func(c *gin.Context) {})
		dish.POST("/create", func(c *gin.Context) {})
		dish.PUT("/update", func(c *gin.Context) {})
		dish.DELETE("/delete", func(c *gin.Context) {})
	}
	category := router.Group("/category")
	{
		category.GET("/", func(c *gin.Context) {})
		category.GET("/showone", func(c *gin.Context) {})
		category.POST("/create", func(c *gin.Context) {})
		category.PUT("/rename", func(c *gin.Context) {})
		category.DELETE("/delete", func(c *gin.Context) {})

	}

	return router
}
