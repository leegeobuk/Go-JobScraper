package read

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/util"
)

const baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

// Page gets all jobs in the page
func Page(pageNum int, mainC chan<- []*job.Job) {
	pageURL := baseURL + "&start=" + strconv.Itoa(pageNum*50)
	fmt.Println("Requesting:", pageURL)

	resp, err := http.Get(pageURL)
	util.CheckErr(err)
	util.CheckStatusCode(resp)

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	defer resp.Body.Close()
	util.CheckErr(err)

	mainC <- addJobs(doc)
}

func addJobs(doc *goquery.Document) []*job.Job {
	jobs := []*job.Job{}
	c := make(chan *job.Job)
	jobCards := extractJobCards(doc, c)
	for i := 0; i < jobCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}
	return jobs
}

func extractJobCards(doc *goquery.Document, c chan<- *job.Job) *goquery.Selection {
	jobCards := doc.Find(".jobsearch-SerpJobCard")
	jobCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
	})
	return jobCards
}

func extractJob(card *goquery.Selection, c chan<- *job.Job) {
	id, _ := card.Attr("data-jk")
	title := util.TrimAllspaces(card.Find(".title>a").Text())
	location := util.TrimAllspaces(card.Find(".sjcl").Text())
	salary := util.TrimAllspaces(card.Find(".salarySnippet").Text())
	summary := util.TrimAllspaces(card.Find(".summary").Text())
	c <- job.New(id, title, location, salary, summary)
}

// GetPageCounts gets total number of pages
func GetPageCounts() (pages int) {
	resp, err := http.Get(baseURL)
	util.CheckErr(err)
	util.CheckStatusCode(resp)

	defer resp.Body.Close()
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	util.CheckErr(err)
	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return
}
