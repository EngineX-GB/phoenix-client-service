package dao

import (
	"database/sql"
	"phoenix-client-service/datasource"
	"phoenix-client-service/model"
	"phoenix-client-service/util"
)

func UpdateDataOverride(userId string, overrideDataString string) error {
	db := datasource.Connect()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	defer db.Close()

	_, err = tx.Exec(util.UpdateOverrideByUserId(userId, overrideDataString))
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		print("An error occured while trying to add an override : " + err.Error())
		return err
	}
	return nil
}

func AddDataOverride(userId string, overrideDataString string) error {
	db := datasource.Connect()

	// create a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	defer db.Close()

	_, err = tx.Exec(util.AddOverride(userId, overrideDataString))
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		print("An error occured while trying to add an override : " + err.Error())
		return err
	}
	return nil
}

func DeleteDataOverride(userId string) error {
	db := datasource.Connect()

	// create a transaction
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()
	defer db.Close()

	_, err = tx.Exec(util.DeleteOverrideByUserId(userId))
	if err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		print("An error occured while trying to delete an override : " + err.Error())
		return err
	}
	return nil
}

func GetOverrideByUserId(userId string) (string, error) {
	db := datasource.Connect()
	defer db.Close()
	var overridesJson string
	err := db.QueryRow(util.GetAllOverridesDataQueryByUserId(userId)).Scan(&overridesJson)

	if err == sql.ErrNoRows {
		return "[]", nil
	}

	if err != nil {
		return "", err
	}
	return overridesJson, nil
}

func GetAllOverrideData() ([]model.OverrideRecord, error) {
	var overrideRecords []model.OverrideRecord
	var overrideRecord model.OverrideRecord
	db := datasource.Connect()
	defer db.Close()

	rows, err := db.Query(util.GetAllOverridesQuery())

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&overrideRecord.Oid, &overrideRecord.UserId, &overrideRecord.UserName, &overrideRecord.Data); err != nil {
			return overrideRecords, err
		}
		overrideRecords = append(overrideRecords, overrideRecord)
	}
	return overrideRecords, nil
}
