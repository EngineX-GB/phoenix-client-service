package dao

import (
	"database/sql"
	"phoenix-client-service/datasource"
	"phoenix-client-service/model"
	"phoenix-client-service/util"

	_ "github.com/go-sql-driver/mysql"
)

func ExecuteClientTrackerChangesQuery(userId string) ([]model.TrackerChange, error) {
	db := datasource.Connect()
	rows, err := db.Query(util.GetClientTrackerChanges(userId))

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	defer db.Close()

	//TODO: Finish the function here
	return nil, nil
}

func ExecuteMinMaxMarkersForFeedbackQuery(userId string) (model.QueryMarker, error) {
	db := datasource.Connect()
	defer db.Close()

	var marker model.QueryMarker
	var min sql.NullInt64
	var max sql.NullInt64

	row := db.QueryRow(util.GetMinMaxIdValuesForFeedbackRecord(userId))
	err := row.Scan(&min, &max)
	if err != nil {
		return marker, err
	}

	if min.Valid {
		marker.Min = min.Int64
	} else {
		marker.Min = 0
	}

	if max.Valid {
		marker.Max = max.Int64
	} else {
		marker.Max = 0
	}

	return marker, nil
}

func ExecuteServiceReportHeadlineQuery(userId string) ([]model.ServiceReportHeadline, error) {
	db := datasource.Connect()
	rows, err := db.Query(util.GetServiceReportHeadlines(userId))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	var serviceReportHeadlineEntries = make([]model.ServiceReportHeadline, 0)
	for rows.Next() {
		var serviceReportHeadline model.ServiceReportHeadline
		if err := rows.Scan(&serviceReportHeadline.Oid,
			&serviceReportHeadline.UserId,
			&serviceReportHeadline.MeetDate,
			&serviceReportHeadline.ReportRating,
			&serviceReportHeadline.HeadLine); err != nil {
			return serviceReportHeadlineEntries, err
		}
		serviceReportHeadlineEntries = append(serviceReportHeadlineEntries, serviceReportHeadline)
	}
	if err = rows.Err(); err != nil {
		return serviceReportHeadlineEntries, err
	}
	return serviceReportHeadlineEntries, nil
}

func ExecuteFeedbackQuery(userId string, offsetId int64, pageDirection string) ([]model.Feedback, error) {
	db := datasource.Connect()
	rows, err := db.Query(util.GetFeedbackData(userId, offsetId, pageDirection))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()

	var feedbackEntries = make([]model.Feedback, 0)
	for rows.Next() {
		var feedback model.Feedback
		if err := rows.Scan(&feedback.Oid,
			&feedback.ByUsername,
			&feedback.RatingDate,
			&feedback.Rating,
			&feedback.Disputed,
			&feedback.Feedback,
			&feedback.FeedbackResponse); err != nil {
			return feedbackEntries, err
		}
		feedbackEntries = append(feedbackEntries, feedback)
	}
	if err = rows.Err(); err != nil {
		return feedbackEntries, err
	}
	return feedbackEntries, nil
}

func ExecuteGetWatchlist(todaysList bool) ([]model.WatchListEntry, error) {
	db := datasource.Connect()
	var queryString string
	if todaysList {
		queryString = util.ShowTodaysWatchList()
	} else {
		queryString = util.ShowAllWatchListEntries()
	}
	rows, err := db.Query(queryString)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer db.Close()
	var watchListEntries = make([]model.WatchListEntry, 0)
	for rows.Next() {
		var watchListEntry model.WatchListEntry
		if err := rows.Scan(&watchListEntry.UserId,
			&watchListEntry.Username,
			&watchListEntry.Nationality,
			&watchListEntry.Telephone,
			&watchListEntry.Location,
			&watchListEntry.Rate); err != nil {
			return watchListEntries, err
		}
		watchListEntries = append(watchListEntries, watchListEntry)
	}
	if err = rows.Err(); err != nil {
		return watchListEntries, err
	}
	return watchListEntries, nil
}

func ExecuteDeleteWatchListEntry(userId string) (int64, error) {
	db := datasource.Connect()
	result, err := db.Exec(util.DeleteWatchListEntry(userId))
	if err != nil {
		print(err.Error())
		return -1, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		panic(err)
	}
	return rowsAffected, err
}

func ExecuteAddWatchListEntry(userId string) (int64, error) {
	db := datasource.Connect()
	result, err := db.Exec(util.AddEntryToWatchList(userId))
	defer db.Close()
	if err != nil {
		print(err.Error())
		return -1, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		panic(err)
	}
	return insertId, nil
}

func ExecuteSearchQuery(searchRequest model.SearchRequest) ([]model.Client, error) {
	db := datasource.Connect()
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
			&client.EyeColor,
			&client.MinimumCharge,
			&client.MaximumCharge,
			&client.NumberOfDaysInService,
			&client.PercentageAvailable,
			&client.TotalRegionsTravelled,
			&client.PreviouslyServicedBB,
		); err != nil {
			return entries, err
		}
		entries = append(entries, client)
	}
	if err = rows.Err(); err != nil {
		return entries, err
	}
	return entries, nil
}
