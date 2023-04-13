package utils

import (
	"encoding/json"
	models "main/BusService/models"
	"net/http"
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
