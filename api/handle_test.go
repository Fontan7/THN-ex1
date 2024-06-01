package api

import (
	"THN-ex1/types"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestHandleHealthCheck(t *testing.T) {
	router := gin.New()
	router.GET("/health", handleHealthCheck)

	// Create a test server
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)

	// Perform the request
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	expectedResponse := `{"status":"ok","code":200}`
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func TestHandleGetFeature(t *testing.T) {
	router := gin.New()
	router.GET("/feature", handleGetFeature)

	// Create a test server
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/feature", nil)
	req.RemoteAddr = "192.168.1.1:12345"
	req.Header.Set("User-Agent", "test-agent")

	// Perform the request
	router.ServeHTTP(w, req)

	// Check the response status code
	assert.Equal(t, http.StatusOK, w.Code)

	// Check the response body
	expectedResponse := types.GetFeatureResponse{
		Code:     http.StatusOK,
		Headers:  req.Header,
		Response: "Hello THN backenders ☺ your IP looks like this: " + req.RemoteAddr,
	}
	assert.JSONEq(t, `{"code": 200, "headers": {"User-Agent": ["test-agent"]}, "response": "Hello THN backenders ☺ your IP looks like this: 192.168.1.1:12345"}`, w.Body.String())

	// Verify that the headers in the response match the headers in the request
	var jsonResponse types.GetFeatureResponse
	err := json.Unmarshal(w.Body.Bytes(), &jsonResponse)
	assert.NoError(t, err)
	assert.Equal(t, expectedResponse.Headers, jsonResponse.Headers)
	assert.Equal(t, expectedResponse.Response, jsonResponse.Response)
}


type MockApp struct{ 
	env       string
	port      string
	clientKey string
	host      string
	ginMode   string
	metrics *types.ReqIPs
 }

func (m *MockApp) Env() string          { return m.env }
func (m *MockApp) Port() string         { return m.port }
func (m *MockApp) ClientKey() string    { return m.clientKey }
func (m *MockApp) Host() string         { return m.host }
func (m *MockApp) GinMode() string      { return m.ginMode }
func (m *MockApp) IPMetrics() *types.ReqIPs { return m.metrics }

func TestHandleGetMetrics(t *testing.T) {
	// Initialize the ReqIPs structure with some data
	reqIPs := &types.ReqIPs{
		Requests: map[string][]types.ReqInfo{
			"192.168.1.1": {
				{
					IP:   "192.168.1.1",
					Url:  "/test",
					Time: "2024-06-01T12:00:00Z",
					Headers: map[string][]string{
						"User-Agent": {"test-agent"},
					},
				},
			},
		},
		RWMutex: sync.RWMutex{},
	}

	// Create a mock app
	app := &MockApp{metrics: reqIPs}

	// Initialize Gin router
	router := gin.New()
	router.GET("/metrics/:ip", func(c *gin.Context) {
		handleGetMetrics(c, app)
	})

	tests := []struct {
		name           string
		ip             string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid IP",
			ip:             "192.168.1.1",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"response":{"amount":1,"ip_metrics":[{"ip":"192.168.1.1","url":"/test","time":"2024-06-01T12:00:00Z","headers":{"User-Agent":["test-agent"]}}]}}`,
		},
		{
			name:           "IP Not Found",
			ip:             "192.168.1.2",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"ip not found in metrics"}`,
		},
		{
			name:           "Invalid IP Format",
			ip:             "invalid_ip",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid IP address format"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a test server
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/metrics/"+tt.ip, nil)

			// Perform the request
			router.ServeHTTP(w, req)

			// Check the response status code
			assert.Equal(t, tt.expectedStatus, w.Code)

			// Check the response body
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
