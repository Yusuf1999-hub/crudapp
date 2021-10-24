package utils

import "fmt"

var (
	dbHost     = "localhost"
	dbUser     = "yusufkhuja"
	dbPassword = "y6451895"
	dbName     = "crud_app"
	dbPort     = 5432
)

var Config = fmt.Sprintf("host=%s user=%s dbname=%s port=%d password=%s",
	dbHost, dbUser, dbName, dbPort, dbPassword)
