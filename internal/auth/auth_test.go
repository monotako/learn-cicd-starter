package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKeyPositiveTest(t *testing.T) {
	want := "brothermanbill"
	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	req.Header.Add("Authorization", "ApiKey "+want)

	if got, err := GetAPIKey(req.Header); err != nil || got != want {
		t.Fatalf("expected %v got %v", want, got)
	}

}

func TestGetAPIKeyNoAuthHeader(t *testing.T) {
	var emptyHeader http.Header
	want := ""

	if got, err := GetAPIKey(emptyHeader); err != ErrNoAuthHeaderIncluded && got != want {
		t.Fatalf("Expected empty string and ErrNoAuthHeaderIncluded")
	}

}

func TestGetAPIKeyAuthHeaderMalformed(t *testing.T) {
	want := ""

	req, _ := http.NewRequest("GET", "https://www.google.com", nil)
	req.Header.Add("Authorization", "bad stuff")

	if got, err := GetAPIKey(req.Header); err != errors.New("malformed authorization header") && got != want {
		t.Fatalf("Expected empty string and malformed authorization error")
	}

}
