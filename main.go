package main

import (
	"fmt"
	"net/http"

	handlers "github.com/chandanaavadhani/BusService/handlers/bookings"
	repository "github.com/chandanaavadhani/BusService/repository"
	"github.com/gorilla/mux"
)

func main() {
	db, err := repository.DBConnection()
	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	defer db.Close()

	router := mux.NewRouter()

	//handling routes
	router.HandleFunc("/bookinghistory/{userId}", handlers.BookingHistory).Methods("GET")
	router.HandleFunc("/cancelbooking/{bookingId}", handlers.CancelBooking).Methods("PUT")
	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
