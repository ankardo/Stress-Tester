package cli

import (
	"os"
	"testing"
)

func TestParseArguments(t *testing.T) {
	tests := []struct {
		name     string
		envVars  map[string]string
		args     []string
		expected Args
		wantErr  bool
	}{
		{
			name: "Valid environment variables",
			envVars: map[string]string{
				"URL":         "http://example.com",
				"REQUESTS":    "10",
				"CONCURRENCY": "2",
			},
			expected: Args{
				URL:         "http://example.com",
				Requests:    10,
				Concurrency: 2,
			},
			wantErr: false,
		},
		{
			name: "Missing URL",
			envVars: map[string]string{
				"REQUESTS":    "100",
				"CONCURRENCY": "10",
			},
			expected: Args{},
			wantErr:  true,
		},
		{
			name: "Flags override environment variables",
			envVars: map[string]string{
				"URL":         "http://example.com",
				"REQUESTS":    "100",
				"CONCURRENCY": "10",
			},
			args: []string{"--url=http://override.com", "--requests=50", "--concurrency=5"},
			expected: Args{
				URL:         "http://override.com",
				Requests:    50,
				Concurrency: 5,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.envVars {
				os.Setenv(k, v)
			}
			defer func() {
				for k := range tt.envVars {
					os.Unsetenv(k)
				}
			}()

			if tt.args != nil {
				os.Args = append([]string{"cmd"}, tt.args...)
			} else {
				os.Args = []string{"cmd"}
			}

			args, err := ParseArguments()
			if (err != nil) != tt.wantErr {
				t.Fatalf("Expected error: %v, got: %v", tt.wantErr, err)
			}

			if !tt.wantErr && args != tt.expected {
				t.Fatalf("Expected args: %+v, got: %+v", tt.expected, args)
			}
		})
	}
}
