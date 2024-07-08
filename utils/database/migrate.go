package database

import (
	"initoko/module/entities"

	"gorm.io/gorm"
)

func MigrateDatabase(db *gorm.DB) {
	err := db.AutoMigrate(entities.UsersModels{}, entities.AlamatPenerimaModels{}, entities.ProductModels{}, entities.ReviewModels{}, entities.FotoProductModels{}, entities.TransaksiModels{}, entities.WishlistModels{}, entities.ProductPesananModels{}, entities.KategoriModels{}, entities.JumbotronModels{}, entities.OtpModels{})
	if err != nil {
		return
	}
}
