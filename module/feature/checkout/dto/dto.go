package dto

type DtoProduct struct {
	IDUser        string `json:"id_user"`
	IDProduct     uint64 `json:"id_product"`
	IDTrasansaksi string `json:"id_transaksi"`
	Email         string `json:"email"`
	NamaProduct   string `json:"nama_product"`
	Quantity      uint64 `json:"quantity"`
	Foto          string `json:"foto"`
	Harga         uint64 `json:"harga"`
}

type DtoTransasksi struct {
	IDUser string `json:"id_user"`
	// IDPembayaran  string       `json:"id_pembayaran"`
	Product []DtoProduct `json:"product"`
	Email   string       `json:"email"`
	//UrlPembayaran string       `json:"url_pembayaran"`
	NamaUser string `json:"nama_user"`
	//Status   string `json:"status"`
	Alamat string `json:"alamat"`
	Harga  uint64 `json:"harga"`
	Ongkir uint64 `json:"ongkir"`
}

type DtoChangeStatusTransaksi struct {
	Status string `json:"status"`
	Resi   string `json:"resi"`
}
