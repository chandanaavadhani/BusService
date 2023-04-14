package main

import (
	"fmt"
	"net/http"

	handlers1 "github.com/chandanaavadhani/BusService/handlers/operators"
	handlers "github.com/chandanaavadhani/BusService/handlers/users"
	repository "github.com/chandanaavadhani/BusService/repository"
)

func main() {
	db, err := repository.DBConnection()

	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	defer db.Close()
	//handling routes
	http.HandleFunc("/operators", handlers1.CreateOperator)
	http.HandleFunc("/v1/signup", handlers.Signup)
	http.HandleFunc("/v1/login", handlers.Login)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
