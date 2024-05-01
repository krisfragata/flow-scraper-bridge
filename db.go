package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)


type OldRelease struct {
	release_date string
}

type Data struct {
	date_posted time.Time `json:"date_posted"`
	date_string string    `json:"date_string"`
	current_cfs string    `json:"current_cfs"`
	time_posted string    `json:"time_posted`
	forecast    []string  `json:"forecast"`
	expires     string    `json:"expires"`
}

func runDB(date time.Time, currentDate string, cfs string, timePosted string, forecast string, expire string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")

	conn, err := pgx.Connect(context.Background(), supabaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	// insert data into supabase
	row := []any{
		 date,
	currentDate,
		 cfs,
		 timePosted,
	  forecast,
		  expire,
	}

	// fmt.Println("data from scheduled release", row, "END")
	for _, val := range row{
		fmt.Println("data:", val)
	}
	current := dayOfYear(currentDate)
	isReleaseToday := isRelease(current)
	fmt.Println("is it a release today?", isReleaseToday)

	//begin query
	query := `INSERT INTO daily_data (date_posted, date_string, current_cfs, time_posted, forecast, expires, scheduled_release) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	rows, err := conn.Query(context.Background(), query, date, currentDate, cfs, timePosted, forecast, expire, isReleaseToday )
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()
	// fmt.Println(rows)
}

