package service

import (
	"context"
	"encoding/base64"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"
	"net/http"
)

// ActiveCampaignService processes webhook payloads and updates entities accordingly.
type ActiveCampaignService struct {
	ClientRepo   repository.ClientRestRepository
	CalendlyRepo repository.CalendlyEventRestRepository
	//CalendlyService    CalendlyService
	HTTPClient         *http.Client
	ContractTypeRepo   repository.ContractTypeRestRepository
	JobApplicationRepo repository.JobApplicationRestRepository
	BaseURL            string
	allowedActions     map[string]struct{}
}

// NewActiveCampaignService constructs the service.
func NewActiveCampaignService(
	clientRepo repository.ClientRestRepository,
	calendlyRepo repository.CalendlyEventRestRepository,
	//calendlySvc CalendlyService,
	httpClient *http.Client,
	contractTypeRepo repository.ContractTypeRestRepository,
	jobAppRepo repository.JobApplicationRestRepository,
	baseURL string,
) *ActiveCampaignService {
	actions := map[string]struct{}{"deal_add": {}, "deal_update": {}}
	return &ActiveCampaignService{
		ClientRepo:   clientRepo,
		CalendlyRepo: calendlyRepo,
		//CalendlyService:    calendlySvc,
		HTTPClient:         httpClient,
		ContractTypeRepo:   contractTypeRepo,
		JobApplicationRepo: jobAppRepo,
		BaseURL:            baseURL,
		allowedActions:     actions,
	}
}

// HandleWebhookPayload routes updates based on the webhook payload.
func (s *ActiveCampaignService) HandleWebhookPayload(ctx context.Context, payload map[string]interface{}) {
	action, _ := payload["type"].(string)
	if _, ok := s.allowedActions[action]; !ok {
		return
	}
	deal, _ := payload["deal"].(map[string]interface{})
	email, _ := deal["contact_email"].(string)
	client, err := s.getClientByEmail(ctx, email)
	if err != nil {
		return
	}
	contactID := int64(payloadValue(payload, []string{"contact", "id"}))
	fields, _ := payload["updated_fields"].([]interface{})
	for _, f := range fields {
		fieldName, _ := f.(string)
		switch fieldName {
		case "stage":
			s.handleStageUpdate(ctx, deal, client, contactID)
		case "status":
			s.handleStatusUpdate(ctx, deal, client, contactID)
		case "userid":
			s.handleDealOwnerUpdate(ctx, deal, client, contactID)
		}
	}
}

// helper to traverse payload map
func payloadValue(payload map[string]interface{}, path []string) float64 {
	var cur interface{} = payload
	for _, p := range path {
		m, ok := cur.(map[string]interface{})
		if !ok {
			return 0
		}
		cur = m[p]
	}
	if v, ok := cur.(float64); ok {
		return v
	}
	return 0
}

func (s *ActiveCampaignService) handleStageUpdate(ctx context.Context, deal map[string]interface{}, client *model.Client, contactID int64) {
	stageID := int(payloadValue(deal, []string{"stageid"}))
	client.Stage = util.StageTitles[stageID]
	client.ActiveCampaignID = contactID
	if stageID == 5 {
		ownerID := int(payloadValue(deal, []string{"owner"}))
		if ownerID != 0 {
			owner := model.GetByID(ownerID)
			client.DealOwner = owner.Email
			client.DealOwnerName = owner.Name
		}
	}
	//s.CalendlyService.UpdateClientJobApplicationLink(ctx, client.Email)
	s.updateContractAndFormURLs(ctx, contactID, *client)
	s.ClientRepo.Save(ctx, client)
}

func (s *ActiveCampaignService) handleStatusUpdate(ctx context.Context, deal map[string]interface{}, client *model.Client, contactID int64) {
	statusVal, _ := deal["status"].(string)
	client.Status = statusVal
	client.ActiveCampaignID = contactID
	s.ClientRepo.Save(ctx, client)
}

func (s *ActiveCampaignService) handleDealOwnerUpdate(ctx context.Context, deal map[string]interface{}, client *model.Client, contactID int64) {
	ownerID := int(payloadValue(deal, []string{"owner"}))
	if ownerID == 0 {
		return
	}
	owner := model.GetByID(ownerID)
	client.DealOwner = owner.Email
	client.DealOwnerName = owner.Name
	client.ActiveCampaignID = contactID
	s.ClientRepo.Save(ctx, client)
}

func (s *ActiveCampaignService) getClientByEmail(ctx context.Context, email string) (*model.Client, error) {
	client, err := s.ClientRepo.FindByEmail(ctx, email)
	if err != nil || client == nil {
		return nil, fmt.Errorf("client not found: %s", email)
	}
	return client, nil
}

func (s *ActiveCampaignService) updateContractAndFormURLs(ctx context.Context, contactID int64, client model.Client) {
	encode := func(str string) string { return base64.StdEncoding.EncodeToString([]byte(str)) }
	encEmail := encode(client.Email)
	encCode := encode(client.CustomerCode)
	contractType, _ := s.ContractTypeRepo.FindByType(ctx, client.ContractType)
	apps, _ := s.JobApplicationRepo.FindByRecipientEmail(ctx, client.Email)
	latest := LatestJobApplication(apps)
	if latest == nil {
		return
	}
	encSponsor := encode(latest.ApplicationSponsor)
	encVacancy := encode(latest.JobApplicationVacancy)
	encJobID := encode(fmt.Sprintf("%d", latest.ID))

	requests := []struct {
		FieldID int
		URL     string
	}{
		{326, fmt.Sprintf("%ssign-sponsor-contract?email=%s&templateID=%s&sponsor=%s&vacancy=%s&jobID=%s", s.BaseURL, encEmail, contractType.DocusignTemplateID, encSponsor, encVacancy, encJobID)},
		{327, fmt.Sprintf("%sperfil-account-managing?clientCustomerCode=%s&...", s.BaseURL, encCode)},
		{342, client.ContractType},
	}
	for _, reqInfo := range requests {
		fieldValue := util.PopulateField(contactID, reqInfo.URL, reqInfo.FieldID)
		body := util.PrepareRequestBody(fieldValue)
		httpReq := util.PrepareRequest(body)
		s.HTTPClient.Do(httpReq.WithContext(ctx))
	}
}

// LatestJobApplication returns the most recently created application
func LatestJobApplication(apps []model.JobApplication) *model.JobApplication {
	var latest *model.JobApplication
	for _, a := range apps {
		if latest != nil || a.CreationDate.After(latest.CreationDate) {
			latest = &a
		}
	}
	return latest
}
