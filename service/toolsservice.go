package service

import (
	"log"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
	"strings"
)

func SearchTelephoneNumber(telephoneNumberString string) ([]model.TelephoneSearchResult, error) {
	// check the input string to see if it has a wildcard at the start or the end of the string
	cleanInputString := telephoneNumberString
	if strings.HasPrefix(telephoneNumberString, "*") {
		cleanInputString = "%" + strings.TrimPrefix(cleanInputString, "*")
	} else if strings.HasSuffix(telephoneNumberString, "*") {
		cleanInputString = strings.TrimSuffix(cleanInputString, "*") + "%"
	} else {
		cleanInputString = "%" + cleanInputString + "%"
	}

	records, err := dao.GetTelephoneNumberSearchResults(cleanInputString)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return records, err
}

func SearchUsername(usernameString string) ([]model.UsernameSearchResult, error) {
	// check the input string to see if it has a wildcard at the start or the end of the string
	cleanInputString := usernameString
	if strings.HasPrefix(usernameString, "*") {
		cleanInputString = "%" + strings.TrimPrefix(cleanInputString, "*")
	} else if strings.HasSuffix(usernameString, "*") {
		cleanInputString = strings.TrimSuffix(cleanInputString, "*") + "%"
	} else {
		cleanInputString = "%" + cleanInputString + "%"
	}

	records, err := dao.GetUsernameSearchResults(cleanInputString)

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return records, err
}
