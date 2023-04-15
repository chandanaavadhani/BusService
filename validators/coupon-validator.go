package validators

import (
	"errors"
	"net/http"

	repository "github.com/chandanaavadhani/BusService/repository"
)

func ValidateCouponCode(tripId string) (int, error) {
	if tripId == "" || tripId == "\n" {
		return http.StatusBadRequest, errors.New("Missing Coupon Code")
	}
	if repository.CheckIfCouponExists(tripId) != true {
		return http.StatusBadRequest, errors.New("Invalid Coupon Code")
	}

	return 200, nil
}
