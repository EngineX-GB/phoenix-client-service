package dao

import (
	"phoenix-client-service/datasource"
	"phoenix-client-service/util"
)

func AddLink(userId1 string, userId2 string, inputType string, notes string) (bool, error) {
	db := datasource.Connect()
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	defer tx.Rollback()
	defer db.Close()
	_, err = tx.Exec(util.AddLink(), userId1, userId2, inputType, notes)
	if err != nil {
		return false, err
	}
	if err := tx.Commit(); err != nil {
		print("An error occured while trying to add a link : " + err.Error())
		return false, err
	}
	return true, nil
}

func RemoveLink(userId1 string, userId2 string) (bool, error) {
	db := datasource.Connect()
	tx, err := db.Begin()
	if err != nil {
		return false, err
	}
	_, err = tx.Exec(util.RemoveLink(), userId1, userId2)
	if err != nil {
		return false, err
	}
	if err := tx.Commit(); err != nil {
		print("An error occured while trying to delete a link: " + err.Error())
		return false, err
	}
	return true, nil
}
