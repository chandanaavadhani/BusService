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

func CreateBus(w http.ResponseWriter, r *http.Request) {
	var data []string
	//if it is post request then add bus
	if r.Method == http.MethodPost {
		fmt.Println("This is called")
		//Establish DB connection
		db, err := repository.DBConnection()
		if err != nil {
			utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
			return
		}
		defer db.Close()
		//get info from token for now it is skipped
		id := "12345"
		isoperator := true
		if !isoperator {
			utils.BuildResponse(w, http.StatusBadRequest, "Invalid User", data)
			return
		}
		//Process the request body to access the data
		body, err := utils.ExtractReqBody(r)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
			return
		}

		var bus models.AddBusRequest
		err = json.Unmarshal(body, &bus)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, "Error parsing request body", data)
			return
		}

		//validate request
		err = validators.ValidateAddBusRequest(bus)
		if err != nil {
			utils.BuildResponse(w, http.StatusBadRequest, err.Error(), data)
		}

		//Check if bus already exists

		//Add bus to DB
		err = repository.AddBusToDB(bus, id, db)
		if err != nil {
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
			return
		}

		utils.BuildResponse(w, http.StatusCreated, "succesfully added new bus", data)
		return
	} else {
		utils.BuildResponse(w, http.StatusBadRequest, "not a POST request", data)
		return
	}
}

func GetAllBusses(w http.ResponseWriter, r *http.Request) {
	var data []string
	//Check for method
	if r.Method != http.MethodGet {
		utils.BuildResponse(w, http.StatusBadRequest, "not a GET request", data)
		return
	}

	//Establish DB connection
	db, err := repository.DBConnection()
	if err != nil {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "not a POST request", data)
		return
	}
	defer db.Close()

	//get info from token for now it is skipped
	id := "1234"
	isoperator := false
	isadmin := true
	if !isoperator {
		if !isadmin {
			utils.BuildResponse(w, http.StatusBadRequest, "previliges has been restricted to admin and operator", data)
			return
		}

		//Admin gets all Buses from DB
		allBusses, err := repository.GetAllBusses(db)
		if err != nil {
			fmt.Println("Error is : ", err.Error())
			utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
			return
		}
		//Send Response
		utils.BuildResponse(w, http.StatusOK, "", allBusses)
		return
	}
	//Get all buses from DB
	allBusses, err := repository.GetAllBussesOfOperator(id, db)
	if err != nil {
		fmt.Println("Error is : ", err.Error())
		utils.BuildResponse(w, http.StatusInternalServerError, err.Error(), data)
		return
	}
	//Send Response
	utils.BuildResponse(w, http.StatusOK, "", allBusses)
}
