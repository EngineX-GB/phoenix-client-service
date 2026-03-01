package util

import (
	"phoenix-client-service/model"
	"strconv"
	"strings"
)

func GetClientsAvailableToday(searchRequest model.SearchRequest) string {
	var stringBuilder strings.Builder
	if searchRequest.Username != "" {
		stringBuilder.WriteString(" AND c.username like '%" + searchRequest.Username + "%'")
	}
	if searchRequest.UserId != "" {
		stringBuilder.WriteString(" AND c.user_Id = '" + searchRequest.UserId + "'")
	}
	if searchRequest.Nationality != "" {
		stringBuilder.WriteString(" AND c.nationality = '" + searchRequest.Nationality + "'")
	}
	if searchRequest.Region != "" {
		stringBuilder.WriteString(" AND c.region = '" + searchRequest.Region + "'")
	}
	conditionalClause := stringBuilder.String()
	// if len(conditionalClause) > 0 {
	// 	conditionalClause = conditionalClause[:len(conditionalClause)-1]
	// }

	print(conditionalClause)
	return "SELECT c.oid, c.username, c.nationality, c.location, c.rating, " +
		"c.age, c.rate_15_min, c.rate_30_min, c.rate_45_min, c.rate_1_hour, c.rate_1_50_hour, c.rate_2_hour, c.rate_2_50_hour, " +
		"c.rate_3_hour, c.rate_3_50_hour, c.rate_4_hour, c.rate_overnight, c.telephone, c.user_id, c.region, c.member_since, " +
		"c.height, c.dress_size, c.hair_colour, c.eye_colour, " +
		"ca.minimum_charge, ca.maximum_charge, ca.number_of_days_in_service, ca.percentage_available, ca.total_regions_travelled, ca.previously_serviced_bb " +
		"FROM tbl_client c " +
		"INNER JOIN tbl_client_analytics ca ON ca.user_id = c.user_id " +
		"where date(c.refresh_time) = date(now())" + conditionalClause
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
	return "insert into tbl_order (user_id, username, location, region, date_of_event, time_of_event, creation_date, modification_date, duration, rate, deductions, surplus, price, status, notes) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetAllOrders() string {
	return "select oid, user_id, username, location, region, date_of_event, time_of_event, duration, rate, deductions, surplus, price, status, notes from tbl_order"
}

func GetOrdersByYear(year int) string {
	return "select oid, user_id, username, location, region, date_of_event, time_of_event, duration, rate, deductions, surplus, price, status, notes from tbl_order where year(date_of_event) = " + strconv.Itoa(year) + ""
}

func GenerateOrderRequest(userId string) string {
	return "select user_id, username, location, region, telephone, rate_15_min, rate_30_min, rate_45_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight from tbl_client where user_id = '" + userId + "' limit 1"
}

func UpdateOrderStatus() string {
	return "update tbl_order set status = ? where oid = ?"
}

func AddLink() string {
	return "insert into tbl_link (user_id_1, user_id_2, input_type, notes) values (?, ?, ?, ?)"
}

func RemoveLink() string {
	return "delete from tbl_link where user_id_1 = ? and user_id_2 = ?"
}

func GetAllLinks() string {
	return "select user_id_1, user_id_2, input_type, notes from tbl_link"
}
