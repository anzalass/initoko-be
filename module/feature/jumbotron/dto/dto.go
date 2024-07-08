package dto

type JumbotronRequest struct {
	Name      string `json:"name" form:"name"`
	Deskripsi string `json:"deskripsi" form:"deskripsi"`
	Status    string `json:"status" form:"status"`
	Foto      string `json:"foto" form:"foto"`
}
