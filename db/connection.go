package db

import (
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "123"
	dbname = "postgres"
)

func ConnectDb() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s pass=%s dbname= %s", host, port, user, pass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic("error connecting to db")
	}
	err = db.Ping()
	if err != nil {
		panic("error pinging on db")
	}
	fmt.Println("Connected to " + dbname)
	return db, nil
}
