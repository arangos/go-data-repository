package main

import (
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/config"
	"go-data-repository/src/main/go/com/mccusa/datarepository/controller"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
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

	// Connect to the database with sqlx
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=require",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)
	customDb, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logrus.Fatalf("Failed to connect to database with sqlx: %v", err)
	}
	logrus.Infof("Connected to database at %s", customDb.DriverName())
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
