package service

import (
	"main/entity"
)

type AppService interface {
	Save(entity.App) entity.App
	FindAll() []entity.App
}

type appService struct {
	apps []entity.App

}


func NewAppService() AppService {
	return &appService{}
}

func(service *appService) Save(app entity.App)entity.App {
	// mimic db 
	service.apps = append(service.apps, app)
	return app

}

func(service *appService) FindAll()[]entity.App{

	// mimic db 
	return service.apps
} 

