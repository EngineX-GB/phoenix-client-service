package dao

import (
	"database/sql"
	"phoenix-client-service/model"
	"phoenix-client-service/util"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteClientTrackerChangesQuery(userId string) ([]model.TrackerChange, error) {
	db := connect()
	rows, err := db.Query(util.GetClientTrackerChanges(userId))

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	//TODO: Finish the function here
	return nil, nil
}

func ExecuteFeedbackQuery(userId string) ([]model.Feedback, error) {
	db := connect()
	rows, err := db.Query(util.GetFeedbackData(userId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()

	var feedbackEntries = make([]model.Feedback, 0)
	for rows.Next() {
		var feedback model.Feedback
		if err := rows.Scan(&feedback.Oid,
			&feedback.ServiceProvider,
			&feedback.Rating,
			&feedback.Comment,
			&feedback.FeedbackDate); err != nil {
			return feedbackEntries, err
		}
		feedbackEntries = append(feedbackEntries, feedback)
	}
	if err = rows.Err(); err != nil {
		return feedbackEntries, err
	}
	return feedbackEntries, nil
}

func ExecuteSearchQuery(searchRequest model.SearchRequest) ([]model.Client, error) {
	db := connect()
	rows, err := db.Query(util.GetClientsAvailableToday(searchRequest))

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	var entries = make([]model.Client, 0)
	for rows.Next() {
		var client model.Client
		if err := rows.Scan(&client.Oid,
			&client.Username,
			&client.Nationality,
			&client.Location,
			&client.Rating,
			&client.Age,
			&client.R15,
			&client.R30,
			&client.R45,
			&client.R100,
			&client.R150,
			&client.R200,
			&client.R250,
			&client.R300,
			&client.R350,
			&client.R400,
			&client.RON,
			&client.Telephone,
			&client.UserId,
			&client.Region,
			&client.MemberSince,
			&client.Height,
			&client.DSize,
			&client.HairColor,
			&client.EyeColor); err != nil {
			return entries, err
		}
		entries = append(entries, client)
	}
	if err = rows.Err(); err != nil {
		return entries, err
	}
	return entries, nil
}

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/db_phoenix?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	return db
}
