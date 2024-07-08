package entities

import "time"

type UsersModels struct {
	ID             string                 `gorm:"column:id;type:VARCHAR(255);primaryKey" json:"id"`
	Name           string                 `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Avatar         string                 `gorm:"column:avatar;type:TEXT" json:"avatar"`
	Email          string                 `gorm:"column:email;type:VARCHAR(255);unique" json:"email"`
	Password       string                 `gorm:"column:password;type:VARCHAR(255)" json:"password"`
	Role           string                 `gorm:"column:role;type:VARCHAR(10)" json:"role"`
	Status         string                 `gorm:"column:status;type:VARCHAR(20)" json:"status"`
	AkunTipe       string                 `gorm:"column:akun_tipe;type:VARCHAR(10)" json:"akun_tipe"`
	Review         []ReviewModels         `gorm:"foreignKey:IDUser" json:"review"`
	Transaksi      []TransaksiModels      `gorm:"foreignKey:IDUser" json:"transaksi"`
	ProductPesanan []ProductPesananModels `gorm:"foreignKey:IDUser" json:"product_pesanan"`
	Wishlist       []WishlistModels       `gorm:"foreignKey:IDUser" json:"wishlist"`
	AlamatPenerima []AlamatPenerimaModels `gorm:"foreignKey:IDUser" json:"alamat_penerima"`
	CreatedAt      time.Time              `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt      time.Time              `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt      *time.Time             `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type AlamatPenerimaModels struct {
	ID            uint64     `gorm:"column:id;primaryKey" json:"id"`
	IDUser        string     `gorm:"column:id_user;type:VARCHAR(255)" json:"id_user"`
	Email         string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Name          string     `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	NoWhatsapp    string     `gorm:"column:no_whatsapp;type:VARCHAR(13)" json:"no_whatsapp"`
	Desa          string     `gorm:"column:desa;type:VARCHAR(255)" json:"desa"`
	Kecamatan     string     `gorm:"column:kecamatan;type:VARCHAR(255)" json:"kecamatan"`
	Kabupaten     string     `gorm:"column:kabupaten;type:VARCHAR(255)" json:"kabupaten"`
	Provinsi      string     `gorm:"column:Provinsi;type:VARCHAR(255)" json:"Provinsi"`
	AlamatLengkap string     `gorm:"column:alamat_lengkap;type:TEXT" json:"alamat_lengkap"`
	CreatedAt     time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type ProductModels struct {
	ID          uint64              `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	Name        string              `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Harga       uint64              `gorm:"column:harga;type:BIGINT" json:"harga"`
	Diskon      uint64              `gorm:"column:diskon;type:INT" json:"diskon"`
	Kategori    string              `gorm:"column:kategori;type:VARCHAR(255)" json:"kategori"`
	Deskripsi   string              `gorm:"column:deskripsi;type:TEXT" json:"deskripsi"`
	Tags        string              `gorm:"column:tags;type:TEXT" json:"tags"`
	Ratings     uint64              `gorm:"column:ratings;type:INT" json:"ratings"`
	Stok        uint64              `gorm:"column:stok;type:INT" json:"stok"`
	Terjual     uint64              `gorm:"column:terjual;type:BIGINT" json:"terjual"`
	DibuatOleh  string              `gorm:"column:dibuat_oleh;type:VARCHAR(255)" json:"dibuat_oleh"`
	FotoProduct []FotoProductModels `gorm:"foreignKey:IDProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"foto"`
	Review      []ReviewModels      `gorm:"foreignKey:IDProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"review"`
	Wishlist    []WishlistModels    `gorm:"foreignKey:IDProduct;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"wishlist"`
	CreatedAt   time.Time           `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time           `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt   *time.Time          `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type FotoProductModels struct {
	ID        uint64        `gorm:"column:id;type:BIGINT ;primaryKey" json:"id" `
	IDProduct uint64        `gorm:"column:id_product;type:BIGINT" json:"id_product" form:"id_product"`
	Product   ProductModels `gorm:"foreignKey:IDProduct;references:ID"`
	Url       string        `gorm:"column:url;type:TEXT" json:"url"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time    `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type ReviewModels struct {
	ID        uint64        `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	IDUser    string        `gorm:"column:id_user;type:VARCHAR(255) " json:"id_user"`
	Email     string        `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	IDProduct uint64        `gorm:"column:id_product;type:BIGINT" json:"id_product"`
	Product   ProductModels `gorm:"foreignKey:IDProduct;references:ID"`
	Ratings   uint64        `gorm:"column:ratings;type:INT " json:"ratings"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time    `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type TransaksiModels struct {
	ID               string                 `gorm:"column:id;type:VARCHAR(300) ;primaryKey" json:"id"`
	IDUser           string                 `gorm:"column:id_user;type:VARCHAR(255) " json:"id_user"`
	Product          []ProductPesananModels `gorm:"foreignKey:IDTransaksi;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"product"`
	Email            string                 `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	UrlPembayaran    string                 `gorm:"column:url_pembayaran;type:TEXT" json:"url_pembayaran"`
	NamaUser         string                 `gorm:"column:nama_user;type:VARCHAR(255)" json:"nama_user"`
	Status           string                 `gorm:"column:status;type:VARCHAR(50)" json:"status"`
	Alamat           string                 `gorm:"column:alamat;type:TEXT" json:"alamat"`
	Harga            uint64                 `gorm:"column:harga;type:BIGINT " json:"harga"`
	Resi             string                 `gorm:"column:resi;type:VARCHAR(255)" json:"resi"`
	Pesan            string                 `gorm:"column:pesan;type:TEXT" json:"pesan"`
	StatusPembayaran string                 `gorm:"column:status_pembayaran;type:VARCHAR(50)" json:"status_pembayaran"`
	CreatedAt        time.Time              `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time              `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time             `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type ProductPesananModels struct {
	ID               uint64     `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	IDUser           string     `gorm:"column:id_user;type:VARCHAR(255)" json:"id_user"`
	IDProduct        uint64     `gorm:"column:id_product;type:BIGINT" json:"id_product"`
	IDTransaksi      string     `gorm:"column:id_transaksi;type:VARCHAR(300) " json:"id_transaksi"`
	Email            string     `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Namaproduct      string     `gorm:"column:nama_product;type:VARCHAR(255)" json:"nama_product"`
	Quantity         uint64     `gorm:"column:quantity;type:BIGINT " json:"quantity"`
	Foto             string     `gorm:"column:foto;type:TEXT" json:"foto"`
	Harga            uint64     `gorm:"column:harga;type:BIGINT " json:"harga"`
	StatusPembayaran string     `gorm:"column:status_pembayaran;type:VARCHAR(50)" json:"status_pembayaran"`
	StatusPengiriman string     `gorm:"column:status_pengiriman;type:VARCHAR(50)" json:"status_pengiriman"`
	CreatedAt        time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt        time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt        *time.Time `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type WishlistModels struct {
	ID        uint64        `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	IDUser    string        `gorm:"column:id_user;type:VARCHAR(255)" json:"id_user"`
	Email     string        `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Name      string        `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Harga     uint64        `gorm:"column:harga;type:INTEGER" json:"harga"`
	IDProduct uint64        `gorm:"column:id_product;type:BIGINT;not null" json:"id_product"`
	Foto      string        `gorm:"column:foto;type:TEXT" json:"foto"`
	Product   ProductModels `gorm:"foreignKey:IDProduct;references:ID"`
	CreatedAt time.Time     `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time     `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time    `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type KategoriModels struct {
	ID        uint64     `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	Name      string     `gorm:"column:name;type:VARCHAR(255);unique" json:"name"`
	Tipe      string     `gorm:"column:tipe;type:VARCHAR(10);" json:"tipe"`
	Foto      string     `gorm:"column:foto;type:TEXT" json:"foto"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:TIMESTAMP;index" json:"deleted_at"`
}

type JumbotronModels struct {
	ID        uint64 `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	Name      string `gorm:"column:name;type:VARCHAR(255)" json:"name"`
	Deskripsi string `gorm:"column:deskripsi;type:TEXT" json:"deskripsi"`
	Status    string `gorm:"column:status;type:VARCHAR(20)" json:"status"`
	Foto      string `gorm:"column:foto;type:TEXT" json:"foto"`
}

type OtpModels struct {
	ID      uint64 `gorm:"column:id;type:BIGINT ;primaryKey" json:"id"`
	Kode    string `gorm:"column:kode;type:VARCHAR(10)" json:"otp"`
	Email   string `gorm:"column:email;type:VARCHAR(255)" json:"email"`
	Expired uint64 `gorm:"column:expired;type:INTEGER" json:"expired"`
}

func (UsersModels) TableName() string {
	return "users"
}
func (AlamatPenerimaModels) TableName() string {
	return "alamats"
}
func (ProductModels) TableName() string {
	return "products"
}
func (TransaksiModels) TableName() string {
	return "transaksis"
}
func (WishlistModels) TableName() string {
	return "wishlists"
}
func (ProductPesananModels) TableName() string {
	return "productpesanans"
}
func (ReviewModels) TableName() string {
	return "reviews"
}
func (FotoProductModels) TableName() string {
	return "fotoproducts"
}
func (KategoriModels) TableName() string {
	return "kategoris"
}
func (JumbotronModels) TableName() string {
	return "jumbotrons"
}
func (OtpModels) TableName() string {
	return "otps"
}
