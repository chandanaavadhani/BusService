package handlers

import (
	"encoding/json"
	"net/http"

	types "github.com/chandanaavadhani/BusService/models"
)

func Signup(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decode the JSON request body into a LoginRequest struct.
	var user types.Login
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

}
