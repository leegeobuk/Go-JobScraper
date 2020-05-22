package site

import (
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/util"
)

const baseURL string = "https://kr.indeed.com/jobs?q="
const limit string = "&limit=50"

// Site struct
type Site struct {
	url string
}

// New Site
func New(query string) *Site {
	url := baseURL + query + limit
	return &Site{url}
}

// ReadPage and all the jobs in page
func (s *Site) ReadPage(pageNum int, mainC chan<- []*job.Job) {
	pageURL := s.url + "&start=" + strconv.Itoa(pageNum*50)

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
		jobs = append(jobs, <-c)
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

// CountPages in the site
func (s *Site) CountPages() (pages int) {
	resp, err := http.Get(s.url)
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
