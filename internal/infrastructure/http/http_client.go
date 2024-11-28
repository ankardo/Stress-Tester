package httpclient

import (
	"net/http"
	"time"

	"github.com/ankardo/Stress-Tester/config/logger"
)

type Client struct {
	httpClient *http.Client
}

func NewClient(timeout time.Duration) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
	}
}

func SendRequest(url, method string, headers map[string]string) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		logger.Debug("Error creating request: " + err.Error())
		return nil, err
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	logger.Debug("Sending request to " + url + " with method " + method)
	resp, err := client.Do(req)
	if err != nil {
		logger.Debug("Request failed: " + err.Error())
		return nil, err
	}

	logger.Debug("Received response with status: " + http.StatusText(resp.StatusCode))
	return resp, nil
}
