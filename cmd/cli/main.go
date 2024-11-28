package main

import (
	"time"

	"github.com/ankardo/Stress-Tester/internal/app/loadtest"
	"github.com/ankardo/Stress-Tester/internal/app/reports"
	"github.com/ankardo/Stress-Tester/internal/infrastructure/cli"
)

func main() {
	args, err := cli.ParseArguments()
	if err != nil {
		panic("Url is a mandatory parameter")
	}
	startTime := time.Now()
	status := loadtest.RunLoadTester(args)
	duration := time.Since(startTime)
	reports.GenerateReport(duration, status)
}
