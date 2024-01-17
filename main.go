package main

import (
	"main/service"
	"main/controller"

	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	appService service.AppService  = service.NewAppService()
	appController controller.AppController = controller.New(appService) 
) 

func main(){

	fmt.Println("Hello world")
	
	server := gin.Default()

	server.GET("/apps", func(ctx *gin.Context) {
		ctx.JSON(200, appController.FindAll())
	})

	server.POST("/apps", func(ctx *gin.Context) {
		ctx.JSON(200, appController.Save(ctx))
	}) 
	
	
	server.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!!!",  
		})

	})
	server.Run(":8080")
} 