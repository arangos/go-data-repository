package config

import (
	"net/http"
	"time"
)

// NewHTTPClient provides a default HTTP client, similar to Spring's RestTemplate
func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 30 * time.Second,
	}
}
