package main

import (
	"facebook_golang/db"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := db.GetCon()
	if err != nil {
		log.Printf("Error with database " + err.Error())
		return
	} else {
		err = db.Ping()
		if err != nil {
			log.Printf("Error with connection to db" + err.Error())
			return
		}
	}

	log.Println("Starting the Http server ")
	router := mux.NewRouter()
	initRoutes(router)
	port := ":8000"

	server := &http.Server{
		Handler: router,
		Addr:    port,
		// timeouts so the server never waits forever...
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Server started at %s", port)
	log.Fatal(server.ListenAndServe())
}




