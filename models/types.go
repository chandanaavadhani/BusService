package models

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type Signup struct {
	UserId    string `json:"userID"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Gender    string `json:"gender"`
	Age       string `json:"age"`
	Phone     string `json:"phone"`
}
