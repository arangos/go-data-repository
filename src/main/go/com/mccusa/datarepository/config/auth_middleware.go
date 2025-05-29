package config

import (
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicAuthMiddleware enforces HTTP Basic Auth using SecurityUserService
func BasicAuthMiddleware(userService service.SecurityUserService) gin.HandlerFunc {
	return func(c *gin.Context) {
		username, password, ok := c.Request.BasicAuth()
		if !ok || !userService.ValidateCredentials(username, password) {
			c.Header("WWW-Authenticate", `Basic realm="Restricted"`)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		c.Next()
	}
}
