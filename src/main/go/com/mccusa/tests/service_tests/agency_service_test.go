package service_tests

import (
	"context"
	//"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/dto"
	"go-data-repository/src/main/go/com/mccusa/datarepository/service"
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock repository
type MockAgencyRestRepository struct {
	mock.Mock
}

func (m *MockAgencyRestRepository) FindAll(ctx context.Context, limit int, offset int) ([]model.Agency, int64, error) {
	args := m.Called(ctx, limit, offset)
	return args.Get(0).([]model.Agency), args.Get(1).(int64), args.Error(2)
}

func (m *MockAgencyRestRepository) FindByEmail(ctx context.Context, email string) (*model.Agency, error) {
	args := m.Called(ctx, email)
	return args.Get(0).(*model.Agency), args.Error(1)
}

func (m *MockAgencyRestRepository) UpdateAgency(ctx context.Context, agency model.Agency) (bool, error) {
	args := m.Called(ctx, agency)
	return args.Bool(0), args.Error(1)
}

func (m *MockAgencyRestRepository) FindByShortName(ctx context.Context, shortName string) (model.Agency, error) {
	args := m.Called(ctx, shortName)
	return args.Get(0).(model.Agency), args.Error(1)
}

func TestAuthenticateAgencySuccess(t *testing.T) {
	mockRepo := new(MockAgencyRestRepository)
	service := service.NewAgencyService(mockRepo)

	ctx := context.Background()
	credentials := dto.UserCredentials{Email: "test@example.com", Password: "password123"}

	// Mock repository behavior
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(&model.Agency{Password: func() *string {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(credentials.Password), bcrypt.DefaultCost)
		str := string(hashed)
		return &str
	}()}, nil)

	// Test valid credentials
	isAuthenticated, err := service.AuthenticateAgency(ctx, credentials)
	assert.NoError(t, err)
	assert.True(t, isAuthenticated)
}

func TestAuthenticateAgencyFail(t *testing.T) {
	mockRepo := new(MockAgencyRestRepository)
	service := service.NewAgencyService(mockRepo)

	ctx := context.Background()
	credentials := dto.UserCredentials{Email: "test@example.com", Password: "password123"}
	// Test invalid password
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(&model.Agency{Password: func() *string {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("wrongpassword"), bcrypt.DefaultCost)
		str := string(hashed)
		return &str
	}()}, nil)

	isAuthenticated, err := service.AuthenticateAgency(ctx, credentials)
	assert.Error(t, err)
	assert.False(t, isAuthenticated)
}

func TestAuthenticateAgencyNotFound(t *testing.T) {
	mockRepo := new(MockAgencyRestRepository)
	service := service.NewAgencyService(mockRepo)

	ctx := context.Background()
	credentials := dto.UserCredentials{Email: "test@example.com", Password: "password123"}
	// Test invalid password
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(&model.Agency{Password: func() *string {
		hashed, _ := bcrypt.GenerateFromPassword([]byte("wrongpassword"), bcrypt.DefaultCost)
		str := string(hashed)
		return &str
	}()}, nil)

	// Test agency not found
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(nil, nil)

	isAuthenticated, err := service.AuthenticateAgency(ctx, credentials)
	assert.Error(t, err)
	assert.False(t, isAuthenticated)
}

/*
func TestUpdateAgencyPassword(t *testing.T) {
	mockRepo := new(MockAgencyRestRepository)
	service := service.NewAgencyService(mockRepo)

	ctx := context.Background()
	credentials := dto.UserCredentials{Email: "test@example.com", Password: "newpassword123"}

	// Mock repository behavior
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(&model.Agency{Email: credentials.Email}, nil)
	mockRepo.On("UpdateAgency", ctx, mock.Anything).Return(true, nil)

	// Test successful password update
	isUpdated, err := service.UpdateAgencyPassword(ctx, credentials)
	assert.NoError(t, err)
	assert.True(t, isUpdated)

	// Test agency not found
	mockRepo.On("FindByEmail", ctx, credentials.Email).Return(nil, nil)

	isUpdated, err = service.UpdateAgencyPassword(ctx, credentials)
	assert.Error(t, err)
	assert.False(t, isUpdated)
}
*/

/*
func TestGetAgenciesByShortName(t *testing.T) {
	mockRepo := new(MockAgencyRestRepository)
	service := service.NewAgencyService(mockRepo)

	ctx := context.Background()
	shortName := "shortName"

	// Mock repository behavior
	mockRepo.On("FindByShortName", ctx, shortName).Return(&model.Agency{AgencyShortName: &shortName}, nil)

	// Test valid short name
	agency, err := service.GetAgenciesByShortName(ctx, shortName)
	assert.NoError(t, err)
	assert.Equal(t, shortName, agency.AgencyShortName)

	// Test agency not found
	mockRepo.On("FindByShortName", ctx, shortName).Return(nil, errors.New("not found"))

	agency, err = service.GetAgenciesByShortName(ctx, shortName)
	assert.Error(t, err)
	assert.Empty(t, agency.AgencyShortName)
}
*/
