package repository

import (
	models "github.com/chandanaavadhani/BusService/models"
)

func InsertReview(ratingid string, userid string, review models.Review) error {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Insert review into Database
	stmt, err := db.Prepare("INSERT INTO reviews (ratingID, userID, busID,comment,rating) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ratingid, userid, review.BusId, review.Comment, review.Rating)
	if err != nil {
		return err
	}

	return nil
}

func GetTripsBasedOnRoute(route models.GetTripsRequest) ([]models.Trips, error) {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get trips from the database

	// Execute the SQL query
	query := `
        SELECT * FROM trips
        WHERE DATE(FROM_UNIXTIME(departure)) = ?
        AND EXISTS (
            SELECT * FROM route
            WHERE route.source = ?
            AND route.destination = ?
            AND route.routeID = trips.routeID
        )
    `
	rows, err := db.Query(query, route.Date, route.Source, route.Destination)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var trips []models.Trips
	for rows.Next() {
		var trip models.Trips
		err := rows.Scan(&trip.TripId, &trip.BusId, &trip.RouteId, &trip.Departure, &trip.Arrival, &trip.Capacity, &trip.Cost, &trip.BusStatus)
		if err != nil {
			return nil, err
		}
		trips = append(trips, trip)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return trips, nil

}

func GetTripDetails(tripId string) ([]models.TripDetails, error) {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

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
