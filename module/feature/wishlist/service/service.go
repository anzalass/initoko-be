package service

import (
	"initoko/module/entities"
	"initoko/module/feature/wishlist"
)

type WishlistService struct {
	repo wishlist.RepositoryWishlistInterface
}

func NewWishlistService(repo wishlist.RepositoryWishlistInterface) wishlist.ServiceWishlistInterface {
	return &WishlistService{
		repo: repo,
	}
}

func (s *WishlistService) AddWishlist(newData *entities.WishlistModels) (*entities.WishlistModels, error) {
	var res *entities.WishlistModels
	var err error
	cekwishlist, _ := s.repo.GetWishListByIDProductAndEmail(newData.IDProduct, newData.Email)
	if cekwishlist == true {
		value := &entities.WishlistModels{
			IDUser:    newData.IDUser,
			IDProduct: newData.IDProduct,
			Email:     newData.Email,
			Harga:     newData.Harga,
			Name:      newData.Name,
			Foto:      newData.Foto,
		}
		res, err = s.repo.AddWishlist(value)
		if err != nil {
			return nil, err
		}
	}

	return res, nil
}
func (s *WishlistService) GetWishlistByIdUser(id string) ([]*entities.WishlistModels, error) {
	res, err := s.repo.GetWishlistByIdUser(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *WishlistService) DeleteWishListById(id uint64) error {
	err := s.repo.DeleteWishListById(id)
	if err != nil {
		return err
	}
	return nil
}
