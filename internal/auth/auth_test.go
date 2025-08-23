package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKeyCorrectHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey 123456789")
	header.Add("Test", "test")

	key, err := GetAPIKey(header)
	if err != nil || key != "123456789" {
		t.Errorf("Couldn't get API key, got key=%s err=%v", key, err)
	}
}

func TestGetAPIKeyMalformedHeader(t *testing.T) {
	header := http.Header{}
	header.Add("Authoration", "123456789")
	header.Add("Test", "test")

	key, err := GetAPIKey(header)
	if err == nil || key == "123456789" {
		t.Errorf("Expected error for malformed header, got key=%s err=%v", key, err)
	}
}
