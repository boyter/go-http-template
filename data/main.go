package data

import (
	"changeme/common"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDb(config common.Config) (*sql.DB, error) {
	// Connect to the database
	//db, err := sql.Open("mysql", config.MySQLConnection)
	//if err != nil {
	//	return nil, err
	//}
	//
	//return db, nil

	return nil, nil
}
