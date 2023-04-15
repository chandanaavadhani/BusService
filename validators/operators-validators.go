package validators

import (
	"database/sql"
	"errors"
	"net/http"
	"strings"

	models "github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
)

func SignupOperatorsValidation(db *sql.DB, newOperator models.Operator) (int, error) {

	if !utils.IsEmailValid(newOperator.Email) {
		return http.StatusBadRequest, errors.New("Email is not valid")
	}
	if strings.Split(newOperator.Email, "@")[1] != "operator.busservice.com" {
		return http.StatusBadRequest, errors.New("Not An Valid Operator Email")
	}
	if err := repository.IsOperatorExist(db, newOperator.Email); err != nil {
		return http.StatusForbidden, err
	}
	if !utils.IsPasswordValid(newOperator.Password) {
		return http.StatusBadRequest, errors.New("Password is not Valid")
	}
	if !utils.IsPhoneNumberValid(newOperator.Contact) {
		return http.StatusBadRequest, errors.New("Invalid phone number")
	}
	if newOperator.OperatorName == "" {
		return http.StatusBadRequest, errors.New("Name is required")
	}

	return http.StatusOK, nil
}

func UpdateOperatorsValidation(db *sql.DB, updateOperator models.UpdateOperator, TokenIsAdmin bool, TokenIsOperator bool, operatorId string) (int, error) {
	if !TokenIsAdmin && !TokenIsOperator {
		return http.StatusForbidden, errors.New("Donot have access to update the Operator")
	}
	if !utils.IsPhoneNumberValid(updateOperator.Contact) {
		return http.StatusBadRequest, errors.New("Invalid phone number")
	}
	if updateOperator.OperatorName == "" {
		return http.StatusBadRequest, errors.New("Name is required")
	}
	err := repository.IsOperatorIdValid(db, operatorId)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return http.StatusOK, nil
}

func DeleteOperatorsValidation(db *sql.DB, TokenIsAdmin bool, TokenIsOperator bool, operatorId string) (int, error) {
	if !TokenIsAdmin && !TokenIsOperator {
		return http.StatusForbidden, errors.New("Donot have access to update the Operator")
	}
	err := repository.IsOperatorIdValid(db, operatorId)
	if err != nil {
		return http.StatusBadRequest, err
	}
	return 204, nil
}
