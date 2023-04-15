package validators

import (
	"errors"
	"net/http"

	"github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
)

func ValidateBookingRequest(paymentId string, booking models.Bookings) (int, error) {

	if booking.TripId == "" {
		return http.StatusBadRequest, errors.New("Trip ID missing")
	}
	if booking.CouponCode == "" {
		return http.StatusBadRequest, errors.New("Coupon Code missing")
	}
	if booking.Passengers == "" {
		return http.StatusBadRequest, errors.New("Passengers data missing")
	}
	if booking.BookingStatus == "" {
		return http.StatusBadRequest, errors.New("Booking Status missing")
	}
	if booking.PaymentStatus == "" {
		return http.StatusBadRequest, errors.New("Payment Status missing")
	}
	if booking.Method == "" {
		return http.StatusBadRequest, errors.New("Payment Method missing")
	}
	if booking.AmountPaid == 0 {
		return http.StatusBadRequest, errors.New("Payment amount missing")
	}
	if repository.CheckIfTripExists(booking.TripId) != true {
		return http.StatusNotFound, errors.New("Invalid Trip ID")
	}
	if repository.CheckIfCouponExists(booking.CouponCode) != true {
		return http.StatusNotFound, errors.New("Invalid coupon code")
	}

	tripCost := repository.GetTripCost(booking.TripId)
	discount := repository.GetCouponAmount(booking.CouponCode)
	if tripCost-discount != booking.AmountPaid {
		return http.StatusBadRequest, errors.New("Invalid amount")
	}

	return 200, nil
}
