package main

import (
	"net/http"

	bus "github.com/chandanaavadhani/BusService/handlers/bus"
	trips "github.com/chandanaavadhani/BusService/handlers/trips"
)

func main() {
	http.HandleFunc("/bus", bus.CreateBus)
	http.HandleFunc("/busses", bus.GetAllBusses)
	http.HandleFunc("/trip", trips.Trip)
	http.HandleFunc("/trip/", trips.GetTripByID)
	http.HandleFunc("/bus/", bus.DeleteOrGetBus)
	http.ListenAndServe(":8000", nil)
}
