package api

import (
	"THN-ex1/types"
	"errors"
	"regexp"

	"github.com/gin-gonic/gin"
)

func findMetrics(findIp string, metrics *types.ReqIPs) ([]types.ReqInfo, error) {
	metrics.RLock()
	defer metrics.RUnlock()

	reqs, ok := metrics.Requests[findIp]
	if !ok {
		return nil, errors.New("ip not found in metrics")
	}

	return reqs, nil
}

func extractIp(c *gin.Context) (string, error) {
	ip := c.Param("ip")
	ipRegex := `^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$`
	validIP := regexp.MustCompile(ipRegex).MatchString(ip)
	if !validIP {
		return "", errors.New("invalid IP address format")
	}

	return ip, nil
}
