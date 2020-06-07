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
		user.PATCH("/:username", UpdateUserByUsername)
		user.DELETE("/:username", DeleteUserByUsername)
	}
	r.Run(":3000")
}
