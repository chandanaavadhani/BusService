package repository

import (
	models "github.com/chandanaavadhani/BusService/models"
)

func InsertReview(ratingid, userid, review models.Review) error {

	// DB connection
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Insert review into Database
	stmt, err := db.Prepare("INSERT INTO reviews (ratingid,userid,busid,comment,rating) VALUES (?, ?,?,?,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ratingid, userid, review.busid, review.comment, review.rating)
	if err != nil {
		return err
	}

	return nil
}
