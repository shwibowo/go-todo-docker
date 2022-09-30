package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	fmt.Println("Hello World")
	route := gin.Default()
	route.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	db := NewConfig()
	db.Migration(&Todo{})

	route.GET("/todos", func(ctx *gin.Context) {
		var todo []Todo
		if err := db.Db.Find(&todo).Error; err != nil {
			ctx.AbortWithStatus(404)
			fmt.Println(err)
			return
		} else {
			ctx.JSON(200, gin.H{
				"message": "OK",
				"data":    todo,
			})
		}
	})

	route.POST("/todos", func(ctx *gin.Context) {
		var todo Todo
		ctx.BindJSON(&todo)
		if err := db.Db.Create(&todo).Error; err != nil {
			ctx.AbortWithStatus(404)
			return
		} else {
			ctx.JSON(200, gin.H{
				"message": "Create todo success",
				"data":    todo,
			})
		}
	})

	port := os.Getenv("API_PORT")
	err := route.Run(":" + port)
	if err != nil {
		return
	}
}
