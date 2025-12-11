package util

import (
	"phoenix-client-service/model"
	"strconv"
	"strings"
)

func GetClientsAvailableToday(searchRequest model.SearchRequest) string {
	var stringBuilder strings.Builder
	if searchRequest.Username != "" {
		stringBuilder.WriteString(" AND username = '" + searchRequest.Username + "',")
	} else if searchRequest.UserId != "" {
		stringBuilder.WriteString(" AND user_Id = '" + searchRequest.UserId + "',")
	} else if searchRequest.Nationality != "" {
		stringBuilder.WriteString(" AND nationality = '" + searchRequest.Nationality + "',")
	} else if searchRequest.Region != "" {
		stringBuilder.WriteString(" AND region = '" + searchRequest.Region + "',")
	}
	conditionalClause := stringBuilder.String()
	if len(conditionalClause) > 0 {
		conditionalClause = conditionalClause[:len(conditionalClause)-1]
	}
	return "SELECT oid, username, nationality, location, rating, " +
		"age, rate_15_min, rate_30_min, rate_45_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, " +
		"rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight, telephone, user_id, region, member_since, " +
		"height, dress_size, hair_colour, eye_colour " +
		"FROM tbl_client where date(refresh_time) = date(now())" + conditionalClause
}

func GetClientTrackerChanges(userId string) string {
	return "select oid, username, user_id, field_value, old_value, new_value, record_datetime from vw_view_client_changes"
}

func GetFeedbackData(userId string) string {
	return "select oid, service_provider, rating, comment, feedback_date from tbl_feedback where user_id = " + userId
}

// write query to add a new watchlist entry -> error when the entry exists
func AddEntryToWatchList(userId string) string {
	return "insert into tbl_client_watchlist (user_id) values ('" + userId + "')"
}

// write a query to remove an existing watchlist entry -> error when the entry does not exist
func RemoveEntryFromWatchList(userId string) string {
	return "delete from tbl_client_watchlist where user_id = '" + userId + "'"
}

// write a query to show watchlist entries of clients that are available today
func ShowTodaysWatchList() string {
	return "select user_id, username, nationality, telephone, location, rate_1_hour from vw_todays_watchlist"
}

// write a query to show ALL watchlist entries.
func ShowAllWatchListEntries() string {
	return "select c.user_id, c.username, c.nationality, c.telephone, c.location, c.rate_1_hour from tbl_client c inner join tbl_client_watchlist tcw on c.user_id = tcw.user_id"
}

func DeleteWatchListEntry(userId string) string {
	return "delete from tbl_client_watchlist where user_Id = '" + userId + "'"
}

func SaveOrderEntry() string {
	return "insert into tbl_order (user_id, location, region, date_of_event, time_of_event, creation_date, modification_date, duration, rate, deductions, surplus, price, status, notes) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetAllOrders() string {
	return "select user_id, location, region, date_of_event, time_of_event, duration, rate, deductions, surplus, price, status, notes from tbl_order"
}

func GetOrdersByYear(year int) string {
	return "select user_id, location, region, date_of_event, time_of_event, duration, rate, deductions, surplus, price, status, notes from tbl_order where year(date_of_event) = " + strconv.Itoa(year) + ""
}
