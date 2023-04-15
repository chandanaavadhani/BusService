package validators

import (
	"errors"
	"net/http"
	"time"

	models "github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
)

func ValidateRouteDetails(route models.GetTripsRequest) (int, error) {

	if route.Source == "" {
		return http.StatusBadRequest, errors.New("Source address missing")
	}
	if route.Destination == "" {
		return http.StatusBadRequest, errors.New("Destination address missing")
	}
	if route.Date == "" {
		return http.StatusBadRequest, errors.New("Date missing")
	}
	if ValidateDate(route.Date) != true {
		return http.StatusBadRequest, errors.New("Invalid Date")
	}

	return 200, nil
}

func ValidateDate(date string) bool {

	// Parse the date string into a time.Time object
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return false // Invalid date format
	}

	// Get the current time
	now := time.Now()

	// Validate Date
	if t.After(now) {
		return true
	} else {
		return false
	}
}

func ValidateTripId(tripId string) (int, error) {
	if tripId == "" || tripId == "\n" {
		return http.StatusBadRequest, errors.New("Missing TripId")
	}
	if repository.CheckIfTripExists(tripId) != true {
		return http.StatusBadRequest, errors.New("Invalid TripId")
	}

	return 200, nil
}
