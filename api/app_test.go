package api_test

import (
	"THN-ex1/api"

	"testing"
)

func Test_NewAppSuccess(t *testing.T) {
	app, err := api.NewApp()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if app.Env() != "local" {
		t.Errorf("expected env to be 'local', got %s", app.Env())
	}
	if app.Port() != ":8080" {
		t.Errorf("expected port to be ':8080', got %s", app.Port())
	}
	if app.ClientKey() != "THN_KEY" {
		t.Errorf("expected clientKey to be 'THN_KEY', got %s", app.ClientKey())
	}
	if app.Host() != "127.0.0.1" {
		t.Errorf("expected host to be '127.0.0.1', got %s", app.Host())
	}
	if app.GinMode() != "release" {
		t.Errorf("expected gin mode to be 'release', got %s", app.GinMode())
	}
	if app.IPMetrics() == nil {
		t.Errorf("expected ip metrics to be not nil")
	}
}
