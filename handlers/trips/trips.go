package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	models "github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
	utils "github.com/chandanaavadhani/BusService/utils"
	validators "github.com/chandanaavadhani/BusService/validators"
)

func GetTrips(w http.ResponseWriter, r *http.Request) {

	//validate method
	if r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
		return
	}

	//extract body from the request
	var route models.GetTripsRequest
	err := json.NewDecoder(r.Body).Decode(&route)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	//Validate request
	status, err := validators.ValidateRouteDetails(route)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}
	// Get trips from the database

	trips, err := repository.GetTripsBasedOnRoute(route)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send response
	utils.BuildResponse(w, http.StatusOK, "Getting Trips Succesfully", trips)
}

func GetTripDetails(w http.ResponseWriter, r *http.Request) {

	//Validate Method
	if r.Method != "GET" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Methos not Allowed", nil)
		return
	}
	//Get trip id
	tripId := strings.Split(r.URL.Path, "/trips/")[1] + "\n"
	fmt.Println(tripId)

	//Validate Trip id
	status, err := validators.ValidateTripId(tripId)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	//Get trips from the database
	tripDetails, err := repository.GetTripDetails(tripId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send response
	utils.BuildResponse(w, http.StatusOK, "Get trip details successfully", tripDetails)

}
