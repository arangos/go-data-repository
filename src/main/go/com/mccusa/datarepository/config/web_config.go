package config

import (
	"github.com/gin-gonic/gin"
)

// SetupWebConfig registers route-specific middleware, e.g., LoggingMiddleware for /client-form/**
func SetupWebConfig(router *gin.Engine) {
	// Only log payloads on /client-form/**
	clientForm := router.Group("/client-form")
	clientForm.Use(LoggingMiddleware())
}
