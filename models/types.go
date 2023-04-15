package models

import "github.com/dgrijalva/jwt-go"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Reviews struct {
	BusID    string `json:"bus_id"`
	RatingID string `json:"rating_id"`
	UserID   string `json:"user_id"`
	Comment  string `json:"comment"`
	Rating   string `json:"rating"`
}
type BusSchema struct {
	BusId      string `json:"bus_id"`
	OperatorId string `json:"operator_id"`
	Contact    string `json:"contact"`
	Capacity   int    `json:"capacity"`
	BusType    string `json:"bus_type"`
	BusNumber  string `json:"bus_number"`
}

type BusDetails struct {
	BusId      string                       `json:"bus_id"`
	OperatorId string                       `json:"operator_id"`
	Contact    string                       `json:"contact"`
	Capacity   int                          `json:"capacity"`
	BusType    string                       `json:"bus_type"`
	BusNumber  string                       `json:"bus_number"`
	Reviews    []ReviewResponseFromDBForBus `json:"reviews"`
}
type AddBusRequest struct {
	Contact   string `json:"contact"`
	Capacity  int    `json:"capacity"`
	BusType   string `json:"bus_type"`
	BusNumber string `json:"bus_number"`
}

type UpdateTripRequest struct {
	TripId    string  `json:"trip_id"`
	BusId     string  `json:"bus_id"`
	RouteId   string  `json:"route_id"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
	Cost      float64 `json:"cost"`
	BusStatus string  `json:"bus_status"`
}

type TokenClaim struct {
	Username   string `json:"username"`
	IsOperator string `json:"isoperator"`
	jwt.StandardClaims
}

type AddTripRequest struct {
	BusId     string  `json:"bus_id"`
	RouteId   string  `json:"route_id"`
	Departure string  `json:"departure"`
	Arrival   string  `json:"arrival"`
	Cost      float64 `json:"cost"`
	BusStatus string  `json:"bus_status"`
}

type ReviewResponseFromDBForBus struct {
	RatingID string `json:"rating_id"`
	UserID   string `json:"user_id"`
	Comment  string `json:"comment"`
	Rating   int64  `json:"rating"`
}

type TripDetails struct {
	TripId        string  `json:"tripID"`
	BusId         string  `json:"busID"`
	RouteId       string  `json:"RouteID"`
	Departure     string  `josn:"departure"`
	Arrival       string  `json:"arrival"`
	Capacity      string  `json:"capacity"`
	Cost          float64 `json:"cost"`
	BusStatus     string  `json:"busStatus"`
	DriverContact string  `json:"driverContact"`
	BusCapacity   string  `json:"busCapacity"`
	BusNumber     string  `json:"busNumber"`
	BusType       string  `json:"busType"`
	OperatorId    string  `json:"operatorID"`
	Source        string  `json:"source"`
	Destination   string  `json:"destination"`
	Distance      string  `json:"distance"`
}
