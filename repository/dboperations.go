package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	types "github.com/chandanaavadhani/BusService/models"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

func InsertUser(user types.Signup) (string, error) {
	db, err := DBConnection()
	if err != nil {
		return "", err
	}
	stmt, err := db.Prepare("INSERT INTO `bus service`.`user`(userID, firstName,lastName, email,password,gender,age,phone) VALUES (?,?,?,?,?,?,?,?)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	userid := uuid.NewV4()

	fmt.Println("ID Generated:")
	fmt.Printf("%s", userid)

	// Hash the user's password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	_, err = stmt.Exec(userid, user.Firstname, user.Lastname, user.Email, hashedPassword, user.Gender, user.Age, user.Phone)
	if err != nil {
		return "", err
	}
	return userid.String(), nil
}

func LoginUser(user types.Login) (string, bool, bool, int, error) {
	IsAdmin := false
	IsOperator := false

	db, err := DBConnection()
	if err != nil {
		return "", IsAdmin, IsOperator, http.StatusInternalServerError, err
	}

	parts := strings.Split(user.Email, "@")
	// Get the hashed password from the database

	//checking admin email or operator email
	if parts[1] == "admin.busservice.com" || parts[1] == "operator.busservice.com" {
		if parts[1] == "admin.busservice.com" {
			IsAdmin = true
		} else {
			IsOperator = true
		}
		//var count int
		var operatorid string
		var hashedPassword string
		err = db.QueryRow("SELECT operatorID,password FROM `bus service`.`operator` WHERE email=?", user.Email).Scan(&operatorid, &hashedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", IsAdmin, IsOperator, http.StatusBadRequest, errors.New("operator not exists")
			}
			return "", IsAdmin, IsOperator, http.StatusInternalServerError, err
		}

		// Compare the hashed password with the user's password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return "", IsAdmin, IsOperator, http.StatusBadRequest, errors.New("incorrect password")
			}
			return "", IsAdmin, IsOperator, http.StatusInternalServerError, err
		}
		return operatorid, IsAdmin, IsOperator, http.StatusOK, nil
	} else {
		// var count int
		var userid string
		var hashedPassword string
		err = db.QueryRow("SELECT userID,password FROM `bus service`.`user` WHERE email=?", user.Email).Scan(&userid, &hashedPassword)
		if err != nil {
			if err == sql.ErrNoRows {
				return "", IsAdmin, IsOperator, http.StatusBadRequest, errors.New("user not exist")
			}
			return "", IsAdmin, IsOperator, http.StatusInternalServerError, err
		}

		// Compare the hashed password with the user's password
		err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				return "", IsAdmin, IsOperator, http.StatusBadRequest, errors.New("incorrect password")
			}
			return "", IsAdmin, IsOperator, http.StatusInternalServerError, err
		}
		return userid, IsAdmin, IsOperator, http.StatusOK, nil
	}

}
