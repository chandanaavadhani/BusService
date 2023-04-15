package validators

import (
	"database/sql"
	"fmt"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
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

func ValidateBusExistence(busid string, db *sql.DB) bool {
	return repository.CheckIfBusExists(busid, db)
}
