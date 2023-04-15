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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(buildRequestResponse)
	// response := models.Response{
	// 	Message: message,
	// 	Data:    data,
	// }
	// jsonResponse, err := json.Marshal(response)
	// if err != nil {
	// 	BuildResponse(w, http.StatusInternalServerError, err.Error(), nil)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(code)
	// w.Write(jsonResponse)
}
