package models

import "github.com/dgrijalva/jwt-go"

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Bus struct {
	BusId      string `json:"bus_id"`
	OperatorId string `json:"opertor_id"`
	Contact    string `json:"contact"`
	Capacity   int    `json:"capacity"`
	BusType    string `json:"bus_type"`
	BusNumber  string `json:"bus_number"`
}

type AddBusRequest struct {
	Contact   string `json:"contact"`
	Capacity  int    `json:"capacity"`
	BusType   string `json:"bus_type"`
	BusNumber string `json:"bus_number"`
}

type Trip struct {
	TripId    string `json:"trip_id"`
	BusId     string `json:"bus_id"`
	RouteId   string `json:"route_id"`
	Departure string `json:"departure"`
	Arrival   string `json:"arrival"`
	Cost      string `json:"cost"`
}

type TokenClaim struct {
	Username   string `json:"username"`
	IsOperator string `json:"isoperator"`
	jwt.StandardClaims
}
