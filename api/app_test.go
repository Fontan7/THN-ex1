package api_test

import (
	"THN-ex1/api"
	ty "THN-ex1/types"
	"os"
	"testing"
)

func Test_NewAppSuccess(t *testing.T) {
	os.Setenv("ENV", "local")
	os.Setenv("PORT", ":8080")
	os.Setenv("CLIENT_KEY", "THN_KEY")
	os.Setenv("SERVER_HOST", "127.0.0.1")
	os.Setenv("GIN_MODE", "debug")

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
	if app.GinMode() != "debug" {
		t.Errorf("expected ginMode to be 'debug', got %s", app.GinMode())
	}
}

func Test_AddReqIpAndGetIPMetricsLogs(t *testing.T) {
	app, _ := api.NewApp()
	reqInfo := ty.ReqInfo{
		IP:   "127.0.0.1",
		Time: "2023-05-01T12:34:56Z",
		Headers: map[string][]string{
			"User-Agent": {"Mozilla/5.0"},
		},
	}

	app.AddReqIp(reqInfo)

	reqs, err := app.GetIPMetricsLogs("127.0.0.1")
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(reqs) != 1 {
		t.Fatalf("expected 1 request, got %d", len(reqs))
	}

	if !compareReqInfo(reqs[0], reqInfo) {
		t.Errorf("expected request %v, got %v", reqInfo, reqs[0])
	}

	_, err = app.GetIPMetricsLogs("192.168.0.1")
	if err == nil {
		t.Errorf("expected error for non-existing IP, got none")
	}
}

func compareReqInfo(a, b ty.ReqInfo) bool {
	if a.IP != b.IP || a.Time != b.Time {
		return false
	}
	if len(a.Headers) != len(b.Headers) {
		return false
	}
	for k, v := range a.Headers {
		if len(v) != len(b.Headers[k]) {
			return false
		}
		for i := range v {
			if v[i] != b.Headers[k][i] {
				return false
			}
		}
	}
	return true
}