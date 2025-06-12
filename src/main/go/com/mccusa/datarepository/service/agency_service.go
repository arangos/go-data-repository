package service

import (
	"context"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/dto"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"golang.org/x/crypto/bcrypt"
)

type AgencyService interface {
	// Add methods that are specific to agency operations
	AuthenticateAgency(ctx context.Context, credentials dto.UserCredentials) (bool, error)
}

type agencyServiceImpl struct {
	agencyRepository repository.AgencyRestRepository
}

func NewAgencyService(agencyRepository repository.AgencyRestRepository) AgencyService {
	return &agencyServiceImpl{agencyRepository: agencyRepository}
}

func (a *agencyServiceImpl) AuthenticateAgency(ctx context.Context, credentials dto.UserCredentials) (bool, error) {
	agency, err := a.agencyRepository.FindByEmail(ctx, credentials.Email)
	if err != nil {
		return false, fmt.Errorf("error finding agency by email %s: %w", credentials.Email, err)
	}
	if agency == nil {
		return false, fmt.Errorf("agency not found by mail %s", credentials.Email)
	}

	err = bcrypt.CompareHashAndPassword([]byte(*agency.Password), []byte(credentials.Password))
	if err != nil {
		return false, fmt.Errorf("invalid password for email %s", credentials.Email)
	}
	return true, nil
}
