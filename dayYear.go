package main

import (
	"time"
)

func dayOfYear (dateString string) int{
	parsedDate, err := time.Parse("Monday, January 2" , dateString)
	if err != nil {
		panic(err)
	}
	day := parsedDate.YearDay()
	return day
}