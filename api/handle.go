package api

import (
	"fmt"
	"net"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
// @Summary		Health check
// @Description	always returns OK
// @Tags			health
// @Produce		json
// @Success		200	{object}	string
// @Failure		500
// @Router			/health [get]
func handleHealthCheck(c *gin.Context) {
	// neat test of error handling
	/*
		if true {
			c.Error(errors.New("non fatal error"))
			c.AbortWithError(http.StatusInternalServerError, errors.New("a fatal error"))
			return
		}
	*/

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"code":   http.StatusOK,
	})
}

// GetFeature godoc
// @Summary      Returns happy response and logs the ip
// @Description
// @Tags         feature
// @Produce      json
// @Success      200  {object}  string
// @Failure      500  {object}  string
// @Router       /v1/feature [get]
func handleGetFeature(c *gin.Context, app App) {
	ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, ErrorResponse{
			Code:  http.StatusInternalServerError,
			Error: fmt.Sprintf("could not get ip: %v", err),
		})
		return
	}

	app.AddReqIp(ip)
	t := time.Now().Format(time.RFC3339)
	headers := make(map[string]string)
	for key, values := range c.Request.Header {
		headers[key] = values[0]
	}

	c.JSON(http.StatusOK, StandardResponse{
		Code:     http.StatusOK,
		Headers:  headers,
		Response: "Hello THN backenders ☺ your IP is: " + ip + " at time: " + t,
	})
}

// GetMetrics godoc
// @Summary      Returns matching metrics for the given IP
// @Description  Returns the number of metrics that match the given IP parameter
// @Tags         metrics
// @Produce      json
// @Param        ip    path     string  true  "IP to search for"
// @Param        X-Auth header   string  true  "Authentication token"
// @Success      200   {object}  GetMetricsResponse
// @Failure      400   {object}  string
// @Failure      404   {object}  string
// @Failure      500   {object}  string
// @Router       /v1/metrics/{ip} [get]
func handleGetMetrics(c *gin.Context, app App) {
	type Res struct {
		Ip     string `json:"ip"`
		Amount int    `json:"amount"`
	}
	findIp := c.Param("ip")
	ipRegex := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	validIP := regexp.MustCompile(ipRegex).MatchString(findIp)
	if !validIP {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid IP address format"})
		return
	}

	amount, err := app.GetMetrics(findIp)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, GetMetricsResponse{
		Code: http.StatusOK,
		Response: Res{
			Ip:     findIp,
			Amount: amount,
		},
	})
}
