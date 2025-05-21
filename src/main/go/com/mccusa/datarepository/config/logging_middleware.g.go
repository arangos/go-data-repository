package config

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

// LoggingMiddleware logs JSON request payloads for POST and PUT methods.
func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		if method == http.MethodPost || method == http.MethodPut {
			// Read the Body
			bodyBytes, err := io.ReadAll(c.Request.Body)
			if err == nil {
				// Restore the Body
				c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
				logrus.Infof("Incoming JSON payload for path: %s, payload: %s", c.Request.RequestURI, string(bodyBytes))
			}
		}
		c.Next()
	}
}
