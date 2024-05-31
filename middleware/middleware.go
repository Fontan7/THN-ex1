package middleware

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CorsConfig() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTION")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, Authorization, X-API-Key")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
		//fmt.Println(c.Request.Method)
		//fmt.Println(c.Request.Response)
		//fmt.Println(c.Request.WithContext(c))
	}
}

func CheckAPIKey(clientKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if clientKey != c.GetHeader("X-API-Key") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, "invalid api key")
		}

		c.Next()
	}
}

func ErrorManager() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next() // Continue processing the request

		// If errors are present after processing
		if len(c.Errors) > 0 {
			var errors []map[string]interface{}

			for _, err := range c.Errors {
				// Log each error
				log.Printf("Time: %s, URL: %s, Error: %s, HTTP Code: %d\n",
					time.Now().Format(time.RFC3339),
					c.Request.URL.String(),
					err.Error(),
					c.Writer.Status(),
				)
				// Add the error to the slice for JSON response
				errors = append(errors, map[string]interface{}{
					"time":       time.Now().Format(time.RFC3339),
					"url":        c.Request.URL.String(),
					"error":      err.Error(),
					"httpStatus": c.Writer.Status(),
				})
			}

			// Respond with all logged errors in JSON format
			c.AbortWithStatusJSON(c.Writer.Status(), gin.H{
				"errors": errors,
			})
		}
	}
}