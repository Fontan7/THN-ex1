package api

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockApp is a mock implementation of the App interface
type MockApp struct {
	mock.Mock
}

func (m *MockApp) Env() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockApp) Port() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockApp) ClientKey() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockApp) Host() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockApp) GinMode() string {
	args := m.Called()
	return args.String(0)
}

func (m *MockApp) AddReqIp(ip string) {
	m.Called(ip)
}

func (m *MockApp) GetMetrics(ip string) (int, error) {
	args := m.Called(ip)
	return args.Int(0), args.Error(1)
}

func TestHandleGetMetrics(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		ip             string
		mockAmount     int
		mockError      error
		expectedStatus int
		expectedBody   string
		setExpectation bool
	}{
		{
			name:           "Invalid IP format",
			ip:             "invalid_ip",
			mockAmount:     0,
			mockError:      nil,
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid IP address format"}`,
			setExpectation: false,
		},
		{
			name:           "IP not found",
			ip:             "192.168.1.1",
			mockAmount:     0,
			mockError:      errors.New("IP not found"),
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"error":"IP not found"}`,
			setExpectation: true,
		},
		{
			name:           "Successful retrieval",
			ip:             "192.168.1.1",
			mockAmount:     5,
			mockError:      nil,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":200,"response":{"ip":"192.168.1.1","amount":5}}`,
			setExpectation: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockApp := new(MockApp)
			if tt.setExpectation {
				mockApp.On("GetMetrics", tt.ip).Return(tt.mockAmount, tt.mockError)
			}

			router := gin.Default()
			router.GET("/metrics/:ip", func(c *gin.Context) {
				handleGetMetrics(c, mockApp)
			})

			req := httptest.NewRequest("GET", "/metrics/"+tt.ip, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())

			if tt.setExpectation {
				mockApp.AssertExpectations(t)
			}
		})
	}
}
