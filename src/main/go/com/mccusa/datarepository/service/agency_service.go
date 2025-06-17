package service

import (
	"context"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/dto"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"golang.org/x/crypto/bcrypt"
)

type AgencyService interface {
	AuthenticateAgency(ctx context.Context, credentials dto.UserCredentials) (bool, error)
	UpdateAgencyPassword(ctx context.Context, credentials dto.UserCredentials) (bool, error)
	GetAgenciesByShortName(ctx context.Context, shortName string) (model.Agency, error)
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

func (a *agencyServiceImpl) UpdateAgencyPassword(ctx context.Context, credentials dto.UserCredentials) (bool, error) {
	agency, err := a.agencyRepository.FindByEmail(ctx, credentials.Email)
	var result = false
	if err != nil {
		return result, fmt.Errorf("Agency with email %s:  does not exist in data repository. %w", credentials.Email, err)
	}
	if agency == nil {
		return result, fmt.Errorf("Agency with email %s:  does not exist in data repository.", credentials.Email)
	}

	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
	if err != nil {
		return result, fmt.Errorf("error hashing password for email %s: %w", credentials.Email, err)
	}
	hashedPassword := string(hashedBytes)
	agency.Password = &hashedPassword
	result, err = a.agencyRepository.UpdateAgency(ctx, *agency)
	if err != nil {
		return false, fmt.Errorf("error updating agency password for email %s: %w", credentials.Email, err)
	}
	return result, nil
}

func (a *agencyServiceImpl) GetAgenciesByShortName(ctx context.Context, shortName string) (model.Agency, error) {
	agency, err := a.agencyRepository.FindByShortName(ctx, shortName)
	if err != nil {
		return agency, fmt.Errorf("error finding agency by short name %s: %w", shortName, err)
	}
	return agency, nil
}
