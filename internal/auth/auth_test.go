package auth_test

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

func TestGetAPIKey_Success(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "ApiKey test123")

	apiKey, err := auth.GetAPIKey(headers)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if apiKey != "test123" {
		t.Fatalf("expected 'test123', got '%s'", apiKey)
	}
}

func TestGetAPIKey_MissingAuthHeader(t *testing.T) {
	headers := http.Header{}

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
}

func TestGetAPIKey_InvalidPrefix(t *testing.T) {
	headers := http.Header{}
	headers.Set("Authorization", "Bearer sometoken")

	_, err := auth.GetAPIKey(headers)
	if err == nil {
		t.Fatal("expected an error due to invalid prefix, got nil")
	}
}
