package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/product"
	"time"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) product.ProductRepositoryInterface {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) CreateProduct(newData *entities.ProductModels) (*entities.ProductModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *ProductRepository) UpdateProduct(id uint64, newData *entities.ProductModels) (*entities.ProductModels, error) {
	products := &entities.ProductModels{}
	if err := r.db.Model(products).Where("id = ?  AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return products, nil
}
func (r *ProductRepository) DeleteProduct(id uint64) error {
	products := &entities.ProductModels{}
	if err := r.db.Model(products).Where("id = ?  AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}
	return nil
}
func (r *ProductRepository) GetProductByID(id uint64) (*entities.ProductModels, error) {
	product := entities.ProductModels{}
	if err := r.db.Where("id = ?  AND deleted_at IS NULL", id).Preload("FotoProduct").First(&product).Error; err != nil {
		return nil, err
	}
	return &product, nil
	// product := entities.ProductModels{}
	// if err := r.db.Where("id = ? AND deleted_at IS NULL", id).
	// 	Preload("FotoProduct", func(db *gorm.DB) *gorm.DB {
	// 		return db.Select("url")
	// 	}).
	// 	First(&product).Error; err != nil {
	// 	return nil, err
	// }
	// return &product, nil
}
func (r *ProductRepository) GetAllProduct() ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	if err := r.db.Where("deleted_at IS NULL").Preload("FotoProduct").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil

}
func (r *ProductRepository) UploadFotoProduct(newData *entities.FotoProductModels) (bool, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return false, err
	}

	return true, nil
}
func (r *ProductRepository) GetAllProductByFilter(min uint64, max uint64, cari string, kategori string, filter string) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	if min == 0 {
		min = 1
	}

	if max == 0 {
		max = 999999999
	}

	query := r.db.Where("harga >= ? AND harga <= ?  AND deleted_at IS NULL", min, max)

	if cari != "" {
		query = query.Where("name ILIKE ?", "%"+cari+"%")
	}
	if filter == "laris" {
		query = query.Order("terjual DESC")
	} else if filter == "kurang laris" {
		query = query.Order("terjual ASC")
	} else if filter == "rating tertinggi" {
		query = query.Order("ratings DESC")
	} else if filter == "rating terendah" {
		query = query.Order("ratings ASC")
	} else if filter == "terbaru" {
		query = query.Order("created_at DESC")
	} else if filter == "terlama" {
		query = query.Order("created_at ASC")
	} else if filter == "diskon terendah" {
		query = query.Order("diskon ASC")
	} else if filter == "diskon tertinggi" {
		query = query.Order("diskon DESC")
	}

	if kategori != "" {
		query = query.Where("kategori = ?", kategori)
	}

	if err := query.Preload("FotoProduct").Find(&products).Limit(10).Error; err != nil {
		return nil, err
	}

	return products, nil

}

func (r *ProductRepository) GetProductByKategori(kategori string) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	if err := r.db.Where("kategori = ? AND deleted_at IS NULL", kategori).Preload("FotoProduct").Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetProductDiscount(limit uint64) ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	if err := r.db.Where("diskon >= 10  AND deleted_at IS NULL").Preload("FotoProduct").Limit(int(limit)).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
func (r *ProductRepository) GetProductHomepage() ([]*entities.ProductModels, error) {
	var products []*entities.ProductModels
	if err := r.db.Where("deleted_at IS NULL").Preload("FotoProduct").Order("terjual DESC").Limit(50).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}

func (r *ProductRepository) GetFotoProductById(id uint64) (string, error) {
	foto := &entities.FotoProductModels{}
	if err := r.db.Where("id_product", id).First(foto).Error; err != nil {
		return "", err
	}

	return foto.Url, nil
}
