package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/kategori"
	"time"

	"gorm.io/gorm"
)

type KategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) kategori.KategoriRepositoryInterface {
	return &KategoriRepository{
		db: db,
	}
}

func (r *KategoriRepository) CreateKategori(newData *entities.KategoriModels) (*entities.KategoriModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *KategoriRepository) UpdateKategori(id uint64, newData *entities.KategoriModels) (*entities.KategoriModels, error) {
	kategori := &entities.KategoriModels{}
	if err := r.db.Model(kategori).Where("id = ?  AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return kategori, nil
}
func (r *KategoriRepository) DeleteKategori(id uint64) error {
	kategori := &entities.KategoriModels{}
	if err := r.db.Model(kategori).Where("id = ?  AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}
func (r *KategoriRepository) GetKategoriByID(id uint64) (*entities.KategoriModels, error) {
	kategori := &entities.KategoriModels{}
	if err := r.db.Where("id = ?  AND deleted_at IS NULL", id).First(kategori).Error; err != nil {
		return nil, err
	}

	return kategori, nil
}
func (r *KategoriRepository) GetKategoriInti() ([]*entities.KategoriModels, error) {
	var kategori []*entities.KategoriModels
	if err := r.db.Where("deleted_at IS NULL AND tipe = ?", "inti").Limit(10).Find(&kategori).Error; err != nil {
		return nil, err
	}
	return kategori, nil
}
func (r *KategoriRepository) GetAllKategori() ([]*entities.KategoriModels, error) {
	var kategori []*entities.KategoriModels
	if err := r.db.Where("deleted_at IS NULL").Find(&kategori).Error; err != nil {
		return nil, err
	}

	return kategori, nil
}
func (r *KategoriRepository) GetAllKategoriName() ([]*entities.KategoriModels, error) {
	var kategori []*entities.KategoriModels
	if err := r.db.Where("deleted_at IS NULL").Select("name").Find(&kategori).Error; err != nil {
		return nil, err
	}

	return kategori, nil
}
