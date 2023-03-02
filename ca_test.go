package gosunetca

import (
	"testing"
)

func mockClient(t *testing.T, severURL string) *Client {
	cfg := Config{
		ServerURL: severURL,
		Token:     "test-token",
	}
	if err := Check(cfg); err != nil {
		t.Errorf("mockClient: %v", err)
		t.FailNow()
	}
	c, err := New(cfg)
	if err != nil {
		t.Errorf("mockClient: %v", err)
		t.FailNow()
	}

	return c
}
