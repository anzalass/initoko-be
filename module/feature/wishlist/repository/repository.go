package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/wishlist"

	"gorm.io/gorm"
)

type ReposioryWishlist struct {
	db *gorm.DB
}

func NewReposioryWishlist(db *gorm.DB) wishlist.RepositoryWishlistInterface {
	return &ReposioryWishlist{
		db: db,
	}
}

func (r *ReposioryWishlist) AddWishlist(newData *entities.WishlistModels) (*entities.WishlistModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}

	return newData, nil
}
func (r *ReposioryWishlist) GetWishListByIDProductAndEmail(idproduct uint64, email string) (bool, error) {
	wishlist := &entities.WishlistModels{}
	if err := r.db.Where("id_product = ? AND email = ?", idproduct, email).First(wishlist).Error; err != nil {
		return true, err
	}

	return false, nil
}
func (r *ReposioryWishlist) GetWishlistByIdUser(id string) ([]*entities.WishlistModels, error) {
	var wishlist []*entities.WishlistModels
	if err := r.db.Where("id_user = ?", id).Find(&wishlist).Error; err != nil {
		return nil, err
	}
	return wishlist, nil
}
func (r *ReposioryWishlist) DeleteWishListById(id uint64) error {
	wishlist := &entities.WishlistModels{}
	if err := r.db.Where("id = ?", id).Delete(wishlist).Error; err != nil {
		return err
	}
	return nil
}
