package dto

type LoginResponse struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Avatar   string `json:"avatar"`
	Token    string `json:"token"`
	TipeAkun string `json:"tipe_akun"`
}
type LoginRequest struct {
	Email    string `json:"email"  validate:"required"`
	Password string `json:"password" form:"password"  validate:"required"`
}

type RegisterRequestGoogle struct {
	ID       string `json:"id" validate:"required"`
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Avatar   string `json:"avatar" form:"avatar"`
	Password string `json:"password" form:"password"`
}
type RegisterRequest struct {
	Name     string `json:"name" form:"name" validate:"required"`
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
}

type AktivasiAkunan struct {
	Email string `json:"email" form:"email" validate:"required"`
}
