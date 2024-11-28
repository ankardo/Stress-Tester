package loadtest

import (
	"testing"

	"github.com/ankardo/Stress-Tester/internal/infrastructure/cli"
)

func TestRunLoadTester(t *testing.T) {
	tests := []struct {
		name        string
		args        cli.Args
		expectError bool
	}{
		{
			name: "Valid load test",
			args: cli.Args{
				URL:         "http://example.com",
				Requests:    100,
				Concurrency: 10,
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := RunLoadTester(tt.args)
			if (status.TotalRequests == 0) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, status)
			}
		})
	}
}
