package utils

import (
	"fmt"
	"time"

	models "main/BusService/models"

	"github.com/dgrijalva/jwt-go"
)

// get secret key
func getSecretKey() []byte {
	return []byte("userauthtestproject")
}

// generate token with username encoded
func GenerateToken(userid string, isoperator bool) string {
	secretKey := getSecretKey()
	//Configure header and payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userid":     userid,
		"isoperator": isoperator,
		"exp":        time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "error"
	}
	return tokenString
}

// Get data from the token
func GetDataFromToken(tokenString string) (*models.TokenClaim, error) {
	claims := &models.TokenClaim{}
	secretKey := getSecretKey()

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return &models.TokenClaim{}, fmt.Errorf(err.Error())
	}

	if token.Valid {
		return claims, nil
	}
	return &models.TokenClaim{}, fmt.Errorf("Invalid Token")
}
