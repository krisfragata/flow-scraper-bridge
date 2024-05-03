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


type Data struct {
	date_posted time.Time `json:"date_posted"`
	date_string string    `json:"date_string"`
	current_cfs string    `json:"current_cfs"`
	time_posted string    `json:"time_posted`
	forecast    string  `json:"forecast"`
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

	current := dayOfYear(currentDate)
	isReleaseToday := isRelease(current)

	//check db's recent posting
	var existingData Data
	query := `SELECT date_string, current_cfs, time_posted FROM daily_data ORDER BY id DESC LIMIT 1`
	row := conn.QueryRow(context.Background(), query)
	err = row.Scan(&existingData.date_string, &existingData.current_cfs, &existingData.time_posted)
	if err != nil {
		log.Fatalf("Error querying database: %v", err)
	}


	//check if recent posting matches new data's cfs
	if existingData.date_string == currentDate && existingData.current_cfs == cfs && existingData.time_posted == timePosted {
		// Update the most recent row
		updateQuery := `UPDATE daily_data SET date_posted = $1, forecast = $3, expires = $4 WHERE date_string = $5 AND current_cfs = $6 AND time_posted = $2`
		_, err := conn.Exec(context.Background(), updateQuery, date, timePosted, forecast, expire, currentDate, cfs)
		if err != nil {
			log.Fatal("Error updating database:", err)
		}
		fmt.Println("Updated existing row in the database")
	} else {
		// Insert new data into the database
		insertQuery := `INSERT INTO daily_data (date_posted, date_string, current_cfs, time_posted, forecast, expires, scheduled_release) VALUES ($1, $2, $3, $4, $5, $6, $7)`
		_, err := conn.Exec(context.Background(), insertQuery, date, currentDate, cfs, timePosted, forecast, expire, isReleaseToday)
		if err != nil {
			log.Fatal("Error inserting into database:", err)
		}
		fmt.Println("Inserted new row into the database")
	}

}

