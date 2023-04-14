package repository

import (
	"github.com/chandanaavadhani/BusService/models"
	_ "github.com/go-sql-driver/mysql"
)

func IsCouponCodeExists(promo models.Coupon) (int, string, error) {
	var count int
	var code string
	db, err := DBConnection()
	if err != nil {
		return 0, "", err
	}
	defer db.Close()
	query, err := db.Query("SELECT count(*), couponCode FROM discounts WHERE couponCode = ? GROUP BY couponCode", promo.CouponCode)
	if err != nil {
		return 0, "", err
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&count, &code); err != nil {
			return 0, "", err
		}
	}
	return count, code, nil
}

func InsertCoupon(promo models.Coupon) error {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Execute the query
	query, err := db.Prepare(`INSERT INTO discounts VALUES(?,?)`)
	if err != nil {
		return err
	}
	defer query.Close()

	_, err = query.Exec(promo.CouponCode, promo.CouponAmount)
	return err
}
