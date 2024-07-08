package wishlist

import (
	"initoko/module/entities"

	"github.com/labstack/echo/v4"
)

type RepositoryWishlistInterface interface {
	AddWishlist(newData *entities.WishlistModels) (*entities.WishlistModels, error)
	GetWishlistByIdUser(id string) ([]*entities.WishlistModels, error)
	DeleteWishListById(id uint64) error
	GetWishListByIDProductAndEmail(idproduct uint64, email string) (bool, error)
}
type ServiceWishlistInterface interface {
	AddWishlist(newData *entities.WishlistModels) (*entities.WishlistModels, error)
	GetWishlistByIdUser(id string) ([]*entities.WishlistModels, error)
	DeleteWishListById(id uint64) error
}
type HandlerWishlistInterface interface {
	AddWishlist() echo.HandlerFunc
	GetWishlistByIdUser() echo.HandlerFunc
	DeleteWishListById() echo.HandlerFunc
}
