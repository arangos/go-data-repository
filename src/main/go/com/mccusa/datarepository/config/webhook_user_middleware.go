package config

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// WebhookUserMiddleware authenticates webhook POSTs by path prefix.
//
// For POSTs to /active-campaign-webhook/* it will set the “user” to “ActiveCampaign”,
// for /docusign-webhook/* to “Docusign”, and for /calendly-webhook/* to “Calendly”.
//
// Downstream you can read it via: user, _ := c.Get("username")
func WebhookUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == http.MethodPost {
			path := c.FullPath()
			switch {
			case strings.HasPrefix(path, "/active-campaign-webhook"):
				c.Set("username", "ActiveCampaign")
				c.Next()
				return
			case strings.HasPrefix(path, "/docusign-webhook"):
				c.Set("username", "Docusign")
				c.Next()
				return
			case strings.HasPrefix(path, "/calendly-webhook"):
				c.Set("username", "Calendly")
				c.Next()
				return
			}
		}
		c.Next()
	}
}
