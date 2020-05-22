package scraper

import (
	"fmt"
	"strconv"

	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/site"
	"github.com/leegeobuk/jobscraper/write"
)

// Scrape page by query
func Scrape(query string) {
	jobs := []*job.Job{}
	site := site.New(query)
	pages := site.CountPages()
	mainC := make(chan []*job.Job)
	fmt.Printf("Scraping jobs for %s...\n", query)
	for i := 0; i < pages; i++ {
		go site.ReadPage(i, mainC)
	}
	for i := 0; i < pages; i++ {
		jobs = append(jobs, <-mainC...)
	}
	write.Jobs(jobs)
	fmt.Println("Done writing", strconv.Itoa(len(jobs)))
}
