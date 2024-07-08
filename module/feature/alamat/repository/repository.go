package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/alamat"

	"gorm.io/gorm"
)

type AlamatRepository struct {
	db *gorm.DB
}

func NewAlamatRepository(db *gorm.DB) alamat.RepositoryAlamatInterface {
	return &AlamatRepository{db: db}
}

func (r *AlamatRepository) CreateAlamat(newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *AlamatRepository) UpdateAlamat(id uint64, newData *entities.AlamatPenerimaModels) (*entities.AlamatPenerimaModels, error) {
	alamat := &entities.AlamatPenerimaModels{}
	if err := r.db.Model(alamat).Where("id = ? AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return alamat, nil
}
func (r *AlamatRepository) DeleteAlamatById(id uint64) error {
	alamat := &entities.AlamatPenerimaModels{}
	if err := r.db.Model(alamat).Where("id = ? AND deleted_at IS NULL", id).Delete(alamat).Error; err != nil {
		return err
	}
	return nil
}
func (r *AlamatRepository) GetAlamatById(id uint64) (*entities.AlamatPenerimaModels, error) {
	alamat := &entities.AlamatPenerimaModels{}
	if err := r.db.Where("id = ? AND  deleted_at IS NULL", id).First(alamat).Error; err != nil {
		return nil, err
	}
	return alamat, nil
}
func (r *AlamatRepository) GetAlamatByIdUser(id uint64) ([]*entities.AlamatPenerimaModels, error) {
	var alamat []*entities.AlamatPenerimaModels
	if err := r.db.Where("id = ? AND deleted_at IS NULL", id).Find(&alamat).Error; err != nil {
		return nil, err
	}
	return alamat, nil
}
