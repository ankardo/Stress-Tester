package httpclient

import (
	"net/http"
	"testing"
)

func TestSendRequest(t *testing.T) {
	tests := []struct {
		name        string
		url         string
		method      string
		headers     map[string]string
		expectError bool
	}{
		{
			name:        "Valid GET request",
			url:         "http://example.com",
			method:      "GET",
			headers:     nil,
			expectError: false,
		},
		{
			name:        "Invalid URL",
			url:         "://invalid-url",
			method:      "GET",
			headers:     nil,
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			response, err := SendRequest(tt.url, tt.method, tt.headers)
			if (err != nil) != tt.expectError {
				t.Errorf("Expected error: %v, got: %v", tt.expectError, err)
			}
			if response != nil && response.StatusCode != http.StatusOK && !tt.expectError {
				t.Errorf("Expected status code 200, got: %v", response.StatusCode)
			}
		})
	}
}
