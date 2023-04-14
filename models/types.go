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
