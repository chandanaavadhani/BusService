package main

import (
	"fmt"
	"net/http"

	coupons "github.com/chandanaavadhani/BusService/handlers/coupons"
	operators "github.com/chandanaavadhani/BusService/handlers/operators"
	"github.com/chandanaavadhani/BusService/repository"
)

func main() {
	db, err := repository.DBConnection()
	defer db.Close()

	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	//handling routes
	http.HandleFunc("/operators", operators.CreateOperator)
	http.HandleFunc("/v1/coupons", coupons.CreateCoupon)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
