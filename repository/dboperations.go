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

func UpdateCoupon(promo models.UpdateCoupon) error {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//Execute the query
	_, err = db.Query("UPDATE discounts SET couponAmount = ? WHERE couponCode = ?", promo.NewCouponAmount, promo.CouponCode)
	return err
}

func IsUpdatedCouponExists(promo models.UpdateCoupon) (int, string, error) {
	var count int
	var code string
	db, err := DBConnection()
	if err != nil {
		return 0, "", err
	}
	defer db.Close()

	//Execute the query
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

func DeleteCoupon(coupon string) error {
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM discounts where couponCode = ?", coupon)
	return err
}

func IsCouponExists(coupon string) (int, error) {
	var count int
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query, err := db.Query("select count(*) from discounts where couponCode = ?", coupon)
	if err != nil {
		return 0, err
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&count); err != nil {
			return 0, err
		}
	}
	return count, nil
}

func GetCoupon(coupon string) (models.Coupon, error) {
	var promo models.Coupon
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	query, err := db.Query("SELECT * FROM discounts where couponCode = ?", coupon)
	if err != nil {
		return promo, err
	}
	defer query.Close()

	for query.Next() {
		if err := query.Scan(&promo.CouponCode, &promo.CouponAmount); err != nil {
			return promo, err
		}
	}
	return promo, nil
}

func GetAllCoupons() ([]models.Coupon, error) {
	var coupons []models.Coupon
	db, err := DBConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	query, err := db.Query("SELECT * FROM discounts")
	if err != nil {
		return coupons, err
	}
	defer query.Close()

	for query.Next() {
		var promo models.Coupon
		err := query.Scan(&promo.CouponCode, &promo.CouponAmount)
		if err != nil {
			return coupons, err
		}
		coupons = append(coupons, promo)
	}

	return coupons, nil
}
