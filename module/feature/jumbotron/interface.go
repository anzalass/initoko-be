package jumbotron

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type JumbotronRepositoryInterface interface {
	CreateJumbotron(newData *entities.JumbotronModels) (*entities.JumbotronModels, error)
	UpdateJumbotron(id uint64, newData *entities.JumbotronModels) (*entities.JumbotronModels, error)
	DeleteJumbotron(id uint64) error
	GetJumbotronById(id uint64) (*entities.JumbotronModels, error)
	GetAllJumbotron() ([]*entities.JumbotronModels, error)
	GetJumbotronAktif() ([]*entities.JumbotronModels, error)
}
type JumbotronServiceInterface interface {
	CreateJumbotron(newData *entities.JumbotronModels, foto interface{}, fotoname string) (*entities.JumbotronModels, error)
	UpdateJumbotron(id uint64, newData *entities.JumbotronModels, foto interface{}, fotoname string) (*entities.JumbotronModels, error)
	DeleteJumbotron(id uint64) error
	GetJumbotronById(id uint64) (*entities.JumbotronModels, error)
	GetAllJumbotron() ([]*entities.JumbotronModels, error)
}
type JumbotronHandlerInterface interface {
	CreateJumbotron() echo.HandlerFunc
	UpdateJumbotron() echo.HandlerFunc
	DeleteJumbotron() echo.HandlerFunc
	GetJumbotronById() echo.HandlerFunc
	GetAllJumbotron() echo.HandlerFunc
}
