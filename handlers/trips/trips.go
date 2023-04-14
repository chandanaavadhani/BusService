package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	"github.com/chandanaavadhani/BusService/validators"
)

func AddTrip(w http.ResponseWriter, r *http.Request) {
	var data []string
	if r.Method != http.MethodPost {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
		return
	}
	//Establish DB connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
		return
	}
	defer db.Close()

	//Validate operator token and get details
	isoperator := true
	if !isoperator {
		utils.BuildResponse(w, http.StatusBadRequest, "preveliges are not given for this user", data)
		return
	}

	//Process the request body to access the data
	body, err := utils.ExtractReqBody(r)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		return
	}

	var tripRequest models.AddTripRequest
	err = json.Unmarshal(body, &tripRequest)
	if err != nil {
		utils.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
		return
	}

	//validate request
	err = validators.ValidateAddTripRequest(tripRequest, db)
	if err != nil {
		fmt.Println("Error is : ", err.Error())
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		return
	}

	//get capacity using bus id
	capacity, err := repository.GetCapacityOfBus(tripRequest.BusId, db)
	if err != nil {
		fmt.Println("Error here is : ", err.Error())
		utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		return
	}

	err = repository.AddTriptoDB(tripRequest, capacity, db)
	if err != nil {
		fmt.Println("Error here here is : ", err.Error())
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
		return
	}

	//succesful response
	utils.BuildResponse(w, http.StatusCreated, "", data)
}
