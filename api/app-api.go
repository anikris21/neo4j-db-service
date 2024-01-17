package api

import (
	"main/controller"
	"main/dto"
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
// @Summary List existing apps
// @Description Get all the existing apps
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
// @Tags apps,create
// @Accept  json
// @Produce  json
// @Param video body entity.App true "Create App"
// @Success 200 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 401 {object} dto.Response
// @Router /apps [post]
func(api *AppApi) CreateApp(ctx *gin.Context) {
	err := api.appController.Save(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &dto.Response{
			Message: err.Error(),  
		})
	} else {
		ctx.JSON(http.StatusOK, &dto.Response{
			Message: "Success!",  
		})
	}
} 