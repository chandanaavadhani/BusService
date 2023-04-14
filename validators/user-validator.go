package validators

import (
	"errors"
	"net/http"

	models "github.com/chandanaavadhani/BusService/models"
	utils "github.com/chandanaavadhani/BusService/utils"
)

func SignupValidation(user models.Signup) (int, error) {
	if user.Firstname == "" {
		return http.StatusBadRequest, errors.New("firstname missing")
	}
	if user.Lastname == "" {
		return http.StatusBadRequest, errors.New("lastname missing")
	}
	if user.Email == "" {
		return http.StatusBadRequest, errors.New("email missing")
	}
	if !utils.IsEmailValid(user.Email) {
		return http.StatusBadRequest, errors.New("email format is not valid")
	}
	if user.Password == "" {
		return http.StatusBadRequest, errors.New("password missing")
	}
	if !utils.IsPasswordValid(user.Password) {
		return http.StatusBadRequest, errors.New("password length is not valid")
	}
	if user.Phone == "" {
		return http.StatusBadRequest, errors.New("phone missing")
	}
	if !utils.IsPhoneNumberValid(user.Phone) {
		return http.StatusBadRequest, errors.New("phone number formate is required")
	}
	return 200, nil
}
func LoginValidation(user models.Login) (int, error) {
	if user.Email == "" {
		return http.StatusBadRequest, errors.New("email missing")
	}
	if !utils.IsEmailValid(user.Email) {
		return http.StatusBadRequest, errors.New("email format is not valid")
	}
	if user.Password == "" {
		return http.StatusBadRequest, errors.New("password missing")
	}
	if !utils.IsPasswordValid(user.Password) {
		return http.StatusBadRequest, errors.New("password length is not valid")
	}
	return 200, nil
}
func UpdatePasswordValidation(user models.PasswordUpdate) (int, error) {
	if user.Email == "" {
		return http.StatusBadRequest, errors.New("email missing")
	}
	if !utils.IsEmailValid(user.Email) {
		return http.StatusBadRequest, errors.New("email format is not valid")
	}
	if user.OldPassword == "" {
		return http.StatusBadRequest, errors.New("oldpassword missing")
	}
	if !utils.IsPasswordValid(user.OldPassword) {
		return http.StatusBadRequest, errors.New("password length is not valid")
	}
	if user.NewPassword == "" {
		return http.StatusBadRequest, errors.New("newpassword missing")
	}
	if !utils.IsPasswordValid(user.NewPassword) {
		return http.StatusBadRequest, errors.New("newpassword length is not valid")
	}
	return 200, nil
}
