package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	validators "github.com/chandanaavadhani/BusService/validators"
	uuid "github.com/satori/go.uuid"
)

func GenerateBookingId() string {
	ratingId := uuid.NewV4()
	return ratingId.String()
}
func GeneratePaymentId() string {
	ratingId := uuid.NewV4()
	return ratingId.String()
}

func AddBookings(w http.ResponseWriter, r *http.Request) {

	//validate method
	if r.Method != "POST" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Methos not Allowed", nil)
	}

	//Generate Payment Id
	paymentId := GeneratePaymentId()
	fmt.Println(paymentId)

	//Generate booking ID
	bookingId := GenerateBookingId()
	fmt.Println(bookingId)

	//Get userID
	userId := "83064af3-bb81-4514-a6d4-afba340825cd"

	//extract body from the request
	var booking models.Bookings
	err := json.NewDecoder(r.Body).Decode(&booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate request
	status, err := validators.ValidateBookingRequest(paymentId, booking)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	//insert payment details into db
	err = repository.InsertPaymentDetails(paymentId, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//insert bookings details into db if payment is successful
	err = repository.InsertBookingDetails(bookingId, userId, paymentId, booking)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send response
	utils.BuildResponse(w, http.StatusCreated, "Bookings added successfully", nil)
}
