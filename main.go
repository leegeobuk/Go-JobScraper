package main

import (
	"fmt"
	"strconv"

	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/read"
	"github.com/leegeobuk/jobscraper/write"
)

func main() {
	jobs := []*job.Job{}
	pages := read.GetPageCounts()
	mainC := make(chan []*job.Job)
	for i := 0; i < pages; i++ {
		go read.Page(i, mainC)
	}
	for i := 0; i < pages; i++ {
		jobs = append(jobs, <-mainC...)
	}
	write.Jobs(jobs)
	fmt.Println("Finished writing", strconv.Itoa(len(jobs)))
}
