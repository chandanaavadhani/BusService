package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
	utils "github.com/chandanaavadhani/BusService/utils"
	validators "github.com/chandanaavadhani/BusService/validators"
)

func CreateOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}

	//reading the body
	var newOperator models.Operator
	err := json.NewDecoder(r.Body).Decode(&newOperator)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	//db connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "DB Connection Failed", nil)
		return
	}
	defer db.Close()
	//validate
	if ErrorCode, err := validators.SignupOperatorsValidation(db, newOperator); err != nil {
		utils.BuildResponse(w, ErrorCode, err.Error(), nil)
		return
	}
	//call the query to execute
	if Errorcode, err := repository.InsertOperator(db, newOperator); err != nil {
		utils.BuildResponse(w, Errorcode, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusCreated, fmt.Sprintf("Operator %s created successfully", newOperator.Email), nil)
}

func UpdateOperator(w http.ResponseWriter, r *http.Request) {

	if r.Method != "PUT" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}

	//db connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "DB Connection Failed", nil)
		return
	}
	defer db.Close()

	//reading the body
	var updateOperator models.UpdateOperator
	err = json.NewDecoder(r.Body).Decode(&updateOperator)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	//token details
	// TokenId, ok := r.Context().Value("Id").(string)
	// TokenIsAdmin, ok := r.Context().Value("IsAdmin").(bool)
	// TokenIsOperator, ok := r.Context().Value("IsOperator").(bool)
	TokenIsAdmin := false
	TokenIsOperator := true
	//operator Id from path
	operatorId := strings.Split(r.URL.Path, "/operator/update/")[1]

	//validate
	if ErrorCode, err := validators.UpdateOperatorsValidation(db, updateOperator, TokenIsAdmin, TokenIsOperator, operatorId); err != nil {
		utils.BuildResponse(w, ErrorCode, err.Error(), nil)
		return
	}
	//execute
	if Errorcode, err := repository.UpdateOperatorDetails(db, updateOperator, operatorId); err != nil {
		utils.BuildResponse(w, Errorcode, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusOK, "Operator is updated Successfully", nil)

}

func DeleteOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	//db connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "DB Connection Failed", nil)
		return
	}
	defer db.Close()
	// TokenId, ok := r.Context().Value("Id").(string)
	// TokenIsAdmin, ok := r.Context().Value("IsAdmin").(bool)
	// TokenIsOperator, ok := r.Context().Value("IsOperator").(bool)
	TokenIsAdmin := false
	TokenIsOperator := true
	operatorId := strings.Split(r.URL.Path, "/operator/delete/")[1]

	//validate
	if ErrorCode, err := validators.DeleteOperatorsValidation(db, TokenIsAdmin, TokenIsOperator, operatorId); err != nil {
		utils.BuildResponse(w, ErrorCode, err.Error(), nil)
		return
	}
	//execute
	if Errorcode, err := repository.DeleteOperatorDetails(db, operatorId); err != nil {
		utils.BuildResponse(w, Errorcode, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusOK, "Operator is Deleted Successfully", nil)
}

func GetOperator(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	//db connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "DB Connection Failed", nil)
		return
	}
	defer db.Close()
	// TokenId, ok := r.Context().Value("Id").(string)
	// TokenIsAdmin, ok := r.Context().Value("IsAdmin").(bool)
	// TokenIsOperator, ok := r.Context().Value("IsOperator").(bool)
	TokenIsAdmin := false
	TokenIsOperator := true
	operatorId := strings.Split(r.URL.Path, "/operator/get/")[1]
	//validate
	if ErrorCode, err := validators.DeleteOperatorsValidation(db, TokenIsAdmin, TokenIsOperator, operatorId); err != nil {
		utils.BuildResponse(w, ErrorCode, err.Error(), nil)
		return
	}
	//execute
	Errorcode, err, operator := repository.GetOperatorDetails(db, operatorId)
	if err != nil {
		utils.BuildResponse(w, Errorcode, err.Error(), nil)
		return
	}

	utils.BuildResponse(w, http.StatusOK, "Operator get Response", operator)

}

func GetAllOperators(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	//db connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "DB Connection Failed", nil)
		return
	}
	defer db.Close()
	// TokenId, ok := r.Context().Value("Id").(string)
	// TokenIsAdmin, ok := r.Context().Value("IsAdmin").(bool)
	// TokenIsOperator, ok := r.Context().Value("IsOperator").(bool)
	TokenIsAdmin := true
	// TokenIsOperator := true
	//validate
	if !TokenIsAdmin {
		utils.BuildResponse(w, http.StatusForbidden, "Donot have access to update the Operator", nil)
		return
	}
	//execute
	Errorcode, err, operators := repository.GetAllOperatorDetails(db)
	if err != nil {
		utils.BuildResponse(w, Errorcode, err.Error(), nil)
		return
	}

	utils.BuildResponse(w, http.StatusOK, "Operator get Response", operators)
}
