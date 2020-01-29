package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/Reckner/rad_product/src/helpers"
)

// CreateConnection - creates database connection
func CreateConnection() (*sql.DB, error) {
	host := "host.docker.internal"
	user := "root"
	password := "root"
	dbName := "db"

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:3308)/%s?parseTime=true", user, password, host, dbName)
	database, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func InitializeDB() {
	db, err := CreateConnection()
	if err != nil {
		panic(err)
	}

	result, err := db.Query("SHOW TABLES")
	if err != nil {
		log.Fatal(err)
	}
	var tableName string
	var tableNames []string
	for result.Next() {
		_ = result.Scan(&tableName)
		tableNames = append(tableNames, tableName)
	}

	if len(tableNames) == 0 {
		query := helpers.Files.ReadFile("src/database/init.sql")

		_, err = db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	defer result.Close()
	db.Close()
}
