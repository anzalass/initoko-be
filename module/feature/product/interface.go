package product

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type ProductRepositoryInterface interface {
	CreateProduct(newData *entities.ProductModels) (*entities.ProductModels, error)
	UpdateProduct(id uint64, newData *entities.ProductModels) (*entities.ProductModels, error)
	DeleteProduct(id uint64) error
	GetProductByID(id uint64) (*entities.ProductModels, error)
	GetAllProduct() ([]*entities.ProductModels, error)
	UploadFotoProduct(newData *entities.FotoProductModels) (bool, error)
	GetAllProductByFilter(min uint64, max uint64, cari string, kategori string, filter string) ([]*entities.ProductModels, error)
	GetProductByKategori(kategori string) ([]*entities.ProductModels, error)
	GetProductDiscount(limit uint64) ([]*entities.ProductModels, error)
	GetProductHomepage() ([]*entities.ProductModels, error)
	GetFotoProductById(id uint64) (string, error)
}
type ProductServiceInterface interface {
	CreateProduct(newData *entities.ProductModels) (*entities.ProductModels, error)
	UpdateProduct(id uint64, newData *entities.ProductModels) (*entities.ProductModels, error)
	DeleteProduct(id uint64) error
	GetProductByID(id uint64) (*entities.ProductModels, error)
	GetAllProduct() ([]*entities.ProductModels, error)
	UploadFotoProduct(newData *entities.FotoProductModels, idproduct uint64) (bool, error)
	GetAllProductByFilter(min uint64, max uint64, cari string, kategori string, filter string) ([]*entities.ProductModels, error)
	Homepage() (any, any, any, any)
	GetFotoProductById(id uint64) (string, error)
	GetProductByKategori(kategori string) ([]*entities.ProductModels, error)
}
type ProductHandlerInterface interface {
	CreateProduct() echo.HandlerFunc
	UpdateProduct() echo.HandlerFunc
	DeleteProduct() echo.HandlerFunc
	GetProductByID() echo.HandlerFunc
	UploadFotoProduct() echo.HandlerFunc
	GetAllProductByFilter() echo.HandlerFunc
	GetAllProduct() echo.HandlerFunc
	Homepage() echo.HandlerFunc
	GetProductByKategori() echo.HandlerFunc
	// GetFotoProductById() echo.HandlerFunc
}
