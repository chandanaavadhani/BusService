package repository

import (
	"database/sql"
	"errors"
	"net/http"

	models "github.com/chandanaavadhani/BusService/models"
	"github.com/chandanaavadhani/BusService/utils"
	"golang.org/x/crypto/bcrypt"
)

func InsertOperator(db *sql.DB, newOperator models.Operator) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newOperator.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	stmt, err := db.Prepare(`INSERT INTO operator(operatorID,name,email,password,phone,address) Values (?,?,?,?,?,?)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	Id := utils.UUID()
	_, err = stmt.Exec(Id, newOperator.OperatorName, newOperator.Email, string(hashedPassword), newOperator.Contact, newOperator.Address)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil

}

func UpdateOperatorDetails(db *sql.DB, updateOperator models.UpdateOperator, id string) (int, error) {
	_, err := db.Query(`UPDATE operator SET name = ?, phone = ?, address = ? WHERE operatorID = ?`, updateOperator.OperatorName, updateOperator.Contact, updateOperator.Address, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusOK, nil
}

func DeleteOperatorDetails(db *sql.DB, id string) (int, error) {
	_, err := db.Query(`DELETE FROM operator WHERE operatorID = ?`, id)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return 204, nil
}

func GetOperatorDetails(db *sql.DB, id string) (int, error, models.GetOperatorResponse) {
	var operator models.GetOperatorResponse
	rows, err := db.Query(`
	SELECT operator.*,
	buses.busID,buses.driverContact,buses.capacity,buses.busType,buses.busNumber,
	reviews.ratingID,reviews.userID, reviews.comment,reviews.rating
	FROM operator
	LEFT JOIN buses
	ON operator.operatorID = buses.operatorID
	LEFT JOIN reviews
	ON buses.busID = reviews.busID
	WHERE buses.operatorID = ?
	`, id)
	if err != nil {
		return http.StatusInternalServerError, err, operator
	}
	defer rows.Close()

	// Iterate through the rows

	for rows.Next() {

		var busID sql.NullString
		var driverContact sql.NullString
		var capacity sql.NullInt64
		var busType sql.NullString
		var busNumber sql.NullString
		var ratingID sql.NullString
		var userID sql.NullString
		var comment sql.NullString
		var rating sql.NullInt64

		err := rows.Scan(&operator.OperatorID, &operator.OperatorName, &operator.Contact,
			&operator.Address, &operator.Email, &operator.Password, &busID, &driverContact,
			&capacity, &busType, &busNumber, &ratingID, &userID, &comment, &rating)
		if err != nil {
			return http.StatusInternalServerError, err, operator
		}

		var busDetails models.BusDetailsOperator
		busDetails.BusId = busID.String
		busDetails.Contact = driverContact.String
		busDetails.BusType = busType.String
		busDetails.Capacity = capacity.Int64
		busDetails.BusNumber = busNumber.String
		busDetails.Reviews = []models.ReviewResponseOperator{}
		var reviewDetails models.ReviewResponseOperator
		reviewDetails.RatingID = ratingID.String
		reviewDetails.UserID = userID.String
		reviewDetails.Comment = comment.String
		reviewDetails.Rating = rating.Int64

		if busID.Valid {
			if busIndex(operator.Buses, busDetails) == -1 {
				operator.Buses = append(operator.Buses, busDetails)
				operator.Buses[len(operator.Buses)-1].Reviews = append(operator.Buses[len(operator.Buses)-1].Reviews, reviewDetails)
			} else {
				operator.Buses[busIndex(operator.Buses, busDetails)].Reviews = append(operator.Buses[busIndex(operator.Buses, busDetails)].Reviews, reviewDetails)
			}
		}

	}
	return http.StatusOK, nil, operator

}

func isOperatorExistsInData(data []models.GetOperatorResponse, operatorid string) int {
	for i := 0; i < len(data); i++ {
		if data[i].OperatorID == operatorid {
			return i
		}
	}
	return -1
}

func GetAllOperatorDetails(db *sql.DB) (int, error, []models.GetOperatorResponse) {
	var operators []models.GetOperatorResponse
	rows, err := db.Query(`
	SELECT operator.*,
	buses.busID,buses.driverContact,buses.capacity,buses.busType,buses.busNumber,
	reviews.ratingID,reviews.userID, reviews.comment,reviews.rating
	FROM operator
	LEFT JOIN buses
	ON operator.operatorID = buses.operatorID
	LEFT JOIN reviews
	ON buses.busID = reviews.busID`)
	if err != nil {
		return http.StatusInternalServerError, err, operators
	}
	defer rows.Close()

	for rows.Next() {

		var operator models.GetOperatorResponse

		var busID sql.NullString
		var driverContact sql.NullString
		var capacity sql.NullInt64
		var busType sql.NullString
		var busNumber sql.NullString
		var ratingID sql.NullString
		var userID sql.NullString
		var comment sql.NullString
		var rating sql.NullInt64

		err := rows.Scan(&operator.OperatorID, &operator.OperatorName, &operator.Contact,
			&operator.Address, &operator.Email, &operator.Password, &busID, &driverContact,
			&capacity, &busType, &busNumber, &ratingID, &userID, &comment, &rating)
		if err != nil {
			return http.StatusInternalServerError, err, operators
		}

		index := isOperatorExistsInData(operators, operator.OperatorID)

		var busDetails models.BusDetailsOperator
		busDetails.BusId = busID.String
		busDetails.Contact = driverContact.String
		busDetails.BusType = busType.String
		busDetails.Capacity = capacity.Int64
		busDetails.BusNumber = busNumber.String
		busDetails.Reviews = []models.ReviewResponseOperator{}
		var reviewDetails models.ReviewResponseOperator
		reviewDetails.RatingID = ratingID.String
		reviewDetails.UserID = userID.String
		reviewDetails.Comment = comment.String
		reviewDetails.Rating = rating.Int64

		if index == -1 {
			if busID.Valid {
				if busIndex(operator.Buses, busDetails) == -1 {
					operator.Buses = append(operator.Buses, busDetails)
					operator.Buses[len(operator.Buses)-1].Reviews = append(operator.Buses[len(operator.Buses)-1].Reviews, reviewDetails)
				} else {
					operator.Buses[busIndex(operator.Buses, busDetails)].Reviews = append(operator.Buses[busIndex(operator.Buses, busDetails)].Reviews, reviewDetails)
				}
			}
			operators = append(operators, operator)
		} else if busID.Valid {
			if busIndex(operators[index].Buses, busDetails) == -1 {
				operators[index].Buses = append(operators[index].Buses, busDetails)
				operators[index].Buses[len(operators[index].Buses)-1].Reviews = append(operators[index].Buses[len(operators[index].Buses)-1].Reviews, reviewDetails)
			} else {
				operators[index].Buses[busIndex(operators[index].Buses, busDetails)].Reviews = append(operators[index].Buses[busIndex(operators[index].Buses, busDetails)].Reviews, reviewDetails)
			}
		}

	}

	return http.StatusOK, nil, operators
}

func busIndex(buses []models.BusDetailsOperator, bus models.BusDetailsOperator) int {
	for i := 0; i < len(buses); i++ {
		if buses[i].BusId == bus.BusId {
			return i
		}
	}
	return -1
}

func IsOperatorExist(db *sql.DB, Email string) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM operator WHERE email = ?", Email).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("Operator already exist")
	}
	return nil
}

func IsOperatorIdValid(db *sql.DB, operatorId string) error {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM operator WHERE operatorID = ?", operatorId).Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		return errors.New("Operator doesnot exist")
	}
	return nil
}
