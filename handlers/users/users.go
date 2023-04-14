package handlers

import (
	"encoding/json"
	"net/http"

	models "github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	"github.com/chandanaavadhani/BusService/validators"
)

// func Signup(w http.ResponseWriter, r *http.Request) {
// 	if r.Method != http.MethodPost {
// 		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
// 		return
// 	}
// 	// Decode the JSON request body into a LoginRequest struct.
// 	var user models.Signup
// 	err := json.NewDecoder(r.Body).Decode(&user)
// 	if err != nil {
// 		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), nil)
// 		return
// 	}
// 	// Validate the user input.
// 	statusCode, err := validators.SignupValidation(user)
// 	if err != nil {
// 		utils.BuildResponse(w, statusCode, err.Error(), nil)
// 		return
// 	}
// 	//Insert user into DB
// 	userid, err := repository.InsertUser(user)
// 	if err != nil {
// 		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
// 		return
// 	}

// 	token, err := utils.GenerateToken(userid, false, false)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// 	utils.BuildResponse(w, http.StatusCreated, "signup successful", map[string]string{
// 		"token": token,
// 	})

// }
func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	// Decode the JSON request body into a LoginRequest struct.
	var user models.Signup
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	// Validate the user input.
	statusCode, err := validators.SignupValidation(user)
	if err != nil {
		utils.BuildResponse(w, statusCode, err.Error(), nil)
		return
	}
	//Insert user into DB
	userid, err := repository.InsertUser(user)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	token, err := utils.GenerateToken(userid, false, false)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusCreated, "signup successful", map[string]string{
		"token": token,
	})
}
func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	var userlogin models.Login
	err := json.NewDecoder(r.Body).Decode(&userlogin)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}
	statusCode, err := validators.LoginValidation(userlogin)
	if err != nil {
		utils.BuildResponse(w, statusCode, err.Error(), nil)
		return
	}
	userid, Isadmin, Isoperator, statusCode, err := repository.LoginUser(userlogin)
	if err != nil {
		utils.BuildResponse(w, statusCode, err.Error(), nil)
		return
	}
	token, err := utils.GenerateToken(userid, Isadmin, Isoperator)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusCreated, "Login successful", map[string]string{
		"token": token,
	})
}
func UpdatePassword(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method Not Allowed", nil)
		return
	}
	var person models.Login
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

}
