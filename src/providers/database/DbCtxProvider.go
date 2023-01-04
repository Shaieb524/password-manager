package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetDbConnectionStringFromEnv() string {
	var host, username, password, databaseName, port string

	if os.Getenv("DB_HOST") == "" {
		host = "localhost"
	} else {
		host = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_USER") == "" {
		username = "postgres"
	} else {
		username = os.Getenv("DB_HOST")
	}

	if os.Getenv("DB_PASSWORD") == "" {
		password = "mypassword"
	} else {
		password = os.Getenv("DB_PASSWORD")
	}

	if os.Getenv("DB_NAME") == "" {
		databaseName = "PmDb"
	} else {
		databaseName = os.Getenv("DB_NAME")
	}

	if os.Getenv("DB_PORT") == "" {
		port = "5432"
	} else {
		port = os.Getenv("DB_PORT")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos",
		host, username, password, databaseName, port)

	return dsn
}

func NewDatabaseContext() *gorm.DB {
	connString := GetDbConnectionStringFromEnv()
	Db, err := gorm.Open(postgres.Open(connString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return Db
}
