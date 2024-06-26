package middleware

import (
	"THN-ex1/types"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_CorsConfig(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(CorsConfig())
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("OPTIONS", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "*", w.Header().Get("Access-Control-Allow-Origin"))
	assert.Equal(t, "86400", w.Header().Get("Access-Control-Max-Age"))
	assert.Equal(t, "GET, OPTION", w.Header().Get("Access-Control-Allow-Methods"))
	assert.Equal(t, "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-API-Key", w.Header().Get("Access-Control-Allow-Headers"))
	assert.Equal(t, "Content-Length", w.Header().Get("Access-Control-Expose-Headers"))
	assert.Equal(t, "true", w.Header().Get("Access-Control-Allow-Credentials"))
}

func Test_CheckAPIKey(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	clientKey := "valid_api_key"
	router.Use(CheckAPIKey(clientKey))
	router.GET("/test", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	t.Run("valid API key", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-Key", clientKey)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "test", w.Body.String())
	})

	t.Run("invalid API key", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("X-API-Key", "invalid_api_key")
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
		assert.JSONEq(t, `{"code":401,"error":"invalid api key"}`, w.Body.String())
	})
}

func Test_ErrorManager(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.Use(ErrorManager())
	router.GET("/test", func(c *gin.Context) {
		c.Error(gin.Error{
			Err:  errors.New("test error"),
			Type: gin.ErrorTypePublic,
		})
		c.String(http.StatusInternalServerError, "test")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	assert.Contains(t, w.Body.String(), "test error")
	assert.Contains(t, w.Body.String(), "test error")
}

func TestLogIpMetrics(t *testing.T) {
	r := gin.New()
	reqIPs := &types.ReqIPs{
		Requests: make(map[string][]types.ReqInfo),
	}
	r.Use(LogIpMetrics(reqIPs))
	r.GET("/test", func(c *gin.Context) {c.JSON(http.StatusOK, gin.H{"message": "success"})})

	// Create a test server
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/test", nil)
	req.RemoteAddr = "192.168.1.1:12345" // Set the remote address for the request

	// Perform the request
	r.ServeHTTP(w, req)

	// Check the response
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"message": "success"}`, w.Body.String())

	// Check the IP metrics
	reqIPs.Lock()
	defer reqIPs.Unlock()
	assert.Contains(t, reqIPs.Requests, "192.168.1.1")
	assert.Len(t, reqIPs.Requests["192.168.1.1"], 1)
	assert.Equal(t, "/test", reqIPs.Requests["192.168.1.1"][0].Url)
}
