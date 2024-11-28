package dto

type LoadTestResponse struct {
	TotalRequests      int         `json:"total_requests"`
	SuccessfulRequests int         `json:"successful_requests"`
	FailedRequests     map[int]int `json:"failed_requests"`
	Duration           string      `json:"duration"`
}
