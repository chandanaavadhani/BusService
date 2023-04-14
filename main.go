package main

import (
	"fmt"
	"net/http"

	handlers "github.com/chandanaavadhani/BusService/handlers/reviews"
	repository "github.com/chandanaavadhani/BusService/repository"
)

func main() {
	db, err := repository.DBConnection()
	defer db.Close()

	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	//handling routes
	// http.HandleFunc("/v1/operators", handlers.CreateOperator)
	http.HandleFunc("/v1/addreview", handlers.AddReviews)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
