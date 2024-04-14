package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"errors"

	//import colly & goQuery for use
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//http://www.h2oline.com/srcs/255122.html

//create a function that checks if data should be posted to db

func main() {

	var htmlContent string

	c := colly.NewCollector(
		colly.AllowURLRevisit(),
	)

	// be a good citizen and limit domain visits
	c.Limit(&colly.LimitRule{
		DomainGlob: "h2oline.com",
		Delay: 60 *time.Second,
		RandomDelay: 1 * time.Second,
	})

	c.OnResponse(func(r *colly.Response){
		htmlContent = string(r.Body)
	})
	err := c.Visit("http://www.h2oline.com/srcs/255122.html")
	if err != nil {
		log.Fatal(err)
	}
	
	extractData(htmlContent)
	

}

func extractData(htmlContent string) {
	//use goquery to find text-line needed
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}

	recentPosting := getRecentPosting(doc)
	cfs, _ := extractCFS(recentPosting)
	timePosted, _ := extractTimePosted(recentPosting)

	fmt.Printf("%v CFS @ %v", cfs, timePosted)

}

func getRecentPosting(doc *goquery.Document) string {
	found := false
	var recentPosting string
	doc.Find("div").Each(func(i int, e *goquery.Selection) {
		if strings.Contains(e.Text(), "the total flow below the dam was") {
			found = true
			recentPosting = e.Text()
			return
		}
	})
	if !found {
		log.Fatal("recent posting not found")
	}
	return recentPosting
}

func extractCFS(recentPosting string) (string, error) {
	parts := strings.Fields(recentPosting)
	if len(parts) >= 8 {
		cfs := parts[11]
		return cfs, nil
	} else {
		fmt.Println("Unable to extract time and CFS measurement.")
		return "", errors.New("unable to extract CFS") 
	}
}

func extractTimePosted(recentPosting string) (string, error){
	parts := strings.Fields(recentPosting)
	if len(parts) >= 8 {
		time := parts[1] + " " + parts[2] + " " + parts[3]
		return time, nil
	} else {
		fmt.Println("Unable to extract time and CFS measurement.")
		return "", errors.New("unable to extract time posted")
	}
}

func extractForecast(htmlContent string) (string, error) {
	found := false
	var forecast string
	doc.Find("div").Each(func(i int, e *goquery.Selection) {
		if strings.Contains(e.Text(), "the total flow below the dam was") {
			found = true
			forecast = e.Text()
			return
		}
	})
	if !found {
		log.Fatal("recent posting not found")
	}
	return forecast
}