package handlers

import (
	"encoding/json"
	"main/BusService/models"
	utils "main/BusService/utils"
	"main/BusService/validators"
	"main/services"
	"net/http"
)

func CreateBus(w http.ResponseWriter, r *http.Request) {
	//if it is post request then add bus
	if r.Method == http.MethodPost {
		var data []string
		//get info from token for now it is skipped
		isoperator := true
		if !isoperator {
			utils.BuildResponse(w, http.StatusBadRequest, "Invalid User", data)
			return
		}
		//Process the request body to access the data
		body, err := services.ExtractReqBody(r)
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
	}
}
