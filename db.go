package main

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/supabase-community/supabase-go"
	"github.com/joho/godotenv"
)

var supabaseUrl string
var supabaseKey string
var supabaseClient *supabase.Client

func runDB(date time.Time, currentDate string, cfs string, forecast []string, expire string) {
	// err := connectDB()
	// if err != nil {
	// 	log.Fatal("Error connecting to database", err)
	// }
	err := insertData()
	if err != nil {
		log.Fatal("Error inserting data into database", err)
	}
}

func init(){
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	supabaseUrl = os.Getenv("SUPABASE_URL")
	supabaseKey = os.Getenv("SUPABASE_KEY")

	//Initialize supabase client
	_, err := supabase.NewClient(supabaseUrl, supabaseKey, nil)
	if err != nil {
		fmt.Println("cannot initialize client:", err)
	}

	

}



func insertData() error {
	return nil
}
