package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	uuid "github.com/satori/go.uuid"

	models "github.com/chandanaavadhani/BusService/models"
	repository "github.com/chandanaavadhani/BusService/repository"
	utils "github.com/chandanaavadhani/BusService/utils"
	validators "github.com/chandanaavadhani/BusService/validators"
)

func generateRatingId() string {
	// newUUID, err := exec.Command("uuidgen").Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Generated UUID:")
	// fmt.Printf("%s", newUUID)
	// return string(newUUID)
	ratingId := uuid.NewV4()
	return ratingId.String()
}

func AddReviews(w http.ResponseWriter, r *http.Request) {

	//Validate method
	if r.Method != "POST" {
		utils.BuildResponse(w, http.StatusMethodNotAllowed, "Method not allowed", nil)
	}

	//Generate rating ID
	ratingId := generateRatingId()
	fmt.Println(ratingId)

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
	utils.BuildResponse(w, http.StatusCreated, "Review added successfully", nil)
}
