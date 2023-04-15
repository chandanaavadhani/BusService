package main

import (
	"fmt"
	"net/http"

	Operators "github.com/chandanaavadhani/BusService/handlers/Operators"
	repository "github.com/chandanaavadhani/BusService/repository"
)

func main() {
	db, err := repository.DBConnection()
	defer db.Close()

	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	//handling routes
	http.HandleFunc("/v1/operator/create", Operators.CreateOperator)
	http.HandleFunc("/v1/operator/update/", Operators.UpdateOperator)
	http.HandleFunc("/v1/operator/delete/", Operators.DeleteOperator)
	http.HandleFunc("/v1/operator/get/", Operators.GetOperator)
	http.HandleFunc("/v1/operators", Operators.GetAllOperators)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
