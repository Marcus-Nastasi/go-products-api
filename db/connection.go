package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "123"
	dbname = "products"
)

func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname= %s sslmode=disable",
		host, port, user, pass, dbname,
	)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("error connecting to db " + err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic("error pinging on db " + err.Error())
	}
	fmt.Println("Database: connected to " + dbname)
	fmt.Println()
	return db, nil
}
