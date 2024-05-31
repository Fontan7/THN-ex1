package api

import (
	"errors"
	"fmt"
	"os"
	"sync"
)

type App interface {
	Env() string
	Port() string
	ClientKey() string
	Host() string
	GinMode() string
	AddReqIp(string)
	GetMetrics(string) (int, error)
}

type app struct {
	env       string
	port      string
	clientKey string
	host      string
	ginMode   string
	reqIPs    *reqIPs
}

type reqIPs struct {
	sync.RWMutex
	m map[string]int
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

	ginMode := os.Getenv("GIN_MODE")

	fmt.Println("Successfully loaded app environment variables")
	return &app{
		env:       env,
		port:      port,
		clientKey: clientKey,
		host:      host,
		ginMode:   ginMode,
		reqIPs:    &reqIPs{m: make(map[string]int)},
	}, nil
}

func (a *app) Env() string       { return a.env }
func (a *app) Port() string      { return a.port }
func (a *app) ClientKey() string { return a.clientKey }
func (a *app) Host() string      { return a.host }
func (a *app) GinMode() string   { return a.ginMode }
func (a *app) AddReqIp(ip string) {
	a.reqIPs.Lock()
	defer a.reqIPs.Unlock()
	a.reqIPs.m[ip] += 1
}
func (a *app) GetMetrics(match string) (int, error) {
	a.reqIPs.RLock()
	defer a.reqIPs.RUnlock()
	
	amount, ok := a.reqIPs.m[match]
	if !ok {
		return 0, errors.New("no match for ip")
	}

	return amount, nil
}
