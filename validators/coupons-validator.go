package validators

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
)

func CreateCouponValidations(promo models.Coupon) (int, error) {

	//Validating Coupon Code
	if promo.CouponCode == " " || promo.CouponCode == "" {
		return http.StatusBadRequest, errors.New("coupon Code is required")
	}
	if len(promo.CouponCode) != 6 {
		return http.StatusBadRequest, errors.New("coupon Code Length of 6 is required")
	}
	numbers := regexp.MustCompile(`^[0-9]$`)
	alphabets := regexp.MustCompile(`^[a-zA-Z]{5}$`)

	if !numbers.MatchString(string(promo.CouponCode[5])) {
		return http.StatusBadRequest, errors.New("last element is not a number")
	}
	if !alphabets.MatchString(promo.CouponCode[:5]) {
		return http.StatusBadRequest, errors.New("first five elements are not alphabets")
	}
	count, _, err := repository.IsCouponCodeExists(promo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count != 0 {
		return http.StatusBadRequest, errors.New("coupon code already exists")
	}

	//Validating Coupon Amount
	if promo.CouponAmount > 1000.00 || promo.CouponAmount < 25.00 {
		return http.StatusBadRequest, errors.New("coupon Amount should be more than 25 and less than 1000")
	}
	return 200, nil
}

func UpdateValidations(promo models.UpdateCoupon) (int, error) {

	//Validating Coupon Code
	count, _, err := repository.IsUpdatedCouponExists(promo)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("coupon code not found")
	}
	if promo.CouponCode == " " || promo.CouponCode == "" {
		return http.StatusBadRequest, errors.New("coupon Code is required")
	}

	//Validating New Coupon Amount
	if promo.NewCouponAmount > 1000.00 || promo.NewCouponAmount < 25.00 {
		return http.StatusBadRequest, errors.New("coupon Amount should be more than 25 and less than 1000")
	}
	return 200, nil
}

func DeleteOrGetValidations(coupon string) (int, error) {
	count, err := repository.IsCouponExists(coupon)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	if count == 0 {
		return http.StatusBadRequest, errors.New("coupon code not found")
	}
	return 200, nil
}
