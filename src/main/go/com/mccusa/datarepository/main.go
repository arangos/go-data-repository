package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"go-data-repository/src/main/go/com/mccusa/datarepository/config"
	"go-data-repository/src/main/go/com/mccusa/datarepository/controller"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"
	"log"
	"os"
	"os/signal"
	"syscall"
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
	var db = config.ConnectGORM(cfg)
	var customDb = config.ConnectNativeDbQueries(cfg)

	// Initialize utilities
	config.NewHTTPClient()
	//scheduler := config.NewTaskScheduler()

	// Initialize repositories
	repository.NewClientRestRepository(db)
	repository.NewCalendlyEventRestRepository(db)
	repository.NewContractTypeRestRepository(db)
	repository.NewJobApplicationRestRepository(db)
	docusignRepository := repository.NewDocusignEnvelopesRepository(db)
	//consultantsRepo := repository.NewConsultantsRestRepository(db)
	//sponsorRepo := repository.NewSponsorRestRepository(db)
	agencyRepository := repository.NewAgencyRestRepository(db)
	agencyCustomRepository := repository.NewAgencyClientCustomRepository(customDb, util.ClientUtils{})

	// Initialize services
	clientService := service.NewClientService(agencyCustomRepository, repository.NewAgencyClientRepository(db), util.NewClientUtils())
	agencyService := service.NewAgencyService(agencyRepository)
	docusignService := service.NewDocusignService(docusignRepository)
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
	controller.RegisterAgencyRoutes(router, agencyService)
	//controller.RegisterActiveCampaignWebhookRoutes(api, activeCampaignService, getActiveCampaignUsers)
	//controller.RegisterCalendlyWebhookRoutes(api, calendlyService)

	// Start background scheduler
	setUpScheduler(docusignService, err)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	logrus.Infof("Server starting on %s", addr)
	err = router.Run(addr)
	if err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}

func setUpScheduler(docusignService service.DocusignService, err error) {
	job := func(ctx context.Context) {
		if err := docusignService.ProcessDocusignEnvelopes(); err != nil {
			logrus.Errorf("Docusign job failed: %v", err)
		}
	}
	scheduler, err := config.DocusignScheduleJob(job)
	if err != nil {
		log.Fatalf("failed to start scheduler: %v", err)
	}
	log.Println("scheduler started; running job every 30m")

	// 3) Wait for SIGINT/SIGTERM to gracefully shut down.
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig

	// 4) Stop the scheduler.
	ctx := scheduler.Stop()
	<-ctx.Done() // wait for running jobs
	log.Println("scheduler stopped")
}
