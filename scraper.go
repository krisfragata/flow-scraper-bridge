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

	publishExpire := getPublishExpire(doc)
	currentDateArray, _ := extractCurrentDate(publishExpire)
	expiresAt, _ := extractExpireDate(publishExpire) 

	recentPosting := getRecentPosting(doc)
	cfs, _ := extractCFS(recentPosting)
	timePosted, _ := extractTimePosted(recentPosting)

	forecastArray := extractForecast(doc)


	fmt.Printf("%v CFS @ %v\n", cfs, timePosted)
	fmt.Println(forecastArray)
	fmt.Println("forecast length:", len(forecastArray))
	fmt.Println("current date and expiry:", publishExpire)
	fmt.Println("published:", currentDateArray)
	fmt.Println("expires at:", expiresAt)

}

func getPublishExpire(doc *goquery.Document) (string){
 found := false
 var desiredLine string
 doc.Find("div").Each(func(i int, e *goquery.Selection) {
		text := e.Text()
		if strings.Contains(text, "Published:") {
			found = true
			lines := strings.Split(text, "\n")
			for _, line := range lines {
				if strings.Contains(line, "Published:") && strings.Contains(line, "Expires:") {
					desiredLine = strings.TrimSpace(line)
					break
				}
			}			
		}
	})

 if !found {
	log.Fatal("current date not found")
 }
 return desiredLine
}

func extractCurrentDate(publishExpire string) ([]string, error) {
	parts := strings.Fields(publishExpire)
	var dateArray []string
	if len(parts) >= 8 {
		for i, v := range parts {
			if i >= 1  && i < 7 {
				dateArray = append(dateArray, v)
			}
		}
	} else {
		fmt.Println("Unable to extract publish date.")
		return  []string{""}, errors.New("unable to extract publish date") 
	}
	return dateArray, nil
}

func extractExpireDate(publishExpire string) ([]string, error) {
	parts := strings.Fields(publishExpire)
	var expireArray []string
	if len(parts) >= 8 {
		for i, v := range parts {
			if  i > 8 {
				expireArray = append(expireArray, v)
			}
		}
	} else {
		fmt.Println("Unable to extract expire date.")
		return  []string{""}, errors.New("unable to extract expire date") 
	}
	return expireArray, nil
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
		fmt.Println("Unable to extract CFS measurement.")
		return "", errors.New("unable to extract CFS") 
	}
}

func extractTimePosted(recentPosting string) (string, error){
	parts := strings.Fields(recentPosting)
	if len(parts) >= 8 {
		time := parts[1] + " " + parts[2] + " " + parts[3]
		return time, nil
	} else {
		fmt.Println("Unable to extract time.")
		return "", errors.New("unable to extract time posted")
	}
}

func extractForecast(doc *goquery.Document) ([]string) {
	var forecastArray []string
	var forecastHtml string
	doc.Find("div:contains('The following forecast for flows') + div").Each(func(i int, e *goquery.Selection) {
			htmlContent, err := e.Html()
			if err != nil {
				log.Fatal("forecast not found")
				return
			}
			forecastHtml = htmlContent
	})

	forecastArray = strings.Split(forecastHtml, "<br/>")
	return forecastArray
}