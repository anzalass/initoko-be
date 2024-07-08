package users

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryUserInterface interface {
	RegisterUser(newData *entities.UsersModels) (*entities.UsersModels, error)
	LoginUser(email, password string) (*entities.UsersModels, error)
	UpdateUser(id string, newData *entities.UsersModels) (*entities.UsersModels, error)
	DeleteUser(id uint64) error
	CreateOTP(newData *entities.OtpModels) error
	DeleteOTPByEmail(email string) error
	CompareOTPByEmail(kode string, email string) bool
	ChangeStatusByEmail(email string) error
	GetAkunByEmail(email string) (bool, error)
	GetAkunByEmailProfile(email string) (*entities.UsersModels, error)
	GetAlamatByEmailProfile(email string) ([]*entities.AlamatPenerimaModels, error)
}

type ServiceUserInterface interface {
	RegisterUser(newData *entities.UsersModels) (*entities.UsersModels, error)
	RegisterUserGoogle(newData *entities.UsersModels) (*entities.UsersModels, error)
	LoginUser(email, password string) (any, error)
	UpdateUser(id string, newData *entities.UsersModels) (*entities.UsersModels, error)
	UpdateAvatar(id string, file interface{}, name string) (bool, error)
	DeleteUser(id uint64) error
	AktifasiAkun(kode string, email string) (bool, error)
	SendOtp(email string) error
	ProfilePage(email string) (*entities.UsersModels, error)
	AlamatPage(email string) ([]*entities.AlamatPenerimaModels, error)
}

type HandlerUserInterface interface {
	RegisterUser() echo.HandlerFunc
	RegisterUserGoogle() echo.HandlerFunc
	AktivasiAkun() echo.HandlerFunc
	LoginUser() echo.HandlerFunc
	UpdateUser() echo.HandlerFunc
	DeleteUser() echo.HandlerFunc
	UpdateAvatar() echo.HandlerFunc
	SendOtp() echo.HandlerFunc
	ProfilePage() echo.HandlerFunc
	AlamatPage() echo.HandlerFunc
}
