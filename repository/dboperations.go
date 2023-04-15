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

func GetAllBusses(db *sql.DB) ([]models.BusDetails, error) {
	var allBusses []models.BusDetails
	//get rows from DB
	rows, err := db.Query(`SELECT buses.*, reviews.ratingID,reviews.userID, reviews.comment,reviews.rating
	FROM buses
	LEFT JOIN reviews
	ON buses.busID = reviews.busID`)
	if err != nil {
		return allBusses, err
	}

	var busDetails models.BusDetails
	for rows.Next() {
		//Declaring variables because reviews can be null
		var ratingID sql.NullString
		var rating sql.NullInt64
		var comment sql.NullString
		var userid sql.NullString

		err = rows.Scan(&busDetails.BusId, &busDetails.OperatorId, &busDetails.Contact, &busDetails.Capacity, &busDetails.BusType, &busDetails.BusNumber, &ratingID, &userid, &comment, &rating)

		if err != nil {
			return allBusses, err
		}

		index := checkIfBusIdExistsInData(allBusses, busDetails.BusId)

		//Building review
		var review models.ReviewResponseFromDBForBus
		review.RatingID = ratingID.String
		review.Comment = comment.String
		review.UserID = userid.String
		review.Rating = rating.Int64

		if index == -1 {
			//If There are no rating for this bus is null in DB
			if !rating.Valid {
				allBusses = append(allBusses, busDetails)
				continue
			}
			//If bus doesn't exist in that array , add bus to that array
			busDetails.Reviews = append(busDetails.Reviews, review)

			allBusses = append(allBusses, busDetails)
			continue
		}

		//If bus exists , add review to reviews array of that bus details
		allBusses[index].Reviews = append(allBusses[index].Reviews, review)
	}

	// Handle any errors that occurred during iteration
	err = rows.Err()
	if err != nil {
		return allBusses, err
	}

	return allBusses, nil
}

func checkIfBusIdExistsInData(data []models.BusDetails, busid string) int {
	for i := 0; i < len(data); i++ {
		if data[i].BusId == busid {
			return i
		}
	}
	return -1
}

