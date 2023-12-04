package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // postgres driverini ishga tushu=irish uchun
)

func ConnectToDbTest() (*sql.DB, error) {
	dataSourceName := "user=postgres dbname=restaurant password=postgres sslmode=disable"

	sqlDb, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Println("error on connecting psql: ", err)
		return nil, err
	}

	err = sqlDb.Ping()
	if err != nil {
		log.Println("error on checking: ", err)
		return nil, err
	}

	return sqlDb, nil
}

func ConnectToDb() *sql.DB {
	sqlDb, err := ConnectToDbTest()
	if err != nil {
		log.Println("error on connect postgres: ", err)
		return nil
	}

	return sqlDb
}
