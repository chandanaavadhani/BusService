package utils

import (
	"encoding/json"
	"net/http"

	models "github.com/chandanaavadhani/BusService/models"
)

// Build a Response
func BuildResponse(w http.ResponseWriter, code int, message string, data interface{}) {
	var buildRequestResponse models.Response
	buildRequestResponse.Message = message
	buildRequestResponse.Data = data

	//write response
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(buildRequestResponse)
}
