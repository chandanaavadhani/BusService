package validators

import (
	"fmt"

	"github.com/chandanaavadhani/BusService/models"
)

func ValidateAddBusRequest(request models.AddBusRequest) error {
	if request.BusNumber == "" {
		return fmt.Errorf("invalid bus number")
	} else if request.Contact == "" {
		return fmt.Errorf("invalid contact details")
	} else if request.Capacity == 0 {
		return fmt.Errorf("invalid capacity")
	} else if request.BusType != "A/C" && request.BusType != "Non A/C" {
		return fmt.Errorf("invalid Bus Type")
	}
	return nil
}
