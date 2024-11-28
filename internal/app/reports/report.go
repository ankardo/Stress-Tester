package reports

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/ankardo/Stress-Tester/config/logger"
	"github.com/ankardo/Stress-Tester/internal/domain"
	"github.com/olekukonko/tablewriter"
	"go.uber.org/zap"
)

func GenerateReport(duration time.Duration, status *domain.Status) {
	status.Lock()
	defer status.Unlock()

	logger.Debug("Starting report generation",
		zap.Int("TotalRequests", status.TotalRequests),
		zap.Duration("Duration", duration),
		zap.Any("StatusCodeCounts", status.StatusCodeCounts),
	)

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"METRIC", "VALUE"})
	table.SetAutoWrapText(false)
	table.SetColWidth(50)
	table.SetAutoFormatHeaders(false)
	table.SetBorder(false)
	table.SetCenterSeparator("|")

	if count := status.StatusCodeCounts[200]; count > 0 {
		table.Append([]string{"Successful Requests (200)", fmt.Sprintf("%d", count)})
	}
	if count := status.StatusCodeCounts[404]; count > 0 {
		table.Append([]string{"Page Not Found Requests (404)", fmt.Sprintf("%d", count)})
	}
	if count := status.StatusCodeCounts[500]; count > 0 {
		table.Append([]string{"Internal Server Errors Requests (500)", fmt.Sprintf("%d", count)})
	}

	var builder strings.Builder
	for code, count := range status.StatusCodeCounts {
		if code != 200 && code != 404 && code != 500 {
			builder.WriteString(fmt.Sprintf("%d: %d, ", code, count))
		}
	}
	otherSummary := strings.TrimSuffix(builder.String(), ", ")
	if otherSummary != "" {
		table.Append([]string{"Other HTTP Responses", otherSummary})
	}

	table.Append([]string{"Total Requests", fmt.Sprintf("%d", status.TotalRequests)})
	table.Append([]string{"Total Duration (s)", fmt.Sprintf("%.2f", duration.Seconds())})

	table.SetRowSeparator("-")
	table.SetHeaderLine(true)
	table.Render()
}
