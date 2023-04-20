package validators

import (
	"strconv"
	"testing"
	"time"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
)

func TestValidateAddTripRequest(t *testing.T) {
	req := models.AddTripRequest{
		BusId:     "4B2A7D78-57C9-4417-B898-3DD4DB0CC665\n",
		RouteId:   "testing",
		Departure: "1682476989",
		Arrival:   "1682491389",
		Cost:      25.99,
		BusStatus: "available",
	}
	db, err := repository.DBConnection()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	//Happy pass - no errors
	actual := ValidateAddTripRequest(req, db)
	var expected error

	if actual != expected {
		t.Errorf("Expected %q doesn't match with actual result %q", actual, expected)
	}

}

func TestIsFuture(t *testing.T) {
	currentTime := time.Now().Unix()

	// future timestamp
	futureTimestamp := strconv.FormatInt(currentTime+3600, 10)
	if !isFuture(futureTimestamp) {
		t.Errorf("isFuture(%v) = false; want true", futureTimestamp)
	}

	// past timestamp
	pastTimestamp := strconv.FormatInt(currentTime-3600, 10)
	if isFuture(pastTimestamp) {
		t.Errorf("isFuture(%v) = true; want false", pastTimestamp)
	}

	// invalid timestamp
	invalidTimestamp := "invalid-timestamp"
	if isFuture(invalidTimestamp) {
		t.Errorf("isFuture(%v) = true; want false", invalidTimestamp)
	}
}

func TestIsArrivalFutureOfDeparture(t *testing.T) {
	// future arrival timestamp and past departure timestamp
	arrivalTimestamp := strconv.FormatInt(time.Now().Unix()+3600, 10)
	departureTimestamp := strconv.FormatInt(time.Now().Unix()-3600, 10)
	if !isArrivalFutureOfDeparture(arrivalTimestamp, departureTimestamp) {
		t.Errorf("isArrivalFutureOfDeparture(%v, %v) = false; want true", arrivalTimestamp, departureTimestamp)
	}

	// past arrival timestamp and future departure timestamp
	arrivalTimestamp = strconv.FormatInt(time.Now().Unix()-3600, 10)
	departureTimestamp = strconv.FormatInt(time.Now().Unix()+3600, 10)
	if isArrivalFutureOfDeparture(arrivalTimestamp, departureTimestamp) {
		t.Errorf("isArrivalFutureOfDeparture(%v, %v) = true; want false", arrivalTimestamp, departureTimestamp)
	}

	// invalid timestamps
	if isArrivalFutureOfDeparture("invalid-timestamp", "invalid-timestamp") {
		t.Errorf("isArrivalFutureOfDeparture(%v, %v) = true; want false", "invalid-timestamp", "invalid-timestamp")
	}
}
