package main

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"
	//import colly & goQuery for use
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
)

//http://www.h2oline.com/srcs/255122.html

//create a function that checks if data should be posted to db

func main(){
	htmlContent := visitSite()	
	date, currentDate, cfs, timePosted, forecast, expire := extractData(htmlContent)
	runDB(date, currentDate, cfs, timePosted, forecast, expire)
}

func visitSite() string {
	var htmlContent string
	//create collector
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
	return htmlContent	
}

func extractData(htmlContent string) (date time.Time, currentDate string, cfs string, timePosted string, forecast string, expire string ) {
	//use goquery to find text-line needed
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatal(err)
	}
	publishExpire := getPublishExpire(doc)
	currentDate, _ = extractCurrentDate(publishExpire)
	date = time.Now().UTC()
	expire , _ = extractExpireDate(publishExpire) 
	recentPosting := getRecentPosting(doc)
	cfs, _ = extractCFS(recentPosting)
	timePosted, _ = extractTimePosted(recentPosting)
	forecastHtml := extractForecastHtml(doc)
	forecast = extractForecastStr(forecastHtml)


	return date, currentDate, cfs, timePosted, forecast, expire


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

func extractCurrentDate(publishExpire string) (string, error) {
	parts := strings.Fields(publishExpire)
	var date string

	if len(parts) >= 8 {
		for i, v := range parts {
			if i >= 1  && i < 4 {
				date += v + " "
			}
		}
	} else {
		fmt.Println("Unable to extract publish date.")
		return  "", errors.New("unable to extract publish date") 
	}
	date = strings.TrimSpace(date)
	return date, nil
}

func extractExpireDate(publishExpire string) (string, error) {
	parts := strings.Fields(publishExpire)
	var expire string
	if len(parts) >= 8 {
		for i, v := range parts {
			if  i > 8 {
				expire += v + " "
			}
		}
	} else {
		fmt.Println("Unable to extract expire date.")
		return  "", errors.New("unable to extract expire date") 
	}
	return expire, nil
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

func extractForecastHtml(doc *goquery.Document) ([]string) {
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

func extractForecastStr(forecastArray []string) string{
	var forecast string
	var parts string
	fmt.Printf("forecast Arr is %v and its length is %v \n", forecastArray, len(forecastArray))
	u := false
	c := false

	for _, val := range forecastArray{
		parts += val + " "
	}
	fmt.Println("parts:", parts)
	partsArr := strings.Fields(parts)
	for index, val := range partsArr{
		if(val == "Until"){
			u = true
		} 
		if(u && !c){
			forecast += val + " "
		}
		if (val == "CFS" && index == len(parts) - 1){
			c = true
		}
	}
	forecast = strings.TrimSpace(forecast)

	 // Length of the string
	 length := len(forecast)

	 // Define the number of characters to remove from the end
	 charactersToRemove := 4 // Length of "</b>"

	 // Check if the string is long enough to contain the characters to remove
	 if length >= charactersToRemove {
			 // Trim the string to remove the last characters
			 trimBrk := forecast[:length-charactersToRemove]

			 // Print the trimmed string
			 fmt.Println("Original string:", forecast)
			 fmt.Println("Trimmed string:", trimBrk)
			 forecast = trimBrk
	 } else {
			 // Handle the case where the string is shorter than the characters to remove
			 fmt.Println("The string is too short to remove characters.")
	 }

	return forecast
}