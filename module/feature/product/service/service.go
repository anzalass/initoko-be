package service

import (
	"initoko/module/entities"
	"initoko/module/feature/jumbotron"
	"initoko/module/feature/kategori"
	"initoko/module/feature/product"
)

type ProductsService struct {
	repo          product.ProductRepositoryInterface
	repojumbotron jumbotron.JumbotronRepositoryInterface
	repokategori  kategori.KategoriRepositoryInterface
}

func NewProductService(repo product.ProductRepositoryInterface, repojumbotron jumbotron.JumbotronRepositoryInterface, repokategori kategori.KategoriRepositoryInterface) product.ProductServiceInterface {
	return &ProductsService{
		repo:          repo,
		repojumbotron: repojumbotron,
		repokategori:  repokategori,
	}
}

func (s *ProductsService) CreateProduct(newData *entities.ProductModels) (*entities.ProductModels, error) {
	value := &entities.ProductModels{
		Name:       newData.Name,
		Harga:      newData.Harga,
		Diskon:     newData.Diskon,
		Kategori:   newData.Kategori,
		Deskripsi:  newData.Deskripsi,
		Tags:       newData.Tags,
		Ratings:    0,
		Stok:       newData.Stok,
		Terjual:    0,
		DibuatOleh: newData.DibuatOleh,
	}

	res, err := s.repo.CreateProduct(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *ProductsService) UpdateProduct(id uint64, newData *entities.ProductModels) (*entities.ProductModels, error) {
	value := &entities.ProductModels{
		Name:      newData.Name,
		Harga:     newData.Harga,
		Diskon:    newData.Diskon,
		Kategori:  newData.Kategori,
		Deskripsi: newData.Deskripsi,
		Tags:      newData.Tags,
	}
	res, err := s.repo.UpdateProduct(id, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *ProductsService) DeleteProduct(id uint64) error {
	err := s.repo.DeleteProduct(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *ProductsService) GetProductByID(id uint64) (*entities.ProductModels, error) {
	res, err := s.repo.GetProductByID(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *ProductsService) GetAllProduct() ([]*entities.ProductModels, error) {
	res, err := s.repo.GetAllProduct()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *ProductsService) UploadFotoProduct(newData *entities.FotoProductModels, idproduct uint64) (bool, error) {
	value := &entities.FotoProductModels{
		IDProduct: idproduct,
		Url:       newData.Url,
	}
	res, err := s.repo.UploadFotoProduct(value)
	if err != nil {
		return false, err
	}

	return res, nil
}
func (s *ProductsService) GetAllProductByFilter(min uint64, max uint64, cari string, kategori string, filter string) ([]*entities.ProductModels, error) {
	res, err := s.repo.GetAllProductByFilter(min, max, cari, kategori, filter)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *ProductsService) Homepage() (any, any, any, any) {
	jumbotron, err := s.repojumbotron.GetJumbotronAktif()
	if err != nil {
		return err, nil, nil, nil
	}

	kategori, err := s.repokategori.GetKategoriInti()
	if err != nil {
		return err, nil, nil, nil
	}

	productdiscount, err := s.repo.GetProductDiscount(9)
	if err != nil {
		return err, nil, nil, nil
	}

	product, err := s.repo.GetProductHomepage()
	if err != nil {
		return err, nil, nil, nil
	}

	return jumbotron, kategori, productdiscount, product
}

func (s *ProductsService) GetFotoProductById(id uint64) (string, error) {
	res, err := s.repo.GetFotoProductById(id)
	if err != nil {
		return "", err
	}

	return res, nil
}

func (s *ProductsService) GetProductByKategori(kategori string) ([]*entities.ProductModels, error) {
	res, err := s.repo.GetProductByKategori(kategori)
	if err != nil {
		return nil, err
	}
	return res, nil
}
