package alamat

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryAlamatInterface interface {
	CreateAlamat(newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error)
	UpdateAlamat(id uint64, newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error)
	DeleteAlamatById(id uint64) error
	GetAlamatById(id uint64) (*entities.AlamatPenerimaModels, error)
	GetAlamatByIdUser(id uint64) ([]*entities.AlamatPenerimaModels, error)
}
type ServiceAlamatInterface interface {
	CreateAlamat(newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error)
	UpdateAlamat(id uint64, newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error)
	DeleteAlamatById(id uint64) error
	GetAlamatById(id uint64) (*entities.AlamatPenerimaModels, error)
	GetAlamatByIdUser(id uint64) ([]*entities.AlamatPenerimaModels, error)
}

type HandlerALamatInterface interface {
	CreateAlamat() echo.HandlerFunc
	UpdateAlamat() echo.HandlerFunc
	DeleteAlamatById() echo.HandlerFunc
	GetAlamatById() echo.HandlerFunc
}
