package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os/exec"

	models "github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
	validators "github.com/chandanaavadhani/BusService/validators"
)

func generateRatingId() string {
	ratingID, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(ratingID)
}

func AddReviews(w http.ResponseWriter, r *http.Request) {

	//Validate method
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}

	//Generate rating ID
	ratingId := generateRatingId()
	fmt.Println(ratingID)

	// get user id
	userId := "83064af3-bb81-4514-a6d4-afba340825cd"

	//extract reviews from the request
	var review models.Review

	err := json.NewDecoder(r.Body).Decode(&review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//validate request
	status, err := validators.ValidateReviews(review)
	if err != nil {
		http.Error(w, err.Error(), status)
		return
	}

	//insert reviews into db
	err = repository.InsertReview(ratingId, userId, review)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//send response

}
