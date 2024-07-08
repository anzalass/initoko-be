package handler

import (
	"initoko/module/entities"
	"initoko/module/feature/product"
	"initoko/module/feature/product/dto"
	"initoko/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type HandlerProduct struct {
	service product.ProductServiceInterface
}

func NewHandlerProduct(service product.ProductServiceInterface) product.ProductHandlerInterface {
	return &HandlerProduct{
		service: service,
	}
}

func (h *HandlerProduct) GetAllProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		res, err := h.service.GetAllProduct()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
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

func (h *HandlerProduct) CreateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.CreateProductRequest)
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
		value := &entities.ProductModels{
			Name:       req.Name,
			Harga:      req.Harga,
			Diskon:     req.Diskon,
			Kategori:   req.Kategori,
			Deskripsi:  req.Deskripsi,
			Tags:       req.Tags,
			Stok:       req.Stok,
			DibuatOleh: req.DibuatOleh,
		}

		res, err := h.service.CreateProduct(value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
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
func (h *HandlerProduct) UpdateProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.CreateProductRequest)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		value := &entities.ProductModels{
			Name:       req.Name,
			Harga:      req.Harga,
			Diskon:     req.Diskon,
			Kategori:   req.Kategori,
			Deskripsi:  req.Deskripsi,
			Tags:       req.Tags,
			Stok:       req.Stok,
			DibuatOleh: req.DibuatOleh,
		}
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res, err := h.service.UpdateProduct(idparse, value)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
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
func (h *HandlerProduct) DeleteProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		err2 := h.service.DeleteProduct(idparse)
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
func (h *HandlerProduct) GetProductByID() echo.HandlerFunc {
	return func(c echo.Context) error {
		idparse, err := strconv.ParseUint(c.Param("id"), 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "id tidak ditemukan",
			})
		}

		res, err := h.service.GetProductByID(idparse)
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
func (h *HandlerProduct) UploadFotoProduct() echo.HandlerFunc {
	return func(c echo.Context) error {
		// req := new(entities.FotoProductModels)
		var res any
		var err error
		form, _ := c.MultipartForm()
		idproduct := c.FormValue("idproduct")
		convidproduct, err := strconv.Atoi(idproduct)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Parameter idproduct tidak valid",
				"error":   err.Error(),
				"id":      convidproduct,
			})
		}
		files := form.File["files"]
		var uploadedPhotos []string
		for _, file := range files {
			src, _ := file.Open()
			defer src.Close()
			uploadUrl, _ := utils.ImageHandler(src, file.Filename, "product")
			value := &entities.FotoProductModels{
				Url: uploadUrl,
			}

			res, err = h.service.UploadFotoProduct(value, uint64(convidproduct))
			if err != nil {
				return c.JSON(http.StatusInternalServerError, map[string]any{
					"message": "gagal",
					"error":   err.Error(),
					"id":      convidproduct,
				})
			}
			uploadedPhotos = append(uploadedPhotos, uploadUrl)

		}

		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res,
			"foto":    uploadedPhotos,
			"id":      convidproduct,
		})

		// file, _ := c.FormFile("file")

		// src, _ := file.Open()
		// defer func(src multipart.File) {
		// 	_ = src.Close()
		// }(src)

		// uploadUrl, _ := utils.ImageHandler(src, file.Filename)

	}
}

// func (h *HandlerProduct) UploadFotoProduct() echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		form, err := c.MultipartForm()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 				"message": "gagal",
// 				"error":   err.Error(),
// 			})
// 		}

// 		files := form.File["files"]
// 		var uploadedPhotos []string

// 		for _, file := range files {
// 			src, err := file.Open()
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 					"message": "gagal",
// 					"error":   err.Error(),
// 				})
// 			}
// 			defer src.Close()

// 			uploadUrl, err := utils.ImageHandler(src, file.Filename)
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 					"message": "gagal",
// 					"error":   err.Error(),
// 				})
// 			}

// 			req := &entities.FotoProductModels{
// 				IDProduct: h.IDProduct, // Anda harus mengatur IDProduct sesuai dengan kebutuhan Anda
// 				Url:       uploadUrl,
// 			}

// 			_, err = h.service.UploadFotoProduct(req)
// 			if err != nil {
// 				return c.JSON(http.StatusInternalServerError, map[string]interface{}{
// 					"message": "gagal",
// 					"error":   err.Error(),
// 				})
// 			}

// 			uploadedPhotos = append(uploadedPhotos, uploadUrl)
// 		}

// 		return c.JSON(http.StatusOK, map[string]interface{}{
// 			"message": "Foto berhasil diunggah",
// 			"data":    uploadedPhotos,
// 		})
// 	}
// }

func (h *HandlerProduct) GetAllProductByFilter() echo.HandlerFunc {
	return func(c echo.Context) error {
		max, _ := strconv.ParseUint(c.QueryParam("max"), 10, 64)
		min, _ := strconv.ParseUint(c.QueryParam("min"), 10, 64)
		filter := c.QueryParam("filter")
		cari := c.QueryParam("cari")
		kategori := c.QueryParam("kategori")

		res, err := h.service.GetAllProductByFilter(min, max, cari, kategori, filter)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
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

func (h *HandlerProduct) Homepage() echo.HandlerFunc {
	return func(c echo.Context) error {
		res1, res2, res3, res4 := h.service.Homepage()
		return c.JSON(http.StatusOK, map[string]any{
			"message": "berhasil",
			"data":    res1,
			"data2":   res2,
			"data3":   res3,
			"data4":   res4,
		})
	}
}

func (h *HandlerProduct) GetProductByKategori() echo.HandlerFunc {
	return func(c echo.Context) error {
		req := new(dto.RelatedKategori)
		if err := c.Bind(req); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]any{
				"message": "gagal",
				"error":   err.Error(),
			})
		}

		res, err := h.service.GetProductByKategori(c.Param("kategori"))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]any{
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
