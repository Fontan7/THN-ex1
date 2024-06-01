package api

import (
	"THN-ex1/types"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestFindMetrics(t *testing.T) {
	// Initialize the ReqIPs structure
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
	}

	// Test case: IP is found
	metrics, err := findMetrics("192.168.1.1", reqIPs)
	assert.NoError(t, err)
	assert.Len(t, metrics, 1)
	assert.Equal(t, "192.168.1.1", metrics[0].IP)

	// Test case: IP is not found
	metrics, err = findMetrics("192.168.1.2", reqIPs)
	assert.Error(t, err)
	assert.Nil(t, metrics)
	assert.Equal(t, "ip not found in metrics", err.Error())
}

func TestExtractIp(t *testing.T) {
	// Initialize Gin router
	router := gin.New()

	// Define a test route
	router.GET("/metrics/:ip", func(c *gin.Context) {
		ip, err := extractIp(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"ip": ip})
	})

	// Test case: Valid IP
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/metrics/192.168.1.1", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"ip": "192.168.1.1"}`, w.Body.String())

	// Test case: Invalid IP
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/metrics/invalid_ip", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.JSONEq(t, `{"error": "invalid IP address format"}`, w.Body.String())
}
