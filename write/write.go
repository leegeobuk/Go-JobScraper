package write

import (
	"encoding/csv"
	"os"

	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/util"
)

// Jobs writes all jobs to a csv file
func Jobs(jobs []*job.Job) {
	file, err := os.Create("jobs.csv")
	util.CheckErr(err)

	w := csv.NewWriter(file)
	// defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	util.CheckErr(wErr)

	c := make(chan []string)
	for _, job := range jobs {
		go parseJob(c, job)
	}

	for range jobs {
		wErr := w.Write(<-c)
		util.CheckErr(wErr)
	}
}

func parseJob(c chan<- []string, job *job.Job) {
	jobSlice := job.ToSlice()
	c <- jobSlice
}
