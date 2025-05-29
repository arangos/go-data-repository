package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-co-op/gocron"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"io"
	"net/http"
)

// GetActiveCampaignUsers periodically syncs contacts from ActiveCampaign.
type GetActiveCampaignUsers struct {
	ClientRepo repository.ClientRestRepository
	HTTPClient *http.Client
	Scheduler  *gocron.Scheduler
	APIKey     string
	Offset     int
	BaseURL    string
}

// NewGetActiveCampaignUsers initializes the service and schedules daily sync.
func NewGetActiveCampaignUsers(
	clientRepo repository.ClientRestRepository,
	httpClient *http.Client,
	scheduler *gocron.Scheduler,
	baseURL, apiKey string,
) *GetActiveCampaignUsers {
	svc := &GetActiveCampaignUsers{
		ClientRepo: clientRepo,
		HTTPClient: httpClient,
		Scheduler:  scheduler,
		APIKey:     apiKey,
		BaseURL:    baseURL,
	}
	// Schedule daily at midnight
	scheduler.Every(24).Hours().At("00:00").Do(func() {
		svc.GetContactsFromAPI(context.Background())
	})
	return svc
}

// GetContactsFromAPI fetches contacts, processes them, and reschedules if needed.
func (svc *GetActiveCampaignUsers) GetContactsFromAPI(ctx context.Context) {
	url := fmt.Sprintf("%s/api/3/contacts?offset=%d&limit=100", svc.BaseURL, svc.Offset)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return
	}
	req.Header.Set("Api-Token", svc.APIKey)

	resp, err := svc.HTTPClient.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var result struct {
		Contacts []model.Client `json:"contacts"`
		Meta     struct {
			Total int `json:"total"`
		} `json:"meta"`
	}
	if err := json.Unmarshal(body, &result); err != nil {
		return
	}

	for _, c := range result.Contacts {
		existing, _ := svc.ClientRepo.FindByEmail(ctx, c.Email)
		if existing == nil {
			svc.ClientRepo.Save(ctx, &c)
		}
	}

	if svc.Offset+100 < result.Meta.Total {
		svc.Offset += 100
		// reschedule next fetch in 5 minutes
		svc.Scheduler.Every(5).Minutes().Do(func() {
			svc.GetContactsFromAPI(ctx)
		})
	}
}
