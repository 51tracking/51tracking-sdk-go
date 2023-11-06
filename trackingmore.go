package tracking51

import (
	"crypto/tls"
	"errors"
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	httpClient *http.Client
}

// NewClient returns the 51Tracking client
func NewClient(apiKey string) (*Client, error) {
	if apiKey == "" {
		return nil, errors.New(ErrEmptyAPIKey)
	}

	return &Client{
		apiKey: apiKey,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			},
		},
	}, nil
}
