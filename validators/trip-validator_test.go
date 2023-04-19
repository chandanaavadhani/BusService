package validators

import (
	"testing"

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
