package main

import (
	"fmt"
	"net/http"

	coupons "github.com/chandanaavadhani/BusService/handlers/coupons"
	"github.com/chandanaavadhani/BusService/repository"
)

func main() {
	db, err := repository.DBConnection()
	if err != nil {
		fmt.Println("Error in Connecting the DB : ", err)
	}
	defer db.Close()

	//handling routes
	http.HandleFunc("/v1/coupons/add", coupons.CreateCoupon)
	http.HandleFunc("/v1/coupons/update", coupons.UpdateCoupon)
	http.HandleFunc("/v1/coupons/delete", coupons.DeleteOrGetCoupon)
	http.HandleFunc("/v1/coupons/coupon", coupons.DeleteOrGetCoupon)
	http.HandleFunc("/v1/coupons/list", coupons.GetAllCoupons)

	//hosting the server
	fmt.Println("Local host is servered at port 8000")
	err = http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}
}
