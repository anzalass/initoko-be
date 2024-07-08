package checkout

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
	_ "github.com/labstack/echo/v4"
)

type RepositoryCheckout interface {
	CreateTransaksi(newData *entities.TransaksiModels) (*entities.TransaksiModels, error)
	CreateProductTransaksi(newData *entities.ProductPesananModels) (*entities.ProductPesananModels, error)
	ChangeStatus(idtransaksi string, status string) error
	ChangeStatusPembayaran(idtransaksi string) error
	AllPesananByEmail(email string) ([]*entities.TransaksiModels, error)
	DetailPesananById(id string) (*entities.TransaksiModels, error)
	AllPesanan() ([]*entities.TransaksiModels, error)
	InputResi(id string, newData *entities.TransaksiModels) error
}

type ServiceCheckout interface {
	CreateTransaksi(newData *entities.TransaksiModels) (*entities.TransaksiModels, error)
	CreateProductTransaksi(newData *entities.ProductPesananModels) (*entities.ProductPesananModels, error)
	ChangeStatus(idtransaksi string, status string) error
	ChangeStatusPembayaran(idtransaksi string) error
	AllPesananByEmail(email string) ([]*entities.TransaksiModels, error)
	DetailPesananById(id string) (*entities.TransaksiModels, error)
	AllPesanan() ([]*entities.TransaksiModels, error)
	InputResi(id string, newData *entities.TransaksiModels) error
}

type HandelerCheckout interface {
	CreateTransaksi() echo.HandlerFunc
	ChangeStatus() echo.HandlerFunc
	ChangeStatusPembayaran() echo.HandlerFunc
	AllPesananByEmail() echo.HandlerFunc
	DetailPesananById() echo.HandlerFunc
	AllPesanan() echo.HandlerFunc
	InputResi() echo.HandlerFunc
}
