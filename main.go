package main

import (
	"encoding/json"
	"facebook_golang/controllers"
	"facebook_golang/db"
	"facebook_golang/entity"
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
	router.HandleFunc("/signup", func(w http.ResponseWriter, r *http.Request) {
		var user entity.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := controllers.CreateUser(user)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}

		}
	}).Methods(http.MethodPost)
	router.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		var status entity.Status
		err := json.NewDecoder(r.Body).Decode(&status)
		if err != nil {
			respondWithError(err, w)
		} else {
			err := controllers.CreateStatus(status)
			if err != nil {
				respondWithError(err, w)
			} else {
				respondWithSuccess(true, w)
			}
		}
	}).Methods(http.MethodPost)
}

func stringToInt64(s string) (int64, error) {
	number, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}
	return number, err
}

func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
