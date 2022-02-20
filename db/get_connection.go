package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func GetCon() (*sql.DB, error) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "fbmockupdb"

	return sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

}
