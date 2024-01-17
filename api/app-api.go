package api

import (
	"main/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AppApi struct {
	appController controller.AppController
}


func NewAppAPI(appController controller.AppController) *AppApi {
	return &AppApi{
		appController: appController,
	}
}


// GetApps godoc
// @Security bearerAuth
// @Summary List existing videos
// @Description Get all the existing videos
// @Tags apps,list
// @Accept  json
// @Produce  json
// @Success 200 {array} entity.App
// @Router /apps [get]

func(api *AppApi) GetApps(ctx *gin.Context) {
	ctx.JSON(200, api.appController.FindAll())
	
}

// CreateApp godoc
// @Security bearerAuth
// @Summary Create new apps
// @Description Create a new app
// @Tags videos,create
// @Accept  json
// @Produce  json
// @Param video body entity.Video true "Create video"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 401 {object} gin.H
// @Router /apps [post]
 
func(api *AppApi) CreateApp(ctx *gin.Context) {
	err := api.appController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),  
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Success!",  
		})
	}
} 