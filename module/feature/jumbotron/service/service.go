package service

import (
	"initoko/module/entities"
	"initoko/module/feature/jumbotron"
	"initoko/utils"
)

type JumbotronService struct {
	repo jumbotron.JumbotronRepositoryInterface
}

func NewJumbotronService(repo jumbotron.JumbotronRepositoryInterface) jumbotron.JumbotronServiceInterface {
	return &JumbotronService{repo: repo}
}

func (s *JumbotronService) CreateJumbotron(newData *entities.JumbotronModels, foto interface{}, fotoname string) (*entities.JumbotronModels, error) {
	fotourl, err := utils.ImageHandler(foto, fotoname, "jumbotron")
	if err != nil {
		return nil, err
	}

	value := &entities.JumbotronModels{
		Name:      newData.Name,
		Deskripsi: newData.Deskripsi,
		Foto:      fotourl,
	}

	res, err2 := s.repo.CreateJumbotron(value)
	if err2 != nil {
		return nil, err2
	}

	return res, nil
}
func (s *JumbotronService) UpdateJumbotron(id uint64, newData *entities.JumbotronModels, foto interface{}, fotoname string) (*entities.JumbotronModels, error) {
	fotourl, err := utils.ImageHandler(foto, fotoname, "jumbotron")
	if err != nil {
		return nil, err
	}

	value := &entities.JumbotronModels{
		Name:      newData.Name,
		Deskripsi: newData.Deskripsi,
		Foto:      fotourl,
	}

	res, err2 := s.repo.UpdateJumbotron(id, value)
	if err2 != nil {
		return nil, err2
	}

	return res, nil
}
func (s *JumbotronService) DeleteJumbotron(id uint64) error {
	err := s.repo.DeleteJumbotron(id)
	if err != nil {
		return err
	}
	return nil
}
func (s *JumbotronService) GetJumbotronById(id uint64) (*entities.JumbotronModels, error) {
	res, err := s.repo.GetJumbotronById(id)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *JumbotronService) GetAllJumbotron() ([]*entities.JumbotronModels, error) {
	res, err := s.repo.GetAllJumbotron()
	if err != nil {
		return nil, err
	}
	return res, nil
}
