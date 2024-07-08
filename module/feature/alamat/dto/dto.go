package dto

type AlamatRequest struct {
	IDUser        string `json:"id_user" form:"id_user" validate:"required"`
	Name          string `json:"name" form:"name" validate:"required"`
	Email         string `json:"email" form:"email" validate:"required"`
	NoWhatsapp    string `json:"no_whatsapp" form:"no_whatsapp" validate:"required"`
	Desa          string `json:"desa" form:"desa" validate:"required"`
	Kecamatan     string `json:"kecamatan" form:"kecamatan" validate:"required"`
	Kabupaten     string `json:"kabupaten" form:"kabupaten" validate:"required"`
	Provinsi      string `json:"provinsi" form:"provinsi" validate:"required"`
	AlamatLengkap string `json:"alamat_lengkap" form:"alamaat_lengkap" validate:"required"`
}
