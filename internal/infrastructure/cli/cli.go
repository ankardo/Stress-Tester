package cli

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ankardo/Stress-Tester/config/logger"
	"go.uber.org/zap"
)

type Args struct {
	URL         string
	Requests    int
	Concurrency int
}

func ParseArguments() (Args, error) {
	fs := flag.NewFlagSet("args", flag.ExitOnError)

	urlFlag := fs.String("url", "", "Target URL for the stress test (required)")
	requestsFlag := fs.Int("requests", 0, "Number of total requests (default: 100)")
	concurrencyFlag := fs.Int("concurrency", 0, "Number of concurrent requests (default: 10)")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		logger.Error("Failed to parse arguments", err)
		return Args{}, fmt.Errorf("failed to parse arguments: %w", err)
	}

	url := prioritizeFlagOrEnv(*urlFlag, os.Getenv("URL"), "URL")
	if url == "" {
		logger.Error("URL is required", fmt.Errorf("URL is required. Please provide it using the --url flag or the URL environment variable"))
		return Args{}, fmt.Errorf("URL is required. Please provide it using the --url flag or the URL environment variable")
	}

	requests := prioritizeEnvOrDefault(*requestsFlag, os.Getenv("REQUESTS"), 100, "REQUESTS")
	concurrency := prioritizeEnvOrDefault(*concurrencyFlag, os.Getenv("CONCURRENCY"), 10, "CONCURRENCY")

	return Args{
		URL:         url,
		Requests:    requests,
		Concurrency: concurrency,
	}, nil
}

func prioritizeFlagOrEnv(flagValue string, envValue string, fieldName string) string {
	if flagValue != "" {
		logger.Debug(fmt.Sprintf("%s value set from flag", fieldName), zap.String("value", flagValue))
		return flagValue
	}
	if envValue != "" {
		logger.Debug(fmt.Sprintf("%s value set from environment variable", fieldName), zap.String("value", envValue))
		return envValue
	}
	logger.Debug(fmt.Sprintf("%s value set to empty string", fieldName))
	return ""
}

func prioritizeEnvOrDefault(flagValue int, envValue string, defaultValue int, fieldName string) int {
	if flagValue > 0 {
		logger.Debug(fmt.Sprintf("%s value set from flag", fieldName), zap.Int("value", flagValue))
		return flagValue
	}
	value, err := strconv.Atoi(envValue)
	if err == nil && value > 0 {
		logger.Debug(fmt.Sprintf("%s value set from environment variable", fieldName), zap.Int("value", value))
		return value
	}
	logger.Debug(fmt.Sprintf("%s value set to default", fieldName), zap.Int("value", defaultValue))
	return defaultValue
}
