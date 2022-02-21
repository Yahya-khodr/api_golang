package controllers

import (
	"facebook_golang/db"
	"facebook_golang/entity"
)

func CreateStatus(status entity.Status) error {
	db, err := db.GetCon()
	if err != nil {
		return err
	}
	_, err = db.Exec("INSERT INTO posts(post_content,created_by) VALUES (?,?)", status.Content, status.UserID)
	return err

}

func UpdateStatus(status entity.Status) error {
	db, err := db.GetCon()
	if err != nil {
		return err
	}
	_, err = db.Exec("UPDATE posts SET post_content = ? WHERE post_id = ?", status.Content, status.ID)
	return err
}
