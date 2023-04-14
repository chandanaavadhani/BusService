package models

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Coupon struct {
	CouponCode   string  `json:"couponcode"`
	CouponAmount float64 `json:"couponamount"`
}
