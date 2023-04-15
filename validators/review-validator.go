package validators

import (
	"errors"
	"net/http"

	models "github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/repository"
)

func ValidateReviews(userId string, review models.Review) (int, error) {

	if review.BusId == "" {
		return http.StatusBadRequest, errors.New("Bus ID missing")
	}
	if review.Comment == "" {
		return http.StatusBadRequest, errors.New("Comment missing")
	}
	if review.Rating == 0 {
		return http.StatusBadRequest, errors.New("Rating missing")
	}
	if repository.CheckIfReviewAdded(review.BusId, userId) != false {
		return http.StatusBadRequest, errors.New("Review added already")
	}
	return 200, nil
}
