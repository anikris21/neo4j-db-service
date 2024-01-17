package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main(){

	fmt.Println("Hello world")
	
	server := gin.Default()
	server.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!!!",  
		})

	})
	server.Run(":8080")
} 