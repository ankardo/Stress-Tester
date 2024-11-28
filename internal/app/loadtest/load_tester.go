package loadtest

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ankardo/Stress-Tester/config/logger"
	"github.com/ankardo/Stress-Tester/internal/domain"
	"github.com/ankardo/Stress-Tester/internal/infrastructure/cli"
	"go.uber.org/zap"
)

func RunLoadTester(args cli.Args) *domain.Status {
	status := &domain.Status{
		FailedRequests:   make(map[int]int),
		StatusCodeCounts: make(map[int]int),
	}

	var wg sync.WaitGroup
	semaphore := make(chan struct{}, args.Concurrency)
	progressChannel := make(chan int, args.Requests)

	logger.Debug("Initializing load test",
		zap.String("URL", args.URL),
		zap.Int("Requests", args.Requests),
		zap.Int("Concurrency", args.Concurrency),
	)

	go displayProgress(progressChannel, args.Requests)

	for i := 0; i < args.Requests; i++ {
		wg.Add(1)
		semaphore <- struct{}{}

		go func(requestNumber int) {
			defer wg.Done()
			defer func() { <-semaphore }()
			progressChannel <- requestNumber

			resp, err := http.Get(args.URL)
			status.IncrementTotalRequests()

			if err != nil {
				logger.Debug("Request failed",
					zap.Int("RequestNumber", requestNumber),
					zap.Error(err),
				)
				status.IncrementFailedRequest(500)
				return
			}
			defer resp.Body.Close()

			logger.Debug("Request succeeded",
				zap.Int("RequestNumber", requestNumber),
				zap.Int("StatusCode", resp.StatusCode),
			)
			status.IncrementStatusCode(resp.StatusCode)
		}(i + 1)
	}

	wg.Wait()
	close(progressChannel)

	logger.Debug("Load test completed",
		zap.Int("TotalRequests", status.TotalRequests),
		zap.Any("FailedRequests", status.FailedRequests),
		zap.Any("StatusCodeCounts", status.StatusCodeCounts),
	)

	return status
}

func displayProgress(progressChannel <-chan int, totalRequests int) {
	startTime := time.Now()
	completedRequests := 0

	for range progressChannel {
		completedRequests++
		elapsed := time.Since(startTime).Seconds()
		progress := float64(completedRequests) / float64(totalRequests) * 100
		fmt.Printf("\rProgress: %d/%d requests completed (%.2f%%) - Elapsed: %.2fs",
			completedRequests, totalRequests, progress, elapsed)
	}
	fmt.Println()
}
