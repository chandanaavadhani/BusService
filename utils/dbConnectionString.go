package utils

import (
	"fmt"
	"os"
)

func GetConnectionString() string {
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "192.168.1.67"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "3306"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "sravan"
	}

	password := os.Getenv("DB_PASS")
	if password == "" {
		password = "Sravan@2509"
	}

	dbName := os.Getenv("DB_NAME")
	if dbName == "" {
		dbName = "golang"
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)
}
