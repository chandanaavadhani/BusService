package utils

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(UserId string, Isadmin bool, Isoperator bool) (string, error) {

	claims := jwt.MapClaims{
		"Id":         UserId,
		"IsAdmin":    Isadmin,
		"IsOperator": Isoperator,
		"exp":        time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey())
}
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// get the token
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}
		fmt.Println(tokenString)
		// parse the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return SecretKey(), nil
		})
		fmt.Println(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		// validate the token and return the correponding error or continue with the handler
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims)
			Id, ok := claims["Id"].([]byte)
			if !ok {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			IsAdmin, ok := claims["IsAdmin"].([]byte)
			if !ok {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			IsOperator, ok := claims["IsOperator"].([]byte)
			if !ok {
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}
			//assigning the email to the context of request
			ctx := context.WithValue(r.Context(), "Id", Id)
			ctx = context.WithValue(r.Context(), "IsAdmin", IsAdmin)
			ctx = context.WithValue(r.Context(), "IsOperator", IsOperator)

			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
	})
}
