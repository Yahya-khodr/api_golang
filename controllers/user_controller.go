package controllers

import (
	"facebook_golang/db"
	"facebook_golang/entity"
)

func GetUserByID(id int64) (entity.User, error) {
	var user entity.User
	con, err := db.GetCon()
	if err != nil {
		return user, err
	}
	row := con.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(user entity.User) error {
	con, err := db.GetCon()
	if err != nil {
		return err
	}
	_, err = con.Exec("INSERT INTO users(first_name,last_name,email,password) VALUES (?,?,?,?)",
		user.FirstName, user.LastName, user.Email, user.Password)
	return err

}
