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



type Data struct{
	date_posted time.Time `json:"date_posted"`
	date_string string `json:"date_string"`
	current_cfs string `json:"current_cfs"`
	time_posted string `json:"time_posted`
	forecast    []string `json:"forecast"`
	expires   string `json:"expires"`
}

type Schedule struct{
	id int64
	release_date string
	last_flow string
	minimum_release bool
}

func runDB(date time.Time, currentDate string, cfs string, timePosted string, forecast []string, expire string) {
	// err := connectDB()
	// if err != nil {
	// 	log.Fatal("Error connecting to database", err)
	// }
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	supabaseUrl := os.Getenv("SUPABASE_URL")
	// supabaseKey := os.Getenv("SUPABASE_KEY")
	// supabase := supa.CreateClient(supabaseUrl, supabaseKey)

	conn, err := pgx.Connect(context.Background(), supabaseUrl)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "SELECT * FROM scheduled_release")
	if err != nil {
		log.Fatal("Error querying database:", err)
	}
	defer rows.Close()

	var schedules []Schedule
	for rows.Next() {
    var s Schedule
    err := rows.Scan(s.id ,&s.release_date, &s.last_flow, s.minimum_release)
    if err != nil {
        log.Fatal(err)
    }
   schedules = append(schedules, s)
}
if err := rows.Err(); err != nil {
    log.Fatal(err)
}

	//insert data into supabase
	// row := Data{
	// 	date_posted: date,
	// 	date_string: currentDate,
	// 	current_cfs: cfs,
	// 	time_posted: timePosted,
	// 	forecast:    forecast,
	// 	expires:     expire,
	// }

	// var results []Data
	// err := supabase.DB.From("daily_data").Insert(row).Execute(&results)
	// if err != nil{
	// 	panic(err)
	// }
	
	// var results map[string]interface{}
	// err = supabase.DB.From("scheduled_release").Select("*").Execute(&results)
	// if err != nil {
	// 	panic(err)
	// }


	fmt.Println("data from scheduled release", schedules, "END")
	// fmt.Println(results)

}
