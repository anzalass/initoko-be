package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/alamat"
	"initoko/module/feature/alamat/dto"
	"initoko/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type AlamatHandler struct {
	service alamat.ServiceAlamatInterface
}

func NewAlamatHandler(service alamat.ServiceAlamatInterface) alamat.HandlerALamatInterface {
	return &AlamatHandler{service: service}
}

func (h *AlamatHandler) CreateAlamat() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.AlamatRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		if err := utils.ValidateStruct(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.AlamatPenerimaModels{
			IDUser:        req.IDUser,
			Name:          req.Name,
			Email:         req.Email,
			NoWhatsapp:    req.NoWhatsapp,
			Desa:          req.Desa,
			Kecamatan:     req.Kecamatan,
			Kabupaten:     req.Kabupaten,
			Provinsi:      req.Provinsi,
			AlamatLengkap: req.AlamatLengkap,
		}

		res, err := h.service.CreateAlamat(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sucess",
			"data":    res,
		})
	}
}
func (h *AlamatHandler) UpdateAlamat() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.AlamatRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		if err := utils.ValidateStruct(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.AlamatPenerimaModels{
			Name:          req.Name,
			NoWhatsapp:    req.NoWhatsapp,
			Desa:          req.Desa,
			Kecamatan:     req.Kecamatan,
			Kabupaten:     req.Kabupaten,
			Provinsi:      req.Provinsi,
			AlamatLengkap: req.AlamatLengkap,
		}
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.UpdateAlamat(idparse, value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "sucess",
			"data":    res,
		})
	}
}
func (h *AlamatHandler) DeleteAlamatById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		err := h.service.DeleteAlamatById(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "sucess",
		})
	}
}
func (h *AlamatHandler) GetAlamatById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.GetAlamatById(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "sucess",
			"data":    res,
		})
	}
}
