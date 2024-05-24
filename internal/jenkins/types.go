package jenkins

type Job struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Color string `json:"color"`
}

type JobResponse struct {
	Jobs []Job `json:"jobs"`
}
