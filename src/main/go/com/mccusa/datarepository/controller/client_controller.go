package controller

import (
	"github.com/gin-gonic/gin"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"net/http"
)

// RegisterClientRoutes registers client endpoints under /client
func RegisterClientRoutes(rg *gin.RouterGroup, svc service.ClientService) {
	grp := rg.Group("/client")
	grp.POST("/create", createAgencyClient(svc))
	grp.PUT("/update", updateClientByEmail(svc))
	grp.GET("/email/:email", getClientByEmail(svc))
}

func createAgencyClient(svc service.ClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var client model.AgencyClient
		if err := ctx.ShouldBindJSON(&client); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err, _ := svc.CreateClient(ctx.Request.Context(), &client); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create client"})
			return
		}
		ctx.JSON(http.StatusOK, client)
	}
}

func updateClientByEmail(svc service.ClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var client model.AgencyClient
		if err := ctx.ShouldBindJSON(&client); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		updated, err := svc.UpdateClientByEmail(ctx.Request.Context(), client.Email, &client)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update client"})
			return
		}
		ctx.JSON(http.StatusOK, updated)
	}
}

func getClientByEmail(svc service.ClientService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		email := ctx.Param("email")
		client, err := svc.GetClientByEmail(ctx.Request.Context(), email)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "client not found"})
			return
		}
		ctx.JSON(http.StatusOK, client)
	}
}
