package controllers

import (
	
	"facebook_golang/db"
	"facebook_golang/entity"
	

	
)



func GetUserByID(id int64) (entity.User, error) {
	var user entity.User
	bd, err := db.GetCon()
	if err != nil {
		return user, err
	}
	row := bd.QueryRow("SELECT id, first_name, last_name, email FROM users WHERE id = ?", id)
	err = row.Scan(&user.ID, &user.FirstName, &user.LastName, &user.Email)
	if err != nil {
		return user, err
	}
	return user, nil
}
