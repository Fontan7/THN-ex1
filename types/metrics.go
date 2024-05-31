package types

import "sync"

type ReqIPs struct {
	sync.RWMutex
	Requests map[string][]ReqInfo
}

type ReqInfo struct {
	IP      string              `json:"ip"`
	Time    string              `json:"time"`
	Headers map[string][]string `json:"headers"`
}
