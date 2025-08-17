package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey 1234567890")
	apiKey, err := GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "1234567890" {
		t.Fatalf("expected api key to be 1234567890, got %s", apiKey)
	}
}

func TestGetAPIKeyNoHeader(t *testing.T) {
	headers := http.Header{}
	apiKey, err := GetAPIKey(headers)
	if err != ErrNoAuthHeaderIncluded {
		t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
	}
	if apiKey != "" {
		t.Fatalf("expected api key to be empty, got %s", apiKey)
	}
}
