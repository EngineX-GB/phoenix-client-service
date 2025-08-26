package util

import (
	"phoenix-client-service/model"
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
