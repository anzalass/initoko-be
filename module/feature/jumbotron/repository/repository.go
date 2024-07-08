package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/jumbotron"

	"gorm.io/gorm"
)

type JumbotronRepository struct {
	db *gorm.DB
}

func NewJumbotronRepository(db *gorm.DB) jumbotron.JumbotronRepositoryInterface {
	return &JumbotronRepository{db: db}
}

func (r *JumbotronRepository) CreateJumbotron(newData *entities.JumbotronModels) (*entities.JumbotronModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *JumbotronRepository) UpdateJumbotron(id uint64, newData *entities.JumbotronModels) (*entities.JumbotronModels, error) {
	jumbotron := &entities.JumbotronModels{}
	if err := r.db.Model(jumbotron).Where("id = ?", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return jumbotron, nil
}
func (r *JumbotronRepository) DeleteJumbotron(id uint64) error {
	jumbotron := &entities.JumbotronModels{}
	if err := r.db.Where("id=?", id).Delete(jumbotron).Error; err != nil {
		return err
	}
	return nil
}
func (r *JumbotronRepository) GetJumbotronById(id uint64) (*entities.JumbotronModels, error) {
	jumbotron := &entities.JumbotronModels{}
	if err := r.db.Where("id =? ", id).First(jumbotron).Error; err != nil {
		return nil, err
	}
	return jumbotron, nil
}
func (r *JumbotronRepository) GetAllJumbotron() ([]*entities.JumbotronModels, error) {
	var jumbotron []*entities.JumbotronModels
	if err := r.db.Find(&jumbotron).Error; err != nil {
		return nil, err
	}

	return jumbotron, nil
}
func (r *JumbotronRepository) GetJumbotronAktif() ([]*entities.JumbotronModels, error) {
	var jumbotron []*entities.JumbotronModels
	if err := r.db.Find(&jumbotron).Error; err != nil {
		return nil, err
	}

	return jumbotron, nil
}
