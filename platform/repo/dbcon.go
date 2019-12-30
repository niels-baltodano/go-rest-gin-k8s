package repo

import (
	"fmt"
	"database/sql"
	)

func GetConexion() (*sql.DB, error) {
	var err error
	var db *sql.DB
	dbInfo := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", DbHost, DbUser, DbPassword, DbName)
	db, err = sql.Open("postgres", dbInfo)
	if err != nil {
		return nil,err
	}
	return db,nil
}