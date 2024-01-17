package main

import (
	"main/service"
	"main/controller"
	"main/middleware"
	"main/api"

	"fmt"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" 
)

var (
	appService service.AppService  = service.NewAppService()
	appController controller.AppController = controller.New(appService) 
) 

func main(){

	fmt.Println("Hello world")

	// docs.SwaggerInfo.Title = "Pragmatic Reviews - Video API"
	// docs.SwaggerInfo.Description = "Pragmatic Reviews - Youtube Video API."
	// docs.SwaggerInfo.Version = "1.0"
	// docs.SwaggerInfo.Host = "pragmatic-video-app.herokuapp.com"
	// docs.SwaggerInfo.BasePath = "/api/v1"
	// docs.SwaggerInfo.Schemes = []string{"https"}
 
	
	server := gin.New()

	server.Use(gin.Recovery(), middleware.Logger(), middleware.BasicAuth())

	appAPI := api.NewAppAPI(appController) 

	apiRoutes := server.Group("/")
	apps := apiRoutes.Group("/apps")
	{
			apps.GET("", appAPI.GetApps)
			apps.POST("", appAPI.CreateApp)
	}
	
	server.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) 
	
	
	server.GET("test", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Ok!!!",  
		})

	})
	server.Run(":8080")
} 