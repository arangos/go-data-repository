package config

import (
	"github.com/CodeBullsCode/mccusa-data-repository-go/src/main/go/com/mccusa/datarepository/service"
	"github.com/CodeBullsCode/mccusa-data-repository-go/src/main/go/com/mccusa/datarepository/util"
	"github.com/gin-gonic/gin"
)

// APIPaths are the endpoints to ignore for CSRF-like behavior
var APIPaths = []string{
	"/clients/*any", "/client/*any", "/agencies/*any",
	"/job-applications/*any", "/calendly-webhook/*any",
	"/active-campaign-webhook/*any", "/agency/*any",
	"/docusign-webhook/*any", "/sponsors/*any",
	"/vacancies/*any", "/calendly-events/*any",
	"/sponsor/*any", "/agency-clients/*any",
	"/client-form/*any", "/contracts/*any",
	"/consultants/*any", "/consultant/*any",
	"/contracts-types/*any", "/client-form/*any",
	"/invoices/*any", "/contract/*any",
	"/invoice/*any", "/independent_agency_payments/*any",
	"/contract-type/*any",
}

// SetupSecurity configures authentication and authorization middleware
func SetupSecurity(
	router *gin.Engine,
	userService service.SecurityUserService,
	jwtMiddleware gin.HandlerFunc,
) {
	// Public routes (no auth)
	for _, path := range []string{"/active-campaign-webhook/*any", "/docusign-webhook/*any", "/calendly-webhook/*any"} {
		router.POST(path, func(c *gin.Context) { c.Status(204) })
	}

	// Basic Auth and JWT for all other APIPaths
	auth := router.Group("/")
	auth.Use(util.BasicAuthMiddleware(userService))
	auth.Use(jwtMiddleware)
}
