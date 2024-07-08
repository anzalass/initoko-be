package dto

type CreateProductRequest struct {
	Name       string `json:"name" form:"name" validate:"required"`
	Harga      uint64 `json:"harga" form:"harga" validate:"required"`
	Diskon     uint64 `json:"diskon" form:"diskon" validate:"required"`
	Kategori   string `json:"kategori" form:"kategori" validate:"required"`
	Deskripsi  string `json:"deskripsi" form:"deskripsi" validate:"required"`
	Tags       string `json:"tags" form:"tags" validate:"required"`
	Stok       uint64 `json:"stok" form:"stok" validate:"required"`
	DibuatOleh string `json:"dibuat_oleh" form:"dibuat_oleh" validate:"required"`
}

type RelatedKategori struct {
	Kategori string `json:"kategori" form:"kategori" validate:"required"`
}
