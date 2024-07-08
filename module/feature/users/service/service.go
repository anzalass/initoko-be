package service

import (
	"errors"
	"initoko/module/entities"
	"initoko/module/feature/users"
	"initoko/module/feature/users/dto"
	"initoko/utils"
	"time"
)

type UsersService struct {
	repo users.RepositoryUserInterface
}

func NewUsersService(repo users.RepositoryUserInterface) users.ServiceUserInterface {
	return &UsersService{
		repo: repo,
	}
}

func (s *UsersService) RegisterUser(newData *entities.UsersModels) (*entities.UsersModels, error) {

	value := &entities.UsersModels{
		ID:       newData.ID,
		Name:     newData.Name,
		Email:    newData.Email,
		Password: newData.Password,
		AkunTipe: "biasa",
		Status:   "nonaktif",
		Role:     "member",
		Avatar:   "https://images.tokopedia.net/img/cache/100-square/tPxBYm/2023/1/20/785ac6cb-d67b-42bd-97f8-6a06b9269130.jpg",
	}

	res, err := s.repo.RegisterUser(value)
	if err != nil {
		return nil, err
	}

	otp := utils.GenerateOTP(6)
	errsendotp := utils.EmailService(res.Email, otp)
	if errsendotp != nil {
		return nil, errsendotp
	}
	newOtp := &entities.OtpModels{
		Email:   res.Email,
		Kode:    otp,
		Expired: uint64(time.Now().Add(7 * time.Minute).Unix()),
	}
	errsaveotp := s.repo.CreateOTP(newOtp)
	if errsaveotp != nil {
		return nil, errsaveotp
	}
	return res, nil
}
func (s *UsersService) RegisterUserGoogle(newData *entities.UsersModels) (*entities.UsersModels, error) {

	value := &entities.UsersModels{
		ID:       newData.ID,
		Name:     newData.Name,
		Email:    newData.Email,
		Password: newData.Password,
		AkunTipe: "google",
		Status:   "aktif",
		Role:     "member",
		Avatar:   newData.Avatar,
	}

	res, err := s.repo.RegisterUser(value)
	if err != nil {
		return nil, err
	}

	return res, nil
}
func (s *UsersService) LoginUser(email, password string) (any, error) {
	res, err := s.repo.LoginUser(email, password)
	if err != nil {
		return nil, err
	}

	token, err2 := utils.GenerateJWT(res.ID, res.Email, res.Role)
	if err2 != nil {
		return nil, err2
	}

	result := &dto.LoginResponse{
		ID:       res.ID,
		Name:     res.Name,
		Email:    res.Email,
		Role:     res.Role,
		Avatar:   res.Avatar,
		TipeAkun: res.AkunTipe,
		Token:    token,
	}

	return result, nil
}
func (s *UsersService) UpdateUser(id string, newData *entities.UsersModels) (*entities.UsersModels, error) {
	value := &entities.UsersModels{
		Name:     newData.Name,
		Email:    newData.Email,
		Password: newData.Password,
		Role:     "member",
		Avatar:   "https://images.tokopedia.net/img/cache/100-square/tPxBYm/2023/1/20/785ac6cb-d67b-42bd-97f8-6a06b9269130.jpg",
	}

	res, err := s.repo.UpdateUser(id, value)
	if err != nil {
		return nil, err
	}
	return res, nil
}
func (s *UsersService) DeleteUser(id uint64) error {
	err := s.repo.DeleteUser(id)
	if err != nil {
		return err
	}

	return nil
}

func (s *UsersService) UpdateAvatar(id string, file interface{}, name string) (bool, error) {
	urlavatar, err := utils.ImageHandler(file, name, "avatar")
	if err != nil {
		return false, errors.New(err.Error())
	}

	value := &entities.UsersModels{
		Avatar: urlavatar,
	}

	_, err2 := s.repo.UpdateUser(id, value)
	if err2 != nil {
		return false, errors.New(err2.Error())
	}
	return true, nil
}

func (s *UsersService) AktifasiAkun(kode string, email string) (bool, error) {
	res := s.repo.CompareOTPByEmail(kode, email)
	if res == true {
		err := s.repo.ChangeStatusByEmail(email)
		if err != nil {
			return false, err
		}
		err2 := s.repo.DeleteOTPByEmail(email)
		if err2 != nil {
			return false, err2
		}
	} else if res == false {

		return false, nil
	}

	return true, nil
}

func (s *UsersService) SendOtp(email string) error {

	akun, err := s.repo.GetAkunByEmail(email)
	if akun == false {
		return err
	} else if akun == true {
		otp := utils.GenerateOTP(6)
		value := &entities.OtpModels{
			Kode:    otp,
			Email:   email,
			Expired: uint64(time.Now().Add(7 * time.Minute).Unix()),
		}
		res := s.repo.CreateOTP(value)
		if res != nil {
			return res
		}
		errsendotp := utils.EmailService(email, otp)
		if errsendotp != nil {
			return errsendotp
		}
	}

	return nil
}

func (s *UsersService) ProfilePage(email string) (*entities.UsersModels, error) {
	akun, err := s.repo.GetAkunByEmailProfile(email)
	if err != nil {
		return nil, err
	}

	return akun, nil
}

func (s *UsersService) AlamatPage(email string) ([]*entities.AlamatPenerimaModels, error) {

	alamat, err := s.repo.GetAlamatByEmailProfile(email)
	if err != nil {
		return nil, err
	}

	return alamat, nil
}
