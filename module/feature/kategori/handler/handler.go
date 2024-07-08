package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/kategori"
	"initoko/module/feature/kategori/dto"
	"initoko/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type KategoriHandler struct {
	service kategori.KategoriServiceInterface
}

func NewKategoriHandler(service kategori.KategoriServiceInterface) kategori.KategoriHandlerInterface {
	return &KategoriHandler{service: service}
}

func (h *KategoriHandler) CreateKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.CreateKategoriRequest)
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

		file, _ := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.KategoriModels{
			Name: req.Name,
			Tipe: req.Tipe,
		}

		res, err := h.service.CreateKategori(value, src, file.Filename)
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
func (h *KategoriHandler) UpdateKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.CreateKategoriRequest)
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

		file, _ := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.KategoriModels{
			Name: req.Name,
			Tipe: req.Tipe,
		}

		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.UpdateKategori(idparse, value, src, file.Filename)
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
func (h *KategoriHandler) DeleteKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		err2 := h.service.DeleteKategori(idparse)
		if err2 != nil {

			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})

		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
		})
	}
}
func (h *KategoriHandler) GetKategoriByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		res, err := h.service.GetKategoriByID(idparse)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
			"data":    res,
		})

	}
}
func (h *KategoriHandler) GetAllKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllKategori()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
			"data":    res,
		})

	}
}
func (h *KategoriHandler) GetAllKategoriName() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllKategoriName()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "success",
			"data":    res,
		})

	}
}
