package config

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"net/http"
	"net/url"
)

// JwtTokenComponent defines an interface for JWT generation.
type JwtTokenComponent interface {
	GenerateJwtToken() (string, error)
}

// GetDocusignAccessToken obtains a DocuSign OAuth access token using a JWT assertion.
func GetDocusignAccessToken(
	ctx context.Context,
	jwtService JwtTokenComponent,
	client *http.Client,
) (string, error) {
	authURL := viper.GetString("docusign.auth.url")
	jwtToken, err := jwtService.GenerateJwtToken()
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT: %w", err)
	}
	endpoint := fmt.Sprintf(
		"%s?grant_type=urn:ietf:params:oauth:grant-type:jwt-bearer&assertion=%s",
		authURL,
		url.QueryEscape(jwtToken),
	)
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create HTTP request: %w", err)
	}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("HTTP request error: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}
	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("DocuSign OAuth error: status %d, body %s", resp.StatusCode, string(body))
	}
	var payload struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &payload); err != nil {
		return "", fmt.Errorf("failed to parse OAuth response: %w", err)
	}
	return payload.AccessToken, nil
}
