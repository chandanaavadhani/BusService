package repository

import (
	"database/sql"
	"fmt"

	utils "github.com/chandanaavadhani/BusService/utils"
	_ "github.com/go-sql-driver/mysql"
)

func DBConnection() (*sql.DB, error) {
	db, err := sql.Open("mysql", utils.GetConnectionString())
	if err != nil {
		fmt.Println("Error in connecting the DB", err)
		return nil, err
	}
	fmt.Println("Database is connected successfully")
	return db, nil
}
