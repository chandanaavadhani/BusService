package handlers

import (
	"net/http"
	"strconv"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	"github.com/gorilla/mux"
)

func BookingHistory(w http.ResponseWriter, r *http.Request) {
	// Get the user ID from the URL parameters

	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userId"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	//Get all Booking details by ID
	rows, _ := repository.GetBookingDetailsById(userID, w)

	// Loop through the rows and build a slice of bookings
	var bookings []models.BookingDetails
	for rows.Next() {
		var booking models.BookingDetails
		err := rows.Scan(&booking.BookingId, &booking.UserId, &booking.TripId, &booking.PaymentId, &booking.Passengers, &booking.BookingStatus, &booking.PaymentStatus, &booking.Contact)
		if err != nil {
			http.Error(w, "Failed to retrieve booking history", http.StatusInternalServerError)
			return
		}
		bookings = append(bookings, booking)
	}

	// Check for any errors that may have occurred during the iteration
	if err := rows.Err(); err != nil {
		http.Error(w, "Failed to retrieve booking history", http.StatusInternalServerError)
		return
	}

	utils.BuildResponse(w, http.StatusOK, "Bookdetails By Id Succesfull", bookings)
}

// CancellingBooking
func CancelBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookingIDStr, ok := vars["bookingId"]
	if !ok {
		http.Error(w, "Missing booking ID", http.StatusBadRequest)
		return
	}
	bookingID, err := strconv.Atoi(bookingIDStr)
	if err != nil {
		http.Error(w, "Invalid booking ID", http.StatusBadRequest)
		return
	}

	// Cancel the booking in the database
	rowsAffected, err := repository.GetDetailsAfterDeleting(bookingID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if rowsAffected == 0 {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Return a success response
	utils.BuildResponse(nil, http.StatusNoContent, "Deletion Done", nil)
}
