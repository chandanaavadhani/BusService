package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	"github.com/chandanaavadhani/BusService/validators"
)

func GetAllCoupons(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func GetCoupon(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello")
}

func CreateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.Coupon
	if r.Method != "POST" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}

	//Get the Data
	err := json.NewDecoder(r.Body).Decode(&coupon)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, "Decoding Error", nil)
		return
	}

	//Validate the Data
	status, err := validators.CreateValidations(coupon)
	if err != nil {
		utils.BuildResponse(w, status, err.Error(), nil)
		return
	}

	//Insert the Data
	err = repository.InsertCoupon(coupon)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusCreated, "Coupon Created", nil)
	w.WriteHeader(http.StatusCreated)
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}
	fmt.Println("Hello")
}
