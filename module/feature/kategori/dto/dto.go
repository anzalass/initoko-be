package dto

type CreateKategoriRequest struct {
	Name string `json:"name" form:"name"`
	Tipe string `json:"tipe" form:"tipe"`
	Foto string `json:"foto" form:"foto"`
}
