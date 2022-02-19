package db
import (
	"database/sql"
	"log"
	_ "github.com/go-sql-driver/mysql"
)



func GetCon() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := ""
	dbName := "golang"

	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)

	if err != nil {
		log.Println("Connection Failed ")
	} else {
		log.Println("Connection Established")
	}
	return db

}