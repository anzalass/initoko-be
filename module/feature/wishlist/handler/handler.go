package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/wishlist"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type WishlistHandler struct {
	service wishlist.ServiceWishlistInterface
}

func NewWishlistHandler(service wishlist.ServiceWishlistInterface) wishlist.HandlerWishlistInterface {
	return &WishlistHandler{service: service}
}

func (h *WishlistHandler) AddWishlist() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(entities.WishlistModels)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}
		value := &entities.WishlistModels{
			IDUser:    req.IDUser,
			IDProduct: req.IDProduct,
			Email:     req.Email,
			Name:      req.Name,
			Harga:     req.Harga,
			Foto:      req.Foto,
		}

		res, err := h.service.AddWishlist(value)
		if err != nil {

			return c.JSON(http.StatusBadRequest, map[string]any{
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
func (h *WishlistHandler) GetWishlistByIdUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse := c.Param("id")
		res, err := h.service.GetWishlistByIdUser(idparse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
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
func (h *WishlistHandler) DeleteWishListById() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, _ := strconv.ParseUint(c.Param("id"), 10, 64)
		err := h.service.DeleteWishListById(idparse)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, map[string]any{
			"message": "success",
		})
	}
}
