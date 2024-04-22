package main

import (
	"fmt"
	"log"
	"os"
	"time"
	// supa "github.com/nedpals/supabase-go"
	"github.com/jackc/pgx/v4"
	"context"
	"github.com/joho/godotenv"
)

type Release struct{
	release_date string
}

type Data struct{
	date_posted time.Time `json:"date_posted"`
	date_string string `json:"date_string"`
	current_cfs string `json:"current_cfs"`
	time_posted string `json:"time_posted`
	forecast    []string `json:"forecast"`
	expires   string `json:"expires"`
}

func runDB(date time.Time, currentDate string, cfs string, timePosted string, forecast []string, expire string) {
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")

	conn, err := pgx.Connect(context.Background(), supabaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT release_date FROM scheduled_release") 
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	var releases []Release
	for rows.Next() {
		var r Release
		err := rows.Scan(&r.release_date)
    if err != nil {
        log.Fatal(err)
    }
   releases = append(releases, r)
	}
	if err := rows.Err(); err != nil {
    log.Fatal(err)
	}
	

	fmt.Println(releases)

	// insert data into supabase
	row := Data{
		date_posted: date,
		date_string: currentDate,
		current_cfs: cfs,
		time_posted: timePosted,
		forecast:    forecast,
		expires:     expire,
	}


	fmt.Println("data from scheduled release", row, "END")
	// fmt.Println(results)
}
