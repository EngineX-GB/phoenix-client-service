package service

import (
	"encoding/csv"
	"log"
	"net/url"
	"phoenix-client-service/dao"
	"phoenix-client-service/model"
	"strconv"
	"strings"
	"time"
)

func ReadWatchListFeed(contents string) {
	reader := strings.NewReader(contents)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '|'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Println("Error in reading the CSV file")
		return
	}

	for _, record := range records {
		// a watchlist entry in the feed will only have one field in it: the URL containing the userId
		u, err := url.Parse(record[0])
		if err != nil {
			panic(err)
		}
		q := u.Query()
		userId := q.Get("userID")
		if userId != "" {
			dao.ExecuteAddWatchListEntry(userId)
		}
	}
}

func ReadOrderFeed(contents string) {
	reader := strings.NewReader(contents)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '|'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Println("Error reading the CSV file")
		return
	}

	var orderList []model.Order

	for _, record := range records {
		var order model.Order
		order.OrderReference = record[1]
		order.UserId = record[2]
		order.UserName = record[3]
		order.Location = record[4]
		order.DateOfEvent = record[5]
		order.TimeOfEvent = record[6]
		order.Duration = convertStrToInt(record[7])
		order.Rate = convertStrToInt(record[8])
		order.Deductions = convertStrToInt(record[9])
		order.Surplus = convertStrToInt(record[10])
		order.Price = convertStrToInt(record[11])
		order.Status = record[12]
		order.Notes = record[13]
		order.CreationDate = record[14]
		order.ModifiedDate = record[14]
		orderList = append(orderList, order)
	}

	// create some sets. One for EXECUTED and one for CANCELLED and put the appropriate order refs in them

	executedSet := make(map[string]struct{})
	cancelledSet := make(map[string]struct{})

	if len(orderList) > 0 {
		for _, order := range orderList {
			if order.Status == "EXECUTED" {
				executedSet[order.OrderReference] = struct{}{}
				dao.SaveOrder(order)
			}
			if order.Status == "CANCELLED" {
				cancelledSet[order.OrderReference] = struct{}{}
				dao.SaveOrder(order)
			}
		}

		for _, order := range orderList {
			if _, ok := executedSet[order.OrderReference]; !ok {
				if _, ok := cancelledSet[order.OrderReference]; !ok {
					if order.Status == "COMMITTED" {
						dao.SaveOrder(order)
					}
				}
			}
		}
	}
}

func convertStrToInt(value string) uint64 {
	if value == "Not Specified" {
		return 0
	}
	num, err := strconv.ParseUint(value, 10, 64)
	if err != nil {
		log.Fatal("Cannot convert string to int, " + err.Error())
	}
	return num
}

func convertStrToDate(value string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	refreshTime, err := time.Parse(timeLayout, value)
	if err != nil {
		log.Fatal("Cannot convert string to time, " + err.Error())
	}
	return refreshTime
}
