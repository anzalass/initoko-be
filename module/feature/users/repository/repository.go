package repository

import (
	"initoko/module/entities"
	"initoko/module/feature/users"
	"time"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) users.RepositoryUserInterface {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) RegisterUser(newData *entities.UsersModels) (*entities.UsersModels, error) {
	if err := r.db.Create(&newData).Error; err != nil {
		return nil, err
	}
	return newData, nil
}
func (r *UserRepository) LoginUser(email, password string) (*entities.UsersModels, error) {
	users := entities.UsersModels{}

	if err := r.db.Where("email = ? AND password = ? AND status = ?", email, password, "aktif").First(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserRepository) UpdateUser(id string, newData *entities.UsersModels) (*entities.UsersModels, error) {
	user := &entities.UsersModels{}
	if err := r.db.Model(user).Where("id = ? AND deleted_at IS NULL", id).Updates(&newData).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) DeleteUser(id uint64) error {
	users := entities.UsersModels{}

	if err := r.db.Model(&users).Where("id = ? AND deleted_at IS NULL", id).Update("deleted_at", time.Now()).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) CreateOTP(newData *entities.OtpModels) error {
	if err := r.db.Create(&newData).Error; err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) DeleteOTPByEmail(email string) error {
	otp := &entities.OtpModels{}
	if err := r.db.Where("email = ?", email).Delete(otp).Error; err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) ChangeStatusByEmail(email string) error {
	users := &entities.UsersModels{}

	if err := r.db.Model(users).Where("email = ?", email).Update("status", "aktif").Error; err != nil {
		return err
	}
	return nil
}
func (r *UserRepository) CompareOTPByEmail(kode string, email string) bool {

	otp := &entities.OtpModels{}

	if err := r.db.Where("kode = ?", kode).First(otp).Error; err != nil {
		return false
	}

	if email == otp.Email && kode == otp.Kode && otp.Expired > uint64(time.Now().Unix()) {
		return true
	}

	return false
}

func (r *UserRepository) GetAkunByEmail(email string) (bool, error) {
	akun := &entities.UsersModels{}

	if err := r.db.Where("email = ? AND status = ?", email, "nonaktif").First(akun).Error; err != nil {
		return false, err
	}

	return true, nil
}

func (r *UserRepository) GetAkunByEmailProfile(email string) (*entities.UsersModels, error) {
	akun := &entities.UsersModels{}
	if err := r.db.Where("email = ? AND status= ?", email, "aktif").Select("name, email, avatar").First(akun).Error; err != nil {
		return nil, err
	}
	return akun, nil
}
func (r *UserRepository) GetAlamatByEmailProfile(email string) ([]*entities.AlamatPenerimaModels, error) {
	var alamat []*entities.AlamatPenerimaModels
	if err := r.db.Where("email = ?", email).Find(&alamat).Error; err != nil {
		return nil, err
	}

	return alamat, nil
}
