package models

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Review struct {
	BusId   string `json:"busid"`
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
}

type GetTripsRequest struct {
	Source      string `json:"source"`
	Destination string `json:"destination"`
	Date        string `json:"date"`
}

type Trips struct {
	TripId    string `json:"tripID"`
	BusId     string `json:"busID"`
	RouteId   string `json:"RouteID"`
	Departure string `josn:"departure"`
	Arrival   string `json:"arrival"`
	Capacity  string `json:"capacity"`
	Cost      string `json:"cost"`
	BusStatus string `json:"busStatus"`
}

type TripDetails struct {
	TripId        string `json:"tripID"`
	BusId         string `json:"busID"`
	RouteId       string `json:"RouteID"`
	Departure     string `josn:"departure"`
	Arrival       string `json:"arrival"`
	Capacity      string `json:"capacity"`
	Cost          string `json:"cost"`
	BusStatus     string `json:"busStatus"`
	DriverContact string `json:"driverContact"`
	BusCapacity   string `json:"busCapacity"`
	BusNumber     string `json:"busNumber"`
	BusType       string `json:"busType"`
	OperatorId    string `json:"operatorID"`
	Source        string `json:"source"`
	Destination   string `json:"destination"`
	Distance      string `json:"distance"`
}

type Bookings struct {
	TripId        string  `json:"tripID"`
	CouponCode    string  `json:"couponCode"`
	Passengers    string  `json:"passengers"`
	BookingStatus string  `json:"bookingStatus"`
	PaymentStatus string  `json:"paymentStatus"`
	Method        string  `json:"method"`
	AmountPaid    float64 `json:"amountPaid"`
}
