package handlers

import (
	"encoding/json"
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

	//Get all coupons
	var promo []models.Coupon
	promo, err := repository.GetAllCoupons()
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusOK, "List of Coupons", promo)
}

func DeleteOrGetCoupon(w http.ResponseWriter, r *http.Request) {

	if r.Method != "DELETE" && r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not Allowed", nil)
		return
	}

	//Get couponCode from Params
	coupon := r.URL.Query().Get("couponcode")

	//Validate the coupon
	status, err := validators.DeleteOrGetValidations(coupon)
	if err != nil {
		utils.BuildResponse(w, status, err.Error(), nil)
		return
	}

	if r.Method == "GET" {
		//Get Coupon Code
		var promo models.Coupon
		promo, err = repository.GetCoupon(coupon)
		if err != nil {
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		utils.BuildResponse(w, http.StatusOK, "Coupon code is retreived", promo)

	} else if r.Method == "DELETE" {
		//Delete the Coupon
		err = repository.DeleteCoupon(coupon)
		if err != nil {
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
			return
		}
		utils.BuildResponse(w, http.StatusOK, "Coupon Deleted", nil)

	}
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
	status, err := validators.CreateCouponValidations(coupon)
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
}

func UpdateCoupon(w http.ResponseWriter, r *http.Request) {
	var coupon models.UpdateCoupon
	if r.Method != "PUT" {
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
	status, err := validators.UpdateValidations(coupon)
	if err != nil {
		utils.BuildResponse(w, status, err.Error(), nil)
		return
	}

	//Insert the Data
	err = repository.UpdateCoupon(coupon)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	utils.BuildResponse(w, http.StatusOK, "Coupon Updated", nil)
}
