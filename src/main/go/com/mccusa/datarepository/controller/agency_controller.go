package controller

import (
	"github.com/gin-gonic/gin"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/dto"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
)

// RegisterClientRoutes registers client endpoints under /client
func RegisterAgencyRoutes(rg *gin.Engine, svc service.AgencyService) {
	grp := rg.Group("/agency")
	grp.POST("/authenticate", authenticateAgency(svc))
	//grp.PUT("/update", updateClientByEmail(svc))
	//grp.GET("/email/:email", getClientByEmail(svc))
}

func authenticateAgency(svc service.AgencyService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userCredentials dto.UserCredentials
		if err := ctx.ShouldBindJSON(&userCredentials); err != nil {
			ctx.JSON(400, gin.H{"error": "Invalid input to authenticate agency"})
			return
		}
		token, err := svc.AuthenticateAgency(ctx.Request.Context(), userCredentials)
		if err != nil {
			ctx.JSON(401, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, gin.H{"token": token})
	}
}
