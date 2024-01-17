package controller

import (
	"main/entity"
	"main/service"

	"github.com/gin-gonic/gin"
)


type AppController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.App
}

type controller struct {
	service service.AppService 
}

func New(service service.AppService) AppController{

	return &controller{service: service} 
}

func(contr *controller) Save(ctx *gin.Context)error {
	var app entity.App
	err := ctx.ShouldBindJSON(&app)
	if(err != nil) {
		return err
	}
	contr.service.Save(app)
	return nil
}

func(contr *controller) FindAll()[]entity.App {
	return contr.service.FindAll()
}  

