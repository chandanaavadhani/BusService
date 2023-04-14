package repository

import (
	models "github.com/chandanaavadhani/BusService/models"
)

func InsertReview(ratingid string, userid string, review models.Review) error {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Insert review into Database
	stmt, err := db.Prepare("INSERT INTO reviews (ratingID, userID, busID,comment,rating) VALUES (?,?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ratingid, userid, review.BusId, review.Comment, review.Rating)
	if err != nil {
		return err
	}

	return nil
}
