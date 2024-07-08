package kategori

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type KategoriRepositoryInterface interface {
	CreateKategori(newData *entities.KategoriModels) (*entities.KategoriModels, error)
	UpdateKategori(id uint64, newData *entities.KategoriModels) (*entities.KategoriModels, error)
	DeleteKategori(id uint64) error
	GetKategoriByID(id uint64) (*entities.KategoriModels, error)
	GetKategoriInti() ([]*entities.KategoriModels, error)
	GetAllKategori() ([]*entities.KategoriModels, error)
	GetAllKategoriName() ([]*entities.KategoriModels, error)
}
type KategoriServiceInterface interface {
	CreateKategori(newData *entities.KategoriModels, foto interface{}, fotoname string) (*entities.KategoriModels, error)
	UpdateKategori(id uint64, newData *entities.KategoriModels, foto interface{}, fotoname string) (*entities.KategoriModels, error)
	DeleteKategori(id uint64) error
	GetKategoriByID(id uint64) (*entities.KategoriModels, error)
	GetAllKategori() ([]*entities.KategoriModels, error)
	GetAllKategoriName() ([]*entities.KategoriModels, error)
}
type KategoriHandlerInterface interface {
	CreateKategori() echo.HandlerFunc
	UpdateKategori() echo.HandlerFunc
	DeleteKategori() echo.HandlerFunc
	GetKategoriByID() echo.HandlerFunc
	GetAllKategori() echo.HandlerFunc
	GetAllKategoriName() echo.HandlerFunc
}
