package credentials

import (
	"context"
	"testing"
	"time"
)

func TestIAMToken_NoExpiry(t *testing.T) {
	creds := IAMToken("static-token")
	tok, err := creds.IAMToken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Token != "static-token" {
		t.Errorf("Token = %q, want %q", tok.Token, "static-token")
	}
	if !tok.ExpiresAt.IsZero() {
		t.Errorf("ExpiresAt = %v, want zero (no expiry tracked for bare static tokens)", tok.ExpiresAt)
	}
}

func TestIAMTokenWithExpiry_PreservesExpiry(t *testing.T) {
	want := time.Date(2030, 1, 2, 3, 4, 5, 0, time.UTC)
	creds := IAMTokenWithExpiry("impersonated-token", want)
	tok, err := creds.IAMToken(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tok.Token != "impersonated-token" {
		t.Errorf("Token = %q, want %q", tok.Token, "impersonated-token")
	}
	if !tok.ExpiresAt.Equal(want) {
		t.Errorf("ExpiresAt = %v, want %v", tok.ExpiresAt, want)
	}
}
