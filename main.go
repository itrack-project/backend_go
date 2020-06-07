package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//GIN
	r := gin.Default()
	user := r.Group("/user")
	{
		user.GET("/", GetUsers)
		user.GET("/:username", GetUserByUsername)
		user.POST("/", NewUser)
		user.PUT("/:username", UpdateUserByUsername)
		user.DELETE("/:username", DeleteUserByUsername)
	}
	check := r.Group("/check")
	{
		check.GET("/", GetChecks)
		check.GET("/:username", GetChecksByUser)
		check.POST("/", NewCheck)
		check.PUT("/:code", UpdateCheckById)
		check.DELETE("/:code", DeleteCheckById)
	}
	r.Run(":8080")
}
