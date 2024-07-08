package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/jumbotron"
	"initoko/module/feature/jumbotron/dto"
	"initoko/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type JumbotronHandler struct {
	service jumbotron.JumbotronServiceInterface
}

func NewJumbotronHandler(service jumbotron.JumbotronServiceInterface) jumbotron.JumbotronHandlerInterface {
	return &JumbotronHandler{service: service}
}

func (h *JumbotronHandler) CreateJumbotron() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.JumbotronRequest)
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

		file, errfoto := c.FormFile("file")
		if errfoto != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   errfoto.Error(),
			})
		}
		src, erropen := file.Open()
		if erropen != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   erropen.Error(),
			})
		}
		value := &entities.JumbotronModels{
			Name:      req.Name,
			Deskripsi: req.Deskripsi,
			Foto:      req.Foto,
		}

		res, err := h.service.CreateJumbotron(value, src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
func (h *JumbotronHandler) UpdateJumbotron() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.JumbotronRequest)
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
		src, _ := file.Open()

		value := &entities.JumbotronModels{
			Name:      req.Name,
			Deskripsi: req.Deskripsi,
			Foto:      req.Foto,
		}
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		res, err := h.service.UpdateJumbotron(idparse, value, src, file.Filename)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
			"data":    res,
		})
	}
}
func (h *JumbotronHandler) DeleteJumbotron() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		err2 := h.service.DeleteJumbotron(idparse)
		if err2 != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"message": "id tidak ditemukan",
				"error":   err2.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
		})
	}
}
func (h *JumbotronHandler) GetJumbotronById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res, err := h.service.GetJumbotronById(idparse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})

	}
}
func (h *JumbotronHandler) GetAllJumbotron() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllJumbotron()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})

	}
}
