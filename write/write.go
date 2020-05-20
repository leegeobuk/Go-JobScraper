package write

import (
	"encoding/csv"
	"os"

	"github.com/leegeobuk/jobscraper/job"
	"github.com/leegeobuk/jobscraper/util"
)

// Jobs writes all jobs on a csv file
func Jobs(jobs []*job.Job) {
	file, err := os.Create("jobs.csv")
	util.CheckErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(headers)
	util.CheckErr(wErr)

	for _, job := range jobs {
		writeOnFile(w, job)
	}
}

func writeOnFile(w *csv.Writer, job *job.Job) {
	jobSlice := job.ToSlice()
	wErr := w.Write(jobSlice)
	util.CheckErr(wErr)
}
