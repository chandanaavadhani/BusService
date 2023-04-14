package validators

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
)

func ValidateAddTripRequest(request models.AddTripRequest, db *sql.DB) error {
	if request.BusId == "" || !repository.CheckIfBusExists(request.BusId, db) {
		return fmt.Errorf("invalid bus id")
	} else if request.RouteId == "" || !repository.CheckIfRouteExists(request.RouteId, db) {
		return fmt.Errorf("invalid route id")
	} else if request.Cost == 0 {
		return fmt.Errorf("invalid cost")
	} else if request.Departure == "" || !isFuture(request.Departure) {
		return fmt.Errorf("invalid departure timestamp")
	} else if request.Arrival == "" {
		return fmt.Errorf("invalid arrival timestamp")
	} else if !isArrivalFutureOfDeparture(request.Arrival, request.Departure) {
		return fmt.Errorf("arrival should be ahead of departure")
	}
	return nil
}

func isFuture(timestamp string) bool {
	tsInt, err := strconv.ParseInt(timestamp, 10, 64)
	if err != nil {
		// handle error
		return false
	}
	currentTime := time.Now().Unix()
	return tsInt > currentTime
}

func isArrivalFutureOfDeparture(timestamp1, timestamp2 string) bool {
	ts1Int, err := strconv.ParseInt(timestamp1, 10, 64)
	if err != nil {
		// handle error
		return false
	}
	ts2Int, err := strconv.ParseInt(timestamp2, 10, 64)
	if err != nil {
		// handle error
		return false
	}
	return ts1Int > ts2Int
}
