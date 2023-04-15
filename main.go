package main

import (
	"fmt"
	"net/http"

	handlers2 "github.com/chandanaavadhani/BusService/handlers/bookings"
	handlers1 "github.com/chandanaavadhani/BusService/handlers/reviews"
	handlers "github.com/chandanaavadhani/BusService/handlers/trips"
)

func main() {

	//handling routes
	// http.HandleFunc("/v1/operators", handlers.CreateOperator)
	http.HandleFunc("/v1/addreview", handlers1.AddReviews)
	http.HandleFunc("/v1/trips", handlers.GetTrips)
	http.HandleFunc("/v1/trips/", handlers.GetTripDetails)
	http.HandleFunc("/v1/addbookings", handlers2.AddBookings)
	http.HandleFunc("/v1/coupon/", handlers2.ValidateCoupon)
	http.HandleFunc("/v1/bookings/", handlers2.GetBookingDetails)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
