package service

import (
	"initoko/module/entities"
	"initoko/module/feature/kategori"
	"initoko/utils"
)

type KategoriService struct {
	repo kategori.KategoriRepositoryInterface
}

func NewKategoriService(repo kategori.KategoriRepositoryInterface) kategori.KategoriServiceInterface {
	return &KategoriService{
		repo: repo,
	}
}

func (s *KategoriService) CreateKategori(newData *entities.KategoriModels, foto interface{}, fotoname string) (*entities.KategoriModels, error) {
	fotourl, err := utils.ImageHandler(foto, fotoname, "kategori")
	if err != nil {
		return nil, err
	}

	value := &entities.KategoriModels{
		Name: newData.Name,
		Tipe: newData.Tipe,
		Foto: fotourl,
	}

	res, err2 := s.repo.CreateKategori(value)
	if err2 != nil {
		return nil, err2
	}

	return res, nil
}
func (s *KategoriService) UpdateKategori(id uint64, newData *entities.KategoriModels, foto interface{}, fotoname string) (*entities.KategoriModels, error) {
	fotourl, err := utils.ImageHandler(foto, fotoname, "kategori")
	if err != nil {
		return nil, err
	}

	value := &entities.KategoriModels{
		Name: newData.Name,
		Tipe: newData.Tipe,
		Foto: fotourl,
	}

	res, err2 := s.repo.UpdateKategori(id, value)
	if err2 != nil {
		return nil, err2
	}

	return res, nil
}
func (s *KategoriService) DeleteKategori(id uint64) error {
	err := s.repo.DeleteKategori(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *KategoriService) GetKategoriByID(id uint64) (*entities.KategoriModels, error) {
	res, err := s.repo.GetKategoriByID(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *KategoriService) GetAllKategori() ([]*entities.KategoriModels, error) {
	res, err := s.repo.GetAllKategori()
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *KategoriService) GetAllKategoriName() ([]*entities.KategoriModels, error) {
	res, err := s.repo.GetAllKategoriName()
	if err != nil {
		return nil, err
	}
	return res, nil
}
