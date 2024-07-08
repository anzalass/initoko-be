package service

import (
	"context"
	"initoko/module/entities"
	"initoko/module/feature/checkout"
	"initoko/utils/midtrans"
)

type CheckoutService struct {
	repo     checkout.RepositoryCheckout
	midtrans midtrans.MidtransServiceInterface
}

func NewCheckoutService(repo checkout.RepositoryCheckout, midtrans midtrans.MidtransServiceInterface) checkout.ServiceCheckout {
	return &CheckoutService{
		repo:     repo,
		midtrans: midtrans,
	}
}

func (s *CheckoutService) CreateTransaksi(newData *entities.TransaksiModels) (*entities.TransaksiModels, error) {
	var ctx context.Context
	urlmidtrans, orderid, err := s.midtrans.GenerateSnapURL(ctx, int64(newData.Harga))
	if err != nil {
		return nil, err
	}

	value := &entities.TransaksiModels{
		ID:               orderid,
		IDUser:           newData.IDUser,
		Email:            newData.Email,
		UrlPembayaran:    urlmidtrans,
		NamaUser:         newData.NamaUser,
		Status:           "diproses",
		Alamat:           newData.Alamat,
		Harga:            newData.Harga,
		StatusPembayaran: "belum bayar",
	}

	res, err := s.repo.CreateTransaksi(value)
	if err != nil {
		return nil, err
	}

	return res, nil

}
func (s *CheckoutService) CreateProductTransaksi(newData *entities.ProductPesananModels) (*entities.ProductPesananModels, error) {
	value := &entities.ProductPesananModels{
		IDUser:      newData.IDUser,
		IDProduct:   newData.IDProduct,
		IDTransaksi: newData.IDTransaksi,
		Email:       newData.Email,
		Namaproduct: newData.Namaproduct,
		Quantity:    newData.Quantity,
		Foto:        newData.Foto,
		Harga:       newData.Harga,
	}

	res, err := s.repo.CreateProductTransaksi(value)
	if err != nil {
		return nil, err
	}

	return res, err
}
func (s *CheckoutService) ChangeStatus(idtransaksi string, status string) error {
	err := s.repo.ChangeStatus(idtransaksi, status)
	if err != nil {
		return err
	}

	return nil
}
func (s *CheckoutService) ChangeStatusPembayaran(idtransaksi string) error {
	err := s.repo.ChangeStatusPembayaran(idtransaksi)
	if err != nil {
		return err
	}

	return nil
}

func (s *CheckoutService) AllPesananByEmail(email string) ([]*entities.TransaksiModels, error) {
	res, err := s.repo.AllPesananByEmail(email)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *CheckoutService) DetailPesananById(id string) (*entities.TransaksiModels, error) {
	res, err := s.repo.DetailPesananById(id)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *CheckoutService) AllPesanan() ([]*entities.TransaksiModels, error) {
	res, err := s.repo.AllPesanan()
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *CheckoutService) InputResi(id string, newData *entities.TransaksiModels) error {
	value := &entities.TransaksiModels{
		Resi:   newData.Resi,
		Status: "sedang dikirim",
	}

	err := s.repo.InputResi(id, value)
	if err != nil {
		return err
	}
	return nil
}
