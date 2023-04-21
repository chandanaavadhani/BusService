package repository

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

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

	// get trips deatils from the database by tripID
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

func InsertPaymentDetails(paymentId string, booking models.Bookings) error {
	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Generate a timestamp
	timestamp := time.Now().Unix()
	fmt.Println(timestamp)

	//Insert review into Database
	stmt, err := db.Prepare("INSERT INTO payments (paymentID, paymentStatus, paymentType, amountPaid, paymentDate ) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(paymentId, booking.PaymentStatus, booking.Method, booking.AmountPaid, timestamp)
	if err != nil {
		return err
	}

	return nil
}

func InsertBookingDetails(bookingId string, userId string, paymentId string, booking models.Bookings) error {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Insert review into Database
	stmt, err := db.Prepare("INSERT INTO bookings (bookingID, userID, tripID,paymentID,couponCode, passengers, bookingStatus) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(bookingId, userId, booking.TripId, paymentId, booking.CouponCode, booking.Passengers, booking.BookingStatus)
	if err != nil {
		return err
	}

	return nil
}

func CheckIfTripExists(tripId string) bool {
	//DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT tripID FROM trips WHERE tripID = ?"

	// Execute the query and check if a record is returned
	var result string
	err = db.QueryRow(query, tripId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}

func CheckIfReviewAdded(busId string, userId string) bool {
	//DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Construct the SQL query with placeholders for the input values
	query := `SELECT * FROM reviews WHERE busID = ? AND userID = ?;`

	// Execute the query with the input values
	rows, err := db.Query(query, busId, userId)
	if err != nil {
		return false
	}
	defer rows.Close()

	// Check if any rows were returned by the query
	exists := false
	for rows.Next() {
		exists = true
	}

	// Return true if any rows were returned, false otherwise
	return exists
}

func CheckIfCouponExists(coupon string) bool {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT couponCode FROM discounts WHERE couponCode = ?"

	// Execute the query and check if a record is returned
	var result string
	err = db.QueryRow(query, coupon).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}

func CheckIfPaymentExists(paymentId string) bool {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT paymentID FROM payments WHERE paymentID = ?"

	// Execute the query and check if a record is returned
	var result string
	err = db.QueryRow(query, paymentId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}
func CheckIfBookingIDExists(bookingId string) bool {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT bookingID FROM bookings WHERE bookingID = ?"

	// Execute the query and check if a record is returned
	var result string
	err = db.QueryRow(query, bookingId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return true
}

func GetTripCost(tripId string) float64 {
	//DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT cost FROM trips WHERE tripId = ?"
	var result string
	err = db.QueryRow(query, tripId).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound
		} else {
			log.Fatal(err)
		}
	}
	cost, err := strconv.ParseFloat(result, 64)
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error: could not convert string to int64")
	}
	return cost
}

func GetCouponAmount(couponCode string) float64 {
	//DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Prepare the SQL query
	query := "SELECT cost FROM trips WHERE couponcode = ?"
	var result string
	err = db.QueryRow(query, couponCode).Scan(&result)
	if err != nil {
		if err == sql.ErrNoRows {
			return http.StatusNotFound
		} else {
			log.Fatal(err)
		}
	}
	couponAmount, err := strconv.ParseFloat(result, 64)
	if err != nil {
		// Handle the error if the string cannot be converted to an integer
		fmt.Println("Error: could not convert string to int64")
	}
	return couponAmount
}

func GetBookingDetails(bookingId string) ([]models.BookingDetails, error) {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// get booking details from the database
	// Execute the SQL query
	query := `
	SELECT 
	b.*,
	py.paymentStatus,
	py.paymentType,
	py.amountPaid,
	py.paymentDate,
	p.PassengerName,
	p.Age,
	p.Gender,
	p.Contact
	FROM bookings b
	JOIN passengers p ON p.PassengerId = b.passengers
	JOIN payments py ON py.paymentID = b.paymentID
	WHERE b.bookingId = ?
    `
	rows, err := db.Query(query, bookingId)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	var bookingDetails []models.BookingDetails
	for rows.Next() {
		var bookingDetail models.BookingDetails
		err := rows.Scan(&bookingDetail.BookingId, &bookingDetail.UserId, &bookingDetail.TripId, &bookingDetail.PaymentId,
			&bookingDetail.CouponCode, &bookingDetail.Passengers, &bookingDetail.BookingStatus, &bookingDetail.PaymentStatus,
			&bookingDetail.Method, &bookingDetail.AmountPaid, &bookingDetail.PaymentDate, &bookingDetail.PassengerName,
			&bookingDetail.Age, &bookingDetail.Gender, &bookingDetail.Contact)
		if err != nil {
			return nil, err

		}
		fmt.Println(bookingDetail)
		bookingDetails = append(bookingDetails, bookingDetail)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return bookingDetails, nil
}