func GetAllBussesOfOperator(operatorId string, db *sql.DB) ([]models.BusDetails, error) {
	var allBusses []models.BusDetails
	//get rows from DB
	rows, err := db.Query(`SELECT buses.*, reviews.ratingID,reviews.userID, reviews.comment,reviews.rating
	FROM buses
	LEFT JOIN reviews
	ON buses.busID = reviews.busID
	WHERE buses.operatorID = ?`, operatorId)
	if err != nil {
		return allBusses, err
	}

	var busDetails models.BusDetails
	for rows.Next() {
		//Declaring variables because reviews can be null
		var ratingID sql.NullString
		var rating sql.NullInt64
		var comment sql.NullString
		var userid sql.NullString

		err = rows.Scan(&busDetails.BusId, &busDetails.OperatorId, &busDetails.Contact, &busDetails.Capacity, &busDetails.BusType, &busDetails.BusNumber, &ratingID, &userid, &comment, &rating)

		if err != nil {
			return allBusses, err
		}

		index := checkIfBusIdExistsInData(allBusses, busDetails.BusId)

		//Building review
		var review models.ReviewResponseFromDBForBus
		review.RatingID = ratingID.String
		review.Comment = comment.String
		review.UserID = userid.String
		review.Rating = rating.Int64

		if index == -1 {
			//If There are no rating for this bus is null in DB
			if !rating.Valid {
				allBusses = append(allBusses, busDetails)
				continue
			}
			//If bus doesn't exist in that array , add bus to that array
			busDetails.Reviews = append(busDetails.Reviews, review)

			allBusses = append(allBusses, busDetails)
			continue
		}

		//If bus exists , add review to reviews array of that bus details
		allBusses[index].Reviews = append(allBusses[index].Reviews, review)
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

func GetBusById(busid string, db *sql.DB) (models.BusDetails, error) {
	var busDetails models.BusDetails
	busid += "\n"
	//get all ratings for that busID
	rows, err := db.Query(`SELECT buses.*, reviews.ratingID,reviews.userID, reviews.comment,reviews.rating
							FROM buses
							LEFT JOIN reviews
							ON buses.busID = reviews.busID
							WHERE buses.busID =?`, busid)

	if err != nil {
		return busDetails, err
	}
	defer rows.Close()
	var count int
	var review models.ReviewResponseFromDBForBus

	//Iterate through all the rows
	for rows.Next() {
		var RatingID sql.NullString
		var Rating sql.NullInt64
		var comment sql.NullString
		var userid sql.NullString

		err = rows.Scan(&busDetails.BusId, &busDetails.OperatorId, &busDetails.Contact, &busDetails.Capacity, &busDetails.BusType, &busDetails.BusNumber, &RatingID, &userid, &comment, &Rating)

		if err != nil {
			return busDetails, err
		}
		//check if rating is null in db
		if RatingID.Valid {
			review.Comment = comment.String
			review.Rating = Rating.Int64
			review.RatingID = RatingID.String
			review.UserID = userid.String
			busDetails.Reviews = append(busDetails.Reviews, review)
		}

		count++
	}

	if count == 0 {
		return busDetails, fmt.Errorf("no data for that busID")
	}

	return busDetails, nil
}

func GetTripDetails(tripId string, db *sql.DB) ([]models.TripDetails, error) {

	// get trips from the database

	// Execute the SQL query
	query := `
	SELECT 
	t.*,
	b.driverContact,
	b.capacity,
	b.busNumber,
	b.busType,
	b.operatorID,
	r.source,
	r.destination,
	r.distance
	FROM trips t
	JOIN buses b ON t.busID = b.busID
	JOIN route r ON t.routeID = r.routeID 
	WHERE t.tripID = ?
    `
	rows, err := db.Query(query, tripId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tripDetails []models.TripDetails
	for rows.Next() {
		var tripDetail models.TripDetails
		err := rows.Scan(&tripDetail.TripId, &tripDetail.BusId, &tripDetail.RouteId, &tripDetail.Departure, &tripDetail.Arrival, &tripDetail.Capacity, &tripDetail.Cost, &tripDetail.BusStatus,
			&tripDetail.DriverContact, &tripDetail.BusCapacity, &tripDetail.BusType, &tripDetail.BusNumber, &tripDetail.OperatorId,
			&tripDetail.Destination, &tripDetail.Source, &tripDetail.Distance)
		if err != nil {
			return nil, err
		}
		tripDetails = append(tripDetails, tripDetail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tripDetails, nil
}

func UpdateTripDetails(request models.UpdateTripRequest, db *sql.DB) error {
	// Prepare the SQL statement
	stmt, err := db.Prepare("UPDATE trips SET routeID=?, busID=?, busStatus=?, departure=?, arrival=?, cost=? WHERE tripID=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Execute the SQL statement
	_, err = stmt.Exec(request.RouteId, request.BusId, request.BusStatus, request.Departure, request.Arrival, request.Cost, request.TripId)
	if err != nil {
		return err
	}

	return nil
}

func GetAllTripsDetails(db *sql.DB) ([]models.TripDetails, error) {

	// get trips from the database

	// Execute the SQL query
	query := `
	SELECT 
	t.*,
	b.driverContact,
	b.capacity,
	b.busNumber,
	b.busType,
	b.operatorID,
	r.source,
	r.destination,
	r.distance
	FROM trips t
	JOIN buses b ON t.busID = b.busID
	JOIN route r ON t.routeID = r.routeID 
    `
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tripDetails []models.TripDetails
	for rows.Next() {
		var tripDetail models.TripDetails
		err := rows.Scan(&tripDetail.TripId, &tripDetail.BusId, &tripDetail.RouteId, &tripDetail.Departure, &tripDetail.Arrival, &tripDetail.Capacity, &tripDetail.Cost, &tripDetail.BusStatus,
			&tripDetail.DriverContact, &tripDetail.BusCapacity, &tripDetail.BusType, &tripDetail.BusNumber, &tripDetail.OperatorId,
			&tripDetail.Destination, &tripDetail.Source, &tripDetail.Distance)
		if err != nil {
			return nil, err
		}
		tripDetails = append(tripDetails, tripDetail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tripDetails, nil
}

func GetAllTripsDetailsOfOperator(operatorid string, db *sql.DB) ([]models.TripDetails, error) {

	// get trips from the database

	// Execute the SQL query
	query := `
	SELECT 
	t.*,
	b.driverContact,
	b.capacity,
	b.busNumber,
	b.busType,
	b.operatorID,
	r.source,
	r.destination,
	r.distance
	FROM trips t
	JOIN buses b ON t.busID = b.busID
	JOIN route r ON t.routeID = r.routeID
	WHERE b.operatorID = ?
    `
	rows, err := db.Query(query, operatorid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tripDetails []models.TripDetails
	for rows.Next() {
		var tripDetail models.TripDetails
		err := rows.Scan(&tripDetail.TripId, &tripDetail.BusId, &tripDetail.RouteId, &tripDetail.Departure, &tripDetail.Arrival, &tripDetail.Capacity, &tripDetail.Cost, &tripDetail.BusStatus,
			&tripDetail.DriverContact, &tripDetail.BusCapacity, &tripDetail.BusType, &tripDetail.BusNumber, &tripDetail.OperatorId,
			&tripDetail.Destination, &tripDetail.Source, &tripDetail.Distance)
		if err != nil {
			return nil, err
		}
		tripDetails = append(tripDetails, tripDetail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tripDetails, nil
}
