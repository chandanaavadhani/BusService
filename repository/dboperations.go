package repository

import (
	"database/sql"
	"fmt"

	"github.com/chandanaavadhani/BusService/models"

	"os/exec"
)

func generateUniqueID() (string, error) {
	newUUID, err := exec.Command("uuidgen").Output()
	if err != nil {
		return "", err
	}
	return string(newUUID), nil
}

func AddBusToDB(bus models.AddBusRequest, operatorId string, db *sql.DB) error {
	//generate unique id for Bus
	busId, err := generateUniqueID()
	if err != nil {
		return err
	}
	fmt.Println("This is the generated BusID : ", busId)
	//Add bus to DB
	_, err = db.Exec("INSERT INTO buses (busID,busType,capacity,driverContact,operatorID,busNumber) VALUES(?,?,?,?,?,?)", busId, bus.BusType, bus.Capacity, bus.Contact, operatorId, bus.BusNumber)
	if err != nil {
		fmt.Println("Error is : ", err.Error())
		return fmt.Errorf("failed to add bus, try again")
	}
	return nil
}

func GetAllBusses(db *sql.DB) ([]models.Bus, error) {
	var allBusses []models.Bus
	//get rows from DB
	rows, err := db.Query("SELECT * FROM buses")
	if err != nil {
		return allBusses, err
	}

	for rows.Next() {
		var row models.Bus
		err = rows.Scan(&row.BusId, &row.OperatorId, &row.Contact, &row.Capacity, &row.BusType, &row.BusNumber)
		if err != nil {
			return allBusses, err
		}

		allBusses = append(allBusses, row)
	}

	// Handle any errors that occurred during iteration
	err = rows.Err()
	if err != nil {
		return allBusses, err
	}

	return allBusses, nil
}

func GetAllBussesOfOperator(operatorId string, db *sql.DB) ([]models.Bus, error) {
	var allBusses []models.Bus
	//get rows from DB
	rows, err := db.Query("SELECT * FROM buses WHERE operatorID = ?", operatorId)
	if err != nil {
		return allBusses, err
	}

	for rows.Next() {
		var row models.Bus
		err = rows.Scan(&row.BusId, &row.OperatorId, &row.Contact, &row.Capacity, &row.BusType, &row.BusNumber)
		if err != nil {
			return allBusses, err
		}

		allBusses = append(allBusses, row)
	}

	// Handle any errors that occurred during iteration
	err = rows.Err()
	if err != nil {
		return allBusses, err
	}

	return allBusses, nil
}

func CheckIfBusExists(busid string, db *sql.DB) bool {
	//check if busId exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM buses WHERE busID = ?)", busid).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return true
}

func CheckIfRouteExists(routeid string, db *sql.DB) bool {
	//check if routeId exists
	var exists bool
	err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM route WHERE routeID = ?)", routeid).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false
	}
	return true
}

func GetCapacityOfBus(busid string, db *sql.DB) (int, error) {
	var capacity int
	err := db.QueryRow("SELECT capacity FROM buses WHERE busID=?", busid).Scan(&capacity)
	fmt.Println("Capacity : ", capacity)
	return capacity, err
}

func AddTriptoDB(request models.AddTripRequest, capacity int, db *sql.DB) error {

	//generate unique id for trip
	tripID, err := generateUniqueID()
	if err != nil {
		return err
	}
	fmt.Println("This is the generated TripID : ", tripID)

	//Add Trip to DB
	_, err = db.Exec("INSERT INTO trips (tripID,busID,routeID,departure,arrival,capacity,cost,busStatus) VALUES(?,?,?,?,?,?,?,?)", tripID, request.BusId, request.RouteId, request.Departure, request.Arrival, capacity, request.Cost, request.BusStatus)
	if err != nil {
		fmt.Println("Error is : ", err.Error())
		return fmt.Errorf("failed to add bus, try again")
	}
	return nil

}
