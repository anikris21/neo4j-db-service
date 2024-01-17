package controller

import (
	"main/entity"
	"main/service"

	"github.com/gin-gonic/gin"
)


type AppController interface {
	Save(ctx *gin.Context) entity.App
	FindAll() []entity.App
}

type controller struct {
	service service.AppService 
}

func New(service service.AppService) AppController{

	return &controller{service: service} 
}

func(contr *controller) Save(ctx *gin.Context)entity.App {
	var app entity.App
	ctx.BindJSON(&app)
	contr.service.Save(app)
	return app
}

func(contr *controller) FindAll()[]entity.App {
	return contr.service.FindAll()
}  

