package reports

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/ankardo/Stress-Tester/internal/domain"
	"github.com/stretchr/testify/assert"
)

func TestGenerateReport(t *testing.T) {
	tests := []struct {
		name           string
		status         *domain.Status
		expectedOutput string
	}{
		{
			name: "Valid report with 200 and 500",
			status: &domain.Status{
				TotalRequests:    100,
				StatusCodeCounts: map[int]int{200: 90, 500: 10},
			},
			expectedOutput: `
                      METRIC                       | VALUE
        ----------------------------------------|--------
          Successful Requests (200)              |    90
          Internal Server Errors Requests (500)  |    10
          Total Requests                         |   100
          Total Duration (s)                     |  2.00
`,
		},
		{
			name: "Valid report with 404 and other codes",
			status: &domain.Status{
				TotalRequests:    120,
				StatusCodeCounts: map[int]int{200: 100, 404: 15, 418: 5},
			},
			expectedOutput: `
                     METRIC             | VALUE
        --------------------------------|---------
          Successful Requests (200)     |   100
          Page Not Found Requests (404) |    15
          Other HTTP Responses          | 418: 5
          Total Requests                |   120
          Total Duration (s)            |  2.00
`,
		},
		{
			name: "Valid report with only 200",
			status: &domain.Status{
				TotalRequests:    50,
				StatusCodeCounts: map[int]int{200: 50},
			},
			expectedOutput: `
                   METRIC           | VALUE
        ----------------------------|--------
          Successful Requests (200) |    50
          Total Requests            |    50
          Total Duration (s)        |  2.00
`,
		},
		{
			name: "Valid report with no successful requests",
			status: &domain.Status{
				TotalRequests:    30,
				StatusCodeCounts: map[int]int{404: 20, 500: 10},
			},
			expectedOutput: `
                      METRIC                       | VALUE
        ----------------------------------------|--------
          Page Not Found Requests (404)          |    20
          Internal Server Errors Requests (500)  |    10
          Total Requests                         |    30
          Total Duration (s)                     |  2.00
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			done := make(chan string)
			go func() {
				var buf bytes.Buffer
				io.Copy(&buf, r)
				done <- buf.String()
			}()

			GenerateReport(2*time.Second, tt.status)

			w.Close()
			os.Stdout = oldStdout
			output := <-done

			t.Logf("Generated Report:\n%s", output)

			assert.Equal(t, normalizeString(tt.expectedOutput), normalizeString(output))
		})
	}
}

func normalizeString(input string) string {
	trimmed := strings.TrimSpace(input)
	unified := strings.ReplaceAll(trimmed, "\r\n", "\n")
	return strings.ReplaceAll(unified, " ", "")
}
