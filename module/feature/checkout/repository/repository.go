package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/checkout"

	"gorm.io/gorm"
)

type CheckoutRepository struct {
	gorm *gorm.DB
}

func NewCheckoutRepository(gorm *gorm.DB) checkout.RepositoryCheckout {
	return &CheckoutRepository{
		gorm: gorm,
	}
}

func (r *CheckoutRepository) CreateTransaksi(newData *entities.TransaksiModels) (*entities.TransaksiModels, error) {
	if err := r.gorm.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *CheckoutRepository) CreateProductTransaksi(newData *entities.ProductPesananModels) (*entities.ProductPesananModels, error) {
	if err := r.gorm.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *CheckoutRepository) ChangeStatus(idtransaksi string, status string) error {
	transaksi := &entities.TransaksiModels{}
	if err := r.gorm.Model(transaksi).Where("id_pembayaran = ?", idtransaksi).Update("status", "dikirim").Error; err != nil {
		return err
	}

	return nil
}

func (r *CheckoutRepository) ChangeStatusPembayaran(idtransaksi string) error {
	transaksi := &entities.TransaksiModels{}
	if err := r.gorm.Model(transaksi).Where("id = ?", idtransaksi).Update("status_pembayaran", "sudah bayar").Error; err != nil {
		return err
	}

	return nil
}

func (r *CheckoutRepository) AllPesananByEmail(email string) ([]*entities.TransaksiModels, error) {
	var transaksi []*entities.TransaksiModels
	if err := r.gorm.Where("email = ?", email).Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}
func (r *CheckoutRepository) DetailPesananById(id string) (*entities.TransaksiModels, error) {
	var transaksi *entities.TransaksiModels
	if err := r.gorm.Where("id = ?", id).Preload("Product").First(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}
func (r *CheckoutRepository) AllPesanan() ([]*entities.TransaksiModels, error) {
	var transaksi []*entities.TransaksiModels
	if err := r.gorm.Find(&transaksi).Error; err != nil {
		return nil, err
	}
	return transaksi, nil
}

func (r *CheckoutRepository) InputResi(id string, newData *entities.TransaksiModels) error {
	transaksi := &entities.TransaksiModels{}
	if err := r.gorm.Model(transaksi).Where("id = ? AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return err

	}
	return nil
}
