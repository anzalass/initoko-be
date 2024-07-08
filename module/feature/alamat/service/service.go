package service

import (
	"initoko/module/entities"
	"initoko/module/feature/alamat"
)

type AlamatService struct {
	repo alamat.RepositoryAlamatInterface
}

func NewAlamatService(repo alamat.RepositoryAlamatInterface) alamat.ServiceAlamatInterface {
	return &AlamatService{repo: repo}
}

func (s *AlamatService) CreateAlamat(newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error) {
	value := &entities.AlamatPenerimaModels{
		IDUser:        newData.IDUser,
		Name:          newData.Name,
		Email:         newData.Email,
		Desa:          newData.Desa,
		Kecamatan:     newData.Kecamatan,
		Kabupaten:     newData.Kabupaten,
		Provinsi:      newData.Provinsi,
		NoWhatsapp:    newData.NoWhatsapp,
		AlamatLengkap: newData.AlamatLengkap,
	}
	res, err := s.repo.CreateAlamat(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *AlamatService) UpdateAlamat(id uint64, newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error) {
	value := &entities.AlamatPenerimaModels{
		IDUser:        newData.IDUser,
		Name:          newData.Name,
		Desa:          newData.Desa,
		Kecamatan:     newData.Kecamatan,
		Kabupaten:     newData.Kabupaten,
		Provinsi:      newData.Provinsi,
		NoWhatsapp:    newData.NoWhatsapp,
		AlamatLengkap: newData.AlamatLengkap,
	}
	res, err := s.repo.UpdateAlamat(id, value)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *AlamatService) DeleteAlamatById(id uint64) error {
	err := s.repo.DeleteAlamatById(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *AlamatService) GetAlamatById(id uint64) (*entities.AlamatPenerimaModels, error) {
	res, err := s.repo.GetAlamatById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *AlamatService) GetAlamatByIdUser(id uint64) ([]*entities.AlamatPenerimaModels, error) {
	res, err := s.repo.GetAlamatByIdUser(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
