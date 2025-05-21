package service

import (
	"context"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
)

const (
	CLIENT_WITH_EMAIL                 = "Client with email "
	DOES_NOT_EXIST_IN_DATA_REPOSITORY = " does not exist in data repository."
)

type ClientService interface {
	CreateClient(ctx context.Context, client *model.AgencyClient) (*model.AgencyClient, error)
	UpdateClientByEmail(ctx context.Context, email string, client *model.AgencyClient) (*model.AgencyClient, error)
	GetClientByEmail(ctx context.Context, email string) (*model.AgencyClient, error)
}

type clientService struct {
	customRepo repository.AgencyClientCustomRepository
	repo       repository.AgencyClientRepository
	utils      util.ClientUtils
}

func NewClientService(
	customRepo repository.AgencyClientCustomRepository,
	repo repository.AgencyClientRepository,
	utils util.ClientUtils,
) ClientService {
	return &clientService{customRepo: customRepo, repo: repo, utils: utils}
}

func (s *clientService) CreateClient(ctx context.Context, client *model.AgencyClient) (*model.AgencyClient, error) {
	return s.customRepo.SaveNewClient(ctx, client)
}

func (s *clientService) UpdateClientByEmail(ctx context.Context, email string, client *model.AgencyClient) (*model.AgencyClient, error) {
	existing, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf("%s%s%w", CLIENT_WITH_EMAIL, email, err)
	}

	// Copy fields if provided
	if client.FirstName != "" {
		existing.FirstName = client.FirstName
	}
	if client.LastName != "" {
		existing.LastName = client.LastName
	}
	// Optional fields
	existing.SecondName = client.SecondName
	existing.SecondLastName = client.SecondLastName

	// Recompute completed name
	existing.CompletedName = s.utils.ConcatenateNameParts(
		existing.FirstName,
		derefString(existing.SecondName),
		existing.LastName,
		derefString(existing.SecondLastName),
	)

	if client.Gender != nil {
		existing.Gender = client.Gender
	}
	if client.Sponsor != nil {
		existing.Sponsor = client.Sponsor
	}
	if client.Vacancy != nil {
		existing.Vacancy = client.Vacancy
	}
	if client.Round != nil {
		existing.Round = client.Round
	}
	if client.ContractType != nil {
		existing.ContractType = client.ContractType
	}

	// Ensure customer code is not null
	if existing.CustomerCode == nil {
		code := ""
		existing.CustomerCode = &code
	}

	return s.repo.Save(ctx, existing)
}

func (s *clientService) GetClientByEmail(ctx context.Context, email string) (*model.AgencyClient, error) {
	client, err := s.repo.FindByEmail(ctx, email)
	if err != nil {
		return nil, fmt.Errorf(CLIENT_WITH_EMAIL + email + DOES_NOT_EXIST_IN_DATA_REPOSITORY)
	}
	return client, nil
}

// Helper to dereference *string or return empty
func derefString(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return ""
}
