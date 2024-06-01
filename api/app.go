package api

import (
	"fmt"

	t "THN-ex1/types"
)

type App interface {
	Env() string
	Port() string
	ClientKey() string
	Host() string
	GinMode() string
	IPMetrics() *t.ReqIPs
}

type app struct {
	env       string
	port      string
	clientKey string
	host      string
	ginMode   string
	metrics   *t.ReqIPs
}

func NewApp() (App, error) {
	env := "local" //os.Getenv("ENV")
	if env == "" {
		return nil, fmt.Errorf("ENV variable is empty")
	}
	port := ":8080" //os.Getenv("PORT")
	if port == "" {
		return nil, fmt.Errorf("PORT variable is empty")
	}
	clientKey := "THN_KEY" // os.Getenv("CLIENT_KEY")
	if clientKey == "" {
		return nil, fmt.Errorf("CLIENT_KEY variable is empty")
	}
	host := "127.0.0.1" //os.Getenv("SERVER_HOST")
	if host == "" {
		return nil, fmt.Errorf("SERVER_HOST variable is empty")
	}
	ginMode := "release"
	if ginMode == "" {
		return nil, fmt.Errorf("GIN_MODE variable is empty")
	}

	fmt.Println("Successfully loaded app environment variables")
	return &app{
		env:       env,
		port:      port,
		clientKey: clientKey,
		host:      host,
		ginMode:   ginMode,
		metrics:   &t.ReqIPs{Requests: make(map[string][]t.ReqInfo)},
	}, nil
}

func (a *app) Env() string          { return a.env }
func (a *app) Port() string         { return a.port }
func (a *app) ClientKey() string    { return a.clientKey }
func (a *app) Host() string         { return a.host }
func (a *app) GinMode() string      { return a.ginMode }
func (a *app) IPMetrics() *t.ReqIPs { return a.metrics }
