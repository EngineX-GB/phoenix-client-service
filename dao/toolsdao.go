package dao

import (
	"log"
	"phoenix-client-service/datasource"
	"phoenix-client-service/model"
	"phoenix-client-service/util"
)

func GetTelephoneNumberSearchResults(queryStringArgument string) ([]model.TelephoneSearchResult, error) {
	db := datasource.Connect()
	defer db.Close()

	print(util.GetTelephoneNumberSearchResults(queryStringArgument))

	rows, err := db.Query(util.GetTelephoneNumberSearchResults(queryStringArgument))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var records = make([]model.TelephoneSearchResult, 0)
	for rows.Next() {
		var record model.TelephoneSearchResult
		if err := rows.Scan(&record.UserId, &record.Username, &record.Nationality, &record.Ethnicity, &record.Telephone); err != nil {
			return records, err
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return records, err
	}

	return records, nil
}

func GetUsernameSearchResults(queryStringArgument string) ([]model.UsernameSearchResult, error) {
	db := datasource.Connect()
	defer db.Close()

	print(util.GetUsernameSearchResults(queryStringArgument))

	rows, err := db.Query(util.GetUsernameSearchResults(queryStringArgument))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()
	var records = make([]model.UsernameSearchResult, 0)
	for rows.Next() {
		var record model.UsernameSearchResult
		if err := rows.Scan(&record.UserId, &record.Username, &record.CurrentUsername, &record.Nationality, &record.Ethnicity, &record.UsernameInTodaysRecord); err != nil {
			return records, err
		}
		records = append(records, record)
	}
	if err = rows.Err(); err != nil {
		return records, err
	}

	return records, nil
}
