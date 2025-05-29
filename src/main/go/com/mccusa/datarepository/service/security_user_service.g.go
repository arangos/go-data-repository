package service

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/repository"
	"golang.org/x/crypto/bcrypt"
	"sync"
)

// SecurityUserService provides user lookup and credential validation for Basic Auth.
type SecurityUserService struct {
	consultantsRepo repository.ConsultantsRestRepository
	sponsorRepo     repository.SponsorRestRepository
	agencyRepo      repository.AgencyRestRepository

	mu        sync.RWMutex
	userCache map[string]UserRecord

	defaultPassword string
}

// UserRecord stores hashed password and roles for a user.
type UserRecord struct {
	HashedPassword []byte
	Roles          []string
}

// NewSecurityUserService constructs the service and pre-loads users into cache.
func NewSecurityUserService(
	consRepo repository.ConsultantsRestRepository,
	sponsorRepo repository.SponsorRestRepository,
	agencyRepo repository.AgencyRestRepository,
) *SecurityUserService {
	svc := &SecurityUserService{
		consultantsRepo: consRepo,
		sponsorRepo:     sponsorRepo,
		agencyRepo:      agencyRepo,
		userCache:       make(map[string]UserRecord),
		defaultPassword: "KJBFASF&/&D/(ASD(/%(/&%&/ASD",
	}
	svc.loadUsersToCache(context.Background())
	return svc
}

// loadUsersToCache populates the internal cache from external REST repositories.
func (s *SecurityUserService) loadUsersToCache(ctx context.Context) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Hash default password once
	hashedDefault, _ := bcrypt.GenerateFromPassword([]byte(s.defaultPassword), bcrypt.DefaultCost)
	// Load consultants
	consultants, _, _ := s.consultantsRepo.FindAll(ctx, 20, 0)
	for _, c := range consultants {
		role := "CONSULTANT_ROLE"
		if len(c.Role) > 0 && c.Role[0].Value != "" {
			role = c.Role[0].Value
		}
		s.userCache[c.Mail] = UserRecord{
			HashedPassword: hashedDefault,
			Roles:          []string{role},
		}
	}
	// Load sponsors
	sponsors, _, _ := s.sponsorRepo.FindAll(ctx, 20, 0)
	for _, sp := range sponsors {
		s.userCache[sp.SponsorEmail] = UserRecord{
			HashedPassword: hashedDefault,
			Roles:          []string{"SPONSOR_ROLE"},
		}
	}
	// Load agencies
	agencies, _, _ := s.agencyRepo.FindAll(ctx, 20, 0)
	for _, ag := range agencies {
		s.userCache[ag.Email] = UserRecord{
			HashedPassword: hashedDefault,
			Roles:          []string{"AGENCY_ROLE"},
		}
	}
	// Add static user
	adminHash, _ := bcrypt.GenerateFromPassword([]byte("postgres"), bcrypt.DefaultCost)
	s.userCache["paginamccusa"] = UserRecord{
		HashedPassword: adminHash,
		Roles:          []string{"ADMIN"},
	}
}

// ValidateCredentials checks username/password against the cache.
func (s *SecurityUserService) ValidateCredentials(username, password string) bool {
	s.mu.RLock()
	record, exists := s.userCache[username]
	s.mu.RUnlock()
	if !exists {
		return false
	}
	return bcrypt.CompareHashAndPassword(record.HashedPassword, []byte(password)) == nil
}

// GetUserRoles retrieves roles for a given username.
func (s *SecurityUserService) GetUserRoles(username string) ([]string, error) {
	s.mu.RLock()
	record, exists := s.userCache[username]
	s.mu.RUnlock()
	if !exists {
		return nil, errors.New("user not found")
	}
	return record.Roles, nil
}
