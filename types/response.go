package types

type GetFeatureResponse struct {
	Code     int                 `json:"code"`
	Headers  map[string][]string `json:"headers"`
	Response string              `json:"response"`
}

type GetMetricsResponse struct {
	Code     int `json:"code"`
	Response struct {
		Amount    int       `json:"amount"`
		IpMetrics []ReqInfo `json:"ip_metrics"`
	} `json:"response"`
}

type ErrorResponse struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
