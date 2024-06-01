package api

import (
	"THN-ex1/types"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck godoc
//	@Summary		Health check
//	@Description	always returns OK
//	@Tags			health
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		500	{object}	string
//	@Router			/health [get]
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
//	@Summary	Returns happy response and logs the ip
//	@Description
//	@Tags		feature
//	@Produce	json
//	@Success	200	{object}	types.GetFeatureResponse
//	@Failure	500	{object}	string
//	@Router		/v1/public/feature [get]
func handleGetFeature(c *gin.Context) {
	headers := c.Request.Header

	c.JSON(http.StatusOK, types.GetFeatureResponse{
		Code:     http.StatusOK,
		Headers:  headers,
		Response: "Hello THN backenders ☺ your IP looks like this: " + c.Request.RemoteAddr,
	})
}

// GetMetrics godoc
//	@Summary		Returns matching metrics for the given IP
//	@Description	Returns the number of metrics that match the given IP parameter
//	@Tags			metrics
//	@Produce		json
//	@Param			ip		path		string	true	"IP to search for"
//	@Param			X-Auth	header		string	true	"Authentication token"
//	@Success		200		{object}	types.GetMetricsResponse
//	@Failure		400		{object}	string
//	@Failure		401		{object}	string
//	@Failure		404		{object}	string
//	@Failure		500		{object}	string
//	@Router			/v1/private/metrics/{ip} [get]
func handleGetMetrics(c *gin.Context, app App) {
	findIp, err := extractIp(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ipMetrics, err := findMetrics(findIp, app.IPMetrics())
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := types.GetMetricsResponse{
		Code: http.StatusOK,
	}
	response.Response.Amount = len(ipMetrics)
	response.Response.IpMetrics = ipMetrics

	c.JSON(http.StatusOK, response)
}
