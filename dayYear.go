package main

import (
	"time"
)

func dayOfYear (dateString OldRelease) int{
	parsedDate, err := time.Parse("Monday, January 2" , dateString.release_date)
	if err != nil {
		panic(err)
	}

	day := parsedDate.YearDay()

	// fmt.Println("day of year:", day)

	return day
}