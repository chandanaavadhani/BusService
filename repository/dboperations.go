package repository

import (
	"database/sql"
	"net/http"
)

func GetBookingDetailsById(userID int, w http.ResponseWriter) (*sql.Rows, error) {

	// Connect to the database
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}

	// Query the database for the user's booking history

	rows, err := db.Query("SELECT * FROM bookings WHERE user_id=?", userID)
	if err != nil {
		http.Error(w, "Failed to retrieve booking history", http.StatusInternalServerError)
		return nil, err
	}
	defer rows.Close()
	defer db.Close()

	return rows, nil
}

func GetDetailsAfterDeleting(bookingID int) (int64, error) {

	// Connect to the database
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	// Execute the SQL query to cancel the booking
	result, err := db.Exec("DELETE FROM bookings WHERE booking_id=?", bookingID)
	if err != nil {
		return 0, err
	}

	// Return the number of rows affected
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
