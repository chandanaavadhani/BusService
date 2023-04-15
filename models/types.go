package models

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Operator struct {
	OperatorName string `json:"operatorName"`
	Contact      string `json:"contact"`
	Address      string `json:"address"`
	Email        string `json:"email"`
	Password     string `json:"password"`
}

type UpdateOperator struct {
	OperatorName string `json:"operatorName"`
	Contact      string `json:"contact"`
	Address      string `json:"address"`
}

type GetOperatorResponse struct {
	OperatorID   string               `json:"operatorID"`
	OperatorName string               `json:"operatorName"`
	Contact      string               `json:"contact"`
	Address      string               `json:"address"`
	Email        string               `json:"email"`
	Password     string               `json:"password"`
	Buses        []BusDetailsOperator `json:"buses"`
}

type BusDetailsOperator struct {
	BusId     string                   `json:"busID"`
	Contact   string                   `json:"contact"`
	Capacity  int64                    `json:"capacity"`
	BusType   string                   `json:"busType"`
	BusNumber string                   `json:"busNumber"`
	Reviews   []ReviewResponseOperator `json:"reviews"`
}
type ReviewResponseOperator struct {
	RatingID string `json:"ratingID"`
	UserID   string `json:"userID"`
	Comment  string `json:"comment"`
	Rating   int64  `json:"rating"`
}
