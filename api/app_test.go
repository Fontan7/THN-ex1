package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	// Test successful creation

	app, err := NewApp()
	assert.NoError(t, err)
	assert.NotNil(t, app)
	assert.Equal(t, "local", app.Env())
	assert.Equal(t, ":8080", app.Port())
	assert.Equal(t, "THN_KEY", app.ClientKey())
	assert.Equal(t, "127.0.0.1", app.Host())
}

func TestAppMethods(t *testing.T) {
	app := &app{
		env:       "local",
		port:      ":8080",
		clientKey: "THN_KEY",
		host:      "127.0.0.1",
		ginMode:   "debug",
		reqIPs:    &reqIPs{m: make(map[string]int)},
	}

	assert.Equal(t, "local", app.Env())
	assert.Equal(t, ":8080", app.Port())
	assert.Equal(t, "THN_KEY", app.ClientKey())
	assert.Equal(t, "127.0.0.1", app.Host())
	assert.Equal(t, "debug", app.GinMode())

	app.AddReqIp("192.168.1.1")
	app.AddReqIp("192.168.1.1")
	app.AddReqIp("192.168.1.2")

	amount, err := app.GetMetrics("192.168.1.1")
	assert.NoError(t, err)
	assert.Equal(t, 2, amount)

	amount, err = app.GetMetrics("192.168.1.2")
	assert.NoError(t, err)
	assert.Equal(t, 1, amount)

	amount, err = app.GetMetrics("192.168.1.3")
	assert.Error(t, err)
	assert.Equal(t, 0, amount)
}
