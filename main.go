package main

import (
	"encoding/json"
	"facebook_golang/controllers"
	"facebook_golang/db"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	db, err := db.GetCon()
	if err != nil {
		log.Printf("Error with database " + err.Error())
		return
	}else {
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

func initRoutes(router *mux.Router) {
	router.HandleFunc("/user/{id}", func(w http.ResponseWriter, r *http.Request) {
		idAsString := mux.Vars(r)["id"]
		id, err := stringToInt64(idAsString)
		if err != nil {
			respondWithError(err, w)
			return
		}
		user, err := controllers.GetUserByID(id)
		if err != nil {
			respondWithError(err, w)
		} else {
			respondWithSuccess(user, w)
		}
	}).Methods(http.MethodGet)
}

func stringToInt64(s string) (int64, error) {
	numero, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return numero, err
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
