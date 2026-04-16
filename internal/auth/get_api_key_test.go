package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAuthKey(t *testing.T) {
	req, _ := http.NewRequest("GET", "https://example.com", nil)
	wantAuthKey := "your_token_here"
	req.Header.Set("Authorization", "ApiKey "+wantAuthKey)
	authKey, errorMessage := GetAPIKey(req.Header)
	if !reflect.DeepEqual(authKey, wantAuthKey) {
		t.Fatalf("Expected: %v, got: %v", wantAuthKey, authKey)
	}
	if !reflect.DeepEqual(errorMessage, nil) {
		t.Fatalf("Expected: %v, got: %v", nil, errorMessage)
	}
}
