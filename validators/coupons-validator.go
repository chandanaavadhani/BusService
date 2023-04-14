package validators

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
)

func CreateValidations(promo models.Coupon) (int, error) {
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
	if promo.CouponAmount > 1000.00 || promo.CouponAmount < 25.00 {
		return http.StatusBadRequest, errors.New("coupon Amount should be between 25-1000")
	}
	return 200, nil
}
