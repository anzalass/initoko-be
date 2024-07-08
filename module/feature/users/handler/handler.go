package handler

import (
	"fmt"
	"initoko/module/entities"
	"initoko/module/feature/users"
	"initoko/module/feature/users/dto"
	"initoko/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	service users.ServiceUserInterface
}

func NewUsersHandler(service users.ServiceUserInterface) users.HandlerUserInterface {
	return &UsersHandler{
		service: service,
	}
}

func (h *UsersHandler) RegisterUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		RegisterRequest := new(dto.RegisterRequest)
		if err := c.Bind(RegisterRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		if err := utils.ValidateStruct(RegisterRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		makeid := time.Now().Unix()
		id := makeid * 10
		idconv := strconv.Itoa(int(id))

		newData := &entities.UsersModels{
			ID:       idconv,
			Name:     RegisterRequest.Name,
			Email:    RegisterRequest.Email,
			Password: RegisterRequest.Password,
		}

		res, err := h.service.RegisterUser(newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "Gagal",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
			"message": "Berhasil",
		})

	}
}

func (h *UsersHandler) RegisterUserGoogle() echo.HandlerFunc {
	return func(c echo.Context) error {
		RegisterRequest := new(dto.RegisterRequestGoogle)
		if err := c.Bind(RegisterRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		if err := utils.ValidateStruct(RegisterRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		newData := &entities.UsersModels{
			ID:       RegisterRequest.ID,
			Name:     RegisterRequest.Name,
			Email:    RegisterRequest.Email,
			Password: RegisterRequest.Password,
			Avatar:   RegisterRequest.Avatar,
		}

		res, err := h.service.RegisterUserGoogle(newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "Gagal",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"success": true,
			"data":    res,
			"message": "Berhasil",
		})

	}
}
func (h *UsersHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		LoginRequest := new(dto.LoginRequest)
		if err := c.Bind(LoginRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		if err := utils.ValidateStruct(LoginRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		newData := &entities.UsersModels{
			Email:    LoginRequest.Email,
			Password: LoginRequest.Password,
		}

		res, err := h.service.LoginUser(newData.Email, newData.Password)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "Gagal total",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data": res,

			"message": "Berhasil",
		})
	}
}
func (h *UsersHandler) UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		RegisterRequest := new(entities.UsersModels)
		if err := c.Bind(RegisterRequest); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		newData := &entities.UsersModels{
			Name:     RegisterRequest.Name,
			Password: RegisterRequest.Password,
		}

		id := c.Param("id")
		res, err := h.service.UpdateUser(id, newData)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "Gagal",
			})
		}

		return c.JSON(http.StatusOK, map[string]any{
			"data":    res,
			"message": "Berhasil",
		})
	}
}

func (h *UsersHandler) DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		err2 := h.service.DeleteUser(idparse)
		if err2 != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "berhasil",
		})

	}
}
func (h *UsersHandler) UpdateAvatar() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse := c.Param("id")

		file, _ := c.FormFile("file")
		src, err := file.Open()
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": fmt.Sprintf("invalid file, %s", err.Error()),
			})
		}

		res, err2 := h.service.UpdateAvatar(idparse, src, file.Filename)
		if err2 != nil {

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

func (h *UsersHandler) AktivasiAkun() echo.HandlerFunc {
	return func(c echo.Context) error {
		kode := c.Param("kode")
		email := c.Param("email")

		res, _ := h.service.AktifasiAkun(kode, email)
		if res == false {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   "error",
			})

		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
		})
	}
}

func (h *UsersHandler) SendOtp() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.AktivasiAkunan)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err,
			})
		}

		if err := utils.ValidateStruct(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"error":   err,
				"message": "Format input yang Anda masukkan tidak sesuai",
			})
		}

		res := h.service.SendOtp(req.Email)
		if res != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   res,
				"message": "error",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"message": "Berhasil Periksa Email" + " " + req.Email + " " + "Sekarang",
		})

	}
}

func (h *UsersHandler) ProfilePage() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		res, err := h.service.ProfilePage(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "error",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data":    res,
			"message": "success",
		})
	}
}
func (h *UsersHandler) AlamatPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		email := c.Param("email")
		res, err := h.service.AlamatPage(email)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
				"error":   err,
				"message": "error",
			})
		}
		return c.JSON(http.StatusOK, map[string]any{
			"data":    res,
			"message": "success",
		})
	}
}
