package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
	"github.com/chandanaavadhani/BusService/utils"
	"github.com/chandanaavadhani/BusService/validators"
)

func Trip(w http.ResponseWriter, r *http.Request) {
	var data []string

	//Establish DB connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
		return
	}
	defer db.Close()

	if r.Method == http.MethodPost {
		//Add Trip To DB

		//Validate operator token and get details
		isoperator := true
		if !isoperator {
			utils.BuildResponse(w, http.StatusBadRequest, "preveliges were not given for this user", data)
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
		return
	} else if r.Method == http.MethodPut {
		//Update Trip

		//Validate operator token and get details
		isoperator := true
		if !isoperator {
			utils.BuildResponse(w, http.StatusBadRequest, "preveliges were not given for this user", data)
			return
		}

		//Process the request body to access the data
		body, err := utils.ExtractReqBody(r)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		var tripRequest models.UpdateTripRequest
		err = json.Unmarshal(body, &tripRequest)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
			return
		}

		//Validate request
		err = validators.ValidateUpdateTripRequest(tripRequest, db)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		//Update Trip Details
		err = repository.UpdateTripDetails(tripRequest, db)
		if err != nil {
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
			return
		}

		//send success response
		utils.BuildResponse(w, http.StatusOK, "updated succesfully", data)
		return

	} else if r.Method == http.MethodGet {
		//Validate operator token and get details
		id := "12345"
		isoperator := false
		isadmin := true
		if !isadmin {
			if !isoperator {
				utils.BuildResponse(w, http.StatusBadRequest, "preveliges were not given for this user", data)
				return
			}

			//Operator Logic
			trips, err := repository.GetAllTripsDetailsOfOperator(id, db)
			if err != nil {
				utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
				return
			}

			//return success response
			utils.BuildResponse(w, http.StatusOK, "", trips)
			return
		}

		//Admin Logic
		trips, err := repository.GetAllTripsDetails(db)
		if err != nil {
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
			return
		}

		//return success response
		utils.BuildResponse(w, http.StatusOK, "", trips)
		return
	} else {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a valid request", data)
	}

}

func GetTripByID(w http.ResponseWriter, r *http.Request) {
	var data []string

	//Establish DB connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
		return
	}
	defer db.Close()

	if r.Method != http.MethodGet {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a valid request", data)
		return
	}

	//Validate operator token and get details
	isoperator := true
	if !isoperator {
		utils.BuildResponse(w, http.StatusBadRequest, "preveliges were not given for this user", data)
		return
	}

	//Get busId from req URL
	path := r.URL.Path
	segments := strings.Split(path, "/")
	tripID := segments[2] + "\n"

	trips, err := repository.GetTripDetails(tripID, db)
	if err != nil {
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
		return
	}

	//return success response
	utils.BuildResponse(w, http.StatusOK, "", trips[0])
}
