package validators

import (
	"database/sql"
	"errors"
	"log"
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
	if isTripIDPresent(tripId) != true {
		return http.StatusBadRequest, errors.New("Invalid Trip ID")
	}
	return 200, nil
}

func isTripIDPresent(tripID string) bool {

	db, err := repository.DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT tripID FROM trips WHERE tripID = ?"

	// Execute the query and check if a record is returned
	var result string
	err = db.QueryRow(query, tripID).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}
