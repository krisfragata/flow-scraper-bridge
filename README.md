# River Flow Scraper & Database Updater

![Go Version](https://img.shields.io/badge/Go-1.22+-00ADD8?logo=go)
![License](https://img.shields.io/badge/license-MIT-green)
![Status](https://img.shields.io/badge/status-active-brightgreen)

A Go application that scrapes daily **river flow data** from [H2OLine](http://www.h2oline.com/srcs/255122.html) and automatically updates a **PostgreSQL (Supabase)** database.  
It detects scheduled release days, prevents duplicate entries, and ensures the latest water flow information is always recorded.

## Overview

This tool automates data collection from H2OLine and updates a Supabase database with the latest:

- **Flow rate (CFS)**
- **Posting time**
- **Forecast summary**
- **Expiration date**
- **Whether today is a release day**

It can be run manually or on a schedule (e.g., via cron).

## Features

**Web Scraping** — Uses [`colly`](https://github.com/gocolly/colly) and [`goquery`](https://github.com/PuerkitoBio/goquery)  
**Smart Parsing** — Extracts clean text from messy HTML  
**Release Detection** — Flags if the current day is a planned release  
**Database Integration** — Inserts or updates data in Supabase PostgreSQL  
**Error Handling** — Logs failed scrapes and prevents duplicate writes  

## Dependencies

| Library | Purpose |
|----------|----------|
| [`github.com/gocolly/colly`](https://github.com/gocolly/colly) | Web scraping |
| [`github.com/PuerkitoBio/goquery`](https://github.com/PuerkitoBio/goquery) | HTML parsing |
| [`github.com/jackc/pgx/v4`](https://github.com/jackc/pgx) | PostgreSQL driver |
| [`github.com/joho/godotenv`](https://github.com/joho/godotenv) | Load environment variables |

## Data Flow

1. **Scrape HTML** using Colly
2. **Parse content** with GoQuery to extract:

   * Publish date
   * Expiration date
   * Flow rate (CFS)
   * Posting time
   * Forecast text
3. **Compare** with most recent database row
4. **Update or insert** based on changes
5. **Mark release days** using predefined map in `isRelease()`


## Future Improvements

* [ ] Split logic into smaller files (`scrape.go`, `db.go`, etc.)
* [ ] Add unit tests for parsing functions
* [ ] Add Slack or email alerts for failed scrapes
* [ ] Support multiple source URLs
* [ ] Use environment variable validation
* [ ] Implement as a Cron job and integrate to larger, What's the Flow application

