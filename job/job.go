package job

const viewURL string = "https://kr.indeed.com/viewjob?jk="

// Job information
type Job struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

// New creates Job
func New(id, title, location, salary, summary string) *Job {
	job := Job{id, title, location, salary, summary}
	return &job
}

// ToSlice makes a slice with attributes of Job
func (j *Job) ToSlice() []string {
	return []string{viewURL + j.id, j.title, j.location, j.salary, j.summary}
}
