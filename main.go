package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"github.com/leegeobuk/jobscraper/scraper"
	"github.com/leegeobuk/jobscraper/util"
)

const fileName = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("index.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(fileName)
	query := strings.ToLower(util.TrimAllspaces(c.FormValue("query")))
	scraper.Scrape(query)
	return c.Attachment(fileName, fileName)
}