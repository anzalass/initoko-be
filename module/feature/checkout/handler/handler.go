package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/checkout"
	"initoko/module/feature/checkout/dto"
	"initoko/utils/midtrans"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CheckoutHandler struct {
	service  checkout.ServiceCheckout
	midtrans midtrans.MidtransServiceInterface
}

func NewCheckoutHandler(service checkout.ServiceCheckout, midtrans midtrans.MidtransServiceInterface) checkout.HandelerCheckout {
	return &CheckoutHandler{service: service, midtrans: midtrans}
}

func (h *CheckoutHandler) CreateTransaksi() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.DtoTransasksi)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		//hargaongkir, _ := strconv.Atoi(req.Ongkir)

		var total = req.Harga + req.Ongkir

		valuetransaksi := &entities.TransaksiModels{
			IDUser:   req.IDUser,
			Email:    req.Email,
			NamaUser: req.NamaUser,
			Alamat:   req.Alamat,
			Harga:    uint64(total),
		}

		transaksi, err := h.service.CreateTransaksi(valuetransaksi)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		var productt *entities.ProductPesananModels

		for _, prdct := range req.Product {
			value2 := &entities.ProductPesananModels{
				IDUser:      transaksi.IDUser,
				IDProduct:   prdct.IDProduct,
				IDTransaksi: transaksi.ID,
				Email:       transaksi.Email,
				Namaproduct: prdct.NamaProduct,
				Quantity:    prdct.Quantity,
				Harga:       prdct.Harga,
				Foto:        prdct.Foto,
			}
			productt, err = h.service.CreateProductTransaksi(value2)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "gagal",
					"error":   err.Error(),
				})
			}
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message":       "berhasil",
			"datatransaksi": transaksi,
			"ID":            transaksi.ID,
			"dataproduct":   productt,
		})

	}
}
func (h *CheckoutHandler) ChangeStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.DtoChangeStatusTransaksi)
		id := c.Param("id")
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		err := h.service.ChangeStatus(id, req.Status)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "sukses",
		})
	}

}
func (h *CheckoutHandler) ChangeStatusPembayaran() echo.HandlerFunc {
	return func(c echo.Context) error {
		var notificationPayload map[string]interface{}
		if err := c.Bind(&notificationPayload); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "gagal mendapatkan notification payload",
			})
		}

		orderId, exists := notificationPayload["order_id"].(string)
		if !exists {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal mendapatkan order id",
			})
		}

		success, err := h.midtrans.VerifyPayment(c.Request().Context(), orderId)
		if err != nil {
			return err
		}
		if success == "success" {
			err := h.service.ChangeStatusPembayaran(orderId)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"error":   err,
					"message": "gagal mengkonfirmasi pembayaran",
				})
			}
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"message": "berhasil mengkonfirmasi pembayaran",
		})
	}

}

func (h *CheckoutHandler) AllPesananByEmail() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		res, err := h.service.AllPesananByEmail(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}
}
func (h *CheckoutHandler) DetailPesananById() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		res, err := h.service.DetailPesananById(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}
}

func (h *CheckoutHandler) AllPesanan() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.AllPesanan()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}

}
func (h *CheckoutHandler) InputResi() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		req := new(dto.DtoChangeStatusTransaksi)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.TransaksiModels{
			Resi: req.Resi,
		}

		err := h.service.InputResi(id, value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
		})
	}

}
