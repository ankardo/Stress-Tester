package dto

type LoadTestRequest struct {
	URL         string `json:"url"`
	Requests    int    `json:"requests"`
	Concurrency int    `json:"concurrency"`
}
