package main

import (
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/config"
	"go-data-repository/src/main/go/com/mccusa/datarepository/controller"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

func main() {
	// Initialize logger
	logrus.SetFormatter(&logrus.JSONFormatter{})
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	// Connect to database
	var db = config.Connect(cfg)
	var customDb = config.ConnectCustomDb(cfg)

	// Initialize utilities
	config.NewHTTPClient()
	//scheduler := config.NewTaskScheduler()

	// Initialize repositories
	repository.NewClientRestRepository(db)
	repository.NewCalendlyEventRestRepository(db)
	repository.NewContractTypeRestRepository(db)
	repository.NewJobApplicationRestRepository(db)
	//consultantsRepo := repository.NewConsultantsRestRepository(db)
	//sponsorRepo := repository.NewSponsorRestRepository(db)
	//agencyRepo := repository.NewAgencyRestRepository(db)
	agencyCustomRepository := repository.NewAgencyClientCustomRepository(customDb, util.ClientUtils{})

	// Initialize services
	clientService := service.NewClientService(agencyCustomRepository, repository.NewAgencyClientRepository(db), util.NewClientUtils())
	//activeCampaignService := service.NewActiveCampaignService(clientRepo, calendlyRepo, service.CalendlyService{}, httpClient, contractTypeRepo, jobAppRepo, cfg.BaseURL)
	//getActiveCampaignUsers := service.NewGetActiveCampaignUsers(clientRepo, httpClient, scheduler, cfg.BaseURL, cfg.APIKey)
	//calendlyService := service.NewCalendlyService(calendlyRepo, clientRepo, httpClient, cfg.BaseURL, cfg.APIKey)

	// Initialize Gin router
	router := gin.Default()

	// Register middleware
	router.Use(config.LoggingMiddleware())
	router.Use(config.WebhookUserMiddleware())

	// Register routes
	controller.RegisterClientRoutes(router, clientService)
	//controller.RegisterActiveCampaignWebhookRoutes(api, activeCampaignService, getActiveCampaignUsers)
	//controller.RegisterCalendlyWebhookRoutes(api, calendlyService)

	// Start background scheduler
	//scheduler.StartAsync()

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	logrus.Infof("Server starting on %s", addr)
	router.Run(addr)
}
