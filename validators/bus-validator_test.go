package validators

import (
	"testing"

	"github.com/chandanaavadhani/BusService/models"
)

func TestValidateAddBusRequest(t *testing.T) {
	req := models.AddBusRequest{
		BusType:   "A/C",
		Capacity:  10,
		Contact:   "1234567890",
		BusNumber: "SAS1886",
	}

	//Happy pass - no errors
	actual := ValidateAddBusRequest(req)
	var expected error

	if actual != expected {
		t.Errorf("Expected %q doesn't match with actual result %q", actual, expected)
	}
}
