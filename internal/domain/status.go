package domain

import "sync"

type Status struct {
	TotalRequests      int
	SuccessfulRequests int
	FailedRequests     map[int]int
	StatusCodeCounts   map[int]int
	mu                 sync.Mutex
}

func (s *Status) IncrementFailedRequest(statusCode int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.FailedRequests[statusCode]++
}

func (s *Status) IncrementStatusCode(statusCode int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.StatusCodeCounts[statusCode]++
}

func (s *Status) IncrementTotalRequests() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.TotalRequests++
}

func (s *Status) Lock() {
	s.mu.Lock()
}

func (s *Status) Unlock() {
	s.mu.Unlock()
}
