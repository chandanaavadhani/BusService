package utils

import (
	"regexp"
)

func IsEmailValid(Email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, err := regexp.MatchString(pattern, Email)
	if err != nil {
		return false
	}
	return match
}

func IsPasswordValid(Password string) bool {
	// pattern := `^(?=.[A-Z])(?=.[a-z])(?=.\d)(?=.[@$!%?&#])[A-Za-z\d@#$!%?&]{8,}$`
	// match, err := regexp.MatchString(pattern, password)
	// if err != nil {
	// 	return false
	// }
	// return match
	return len(Password) >= 8
}

func IsPhoneNumberValid(Phone string) bool {

	// 	123-456-7890
	// (123) 456-7890
	// 123 456 7890
	// 123.456.7890
	// +1 (123) 456-7890

	pattern := `^(\+\d{1,2}\s)?\(?\d{3}\)?[\s.-]?\d{3}[\s.-]?\d{4}$`
	match, err := regexp.MatchString(pattern, Phone)
	if err != nil {
		return false
	}
	return match
}
