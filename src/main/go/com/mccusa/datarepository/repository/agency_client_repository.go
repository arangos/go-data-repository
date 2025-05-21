package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

type AgencyClientRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]model.AgencyClient, int64, error)
	FindTopByAgencyOrderByCustomerCodeDesc(ctx context.Context, agencyCode string) (*model.AgencyClient, error)
	FindByEmail(ctx context.Context, email string) (*model.AgencyClient, error)
	Save(ctx context.Context, client *model.AgencyClient) (*model.AgencyClient, error)
}

type agencyClientRepository struct {
	db *gorm.DB
}

// NewAgencyClientRepository constructs a new repository
func NewAgencyClientRepository(db *gorm.DB) AgencyClientRepository {
	return &agencyClientRepository{db: db}
}

// FindAll retrieves a paginated list of clients and total count
func (r *agencyClientRepository) FindAll(ctx context.Context, limit, offset int) ([]model.AgencyClient, int64, error) {
	var clients []model.AgencyClient
	var count int64

	// Count total
	if err := r.db.WithContext(ctx).Model(&model.AgencyClient{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// Fetch paginated data
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&clients).
		Error; err != nil {
		return nil, 0, err
	}
	return clients, count, nil
}

// FindTopByAgencyOrderByCustomerCodeDesc returns the latest client by customer code
func (r *agencyClientRepository) FindTopByAgencyOrderByCustomerCodeDesc(ctx context.Context, agencyCode string) (*model.AgencyClient, error) {
	var client model.AgencyClient
	err := r.db.WithContext(ctx).
		Where("agency = ? AND customer_code IS NOT NULL", agencyCode).
		Order("customer_code DESC").
		Limit(1).
		First(&client).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &client, nil
}

// FindByEmail looks up a client by email (case-insensitive)
func (r *agencyClientRepository) FindByEmail(ctx context.Context, email string) (*model.AgencyClient, error) {
	var client model.AgencyClient
	err := r.db.WithContext(ctx).
		Where("LOWER(email) = LOWER(?)", email).
		First(&client).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, gorm.ErrRecordNotFound
	}
	return &client, err
}

// Save creates or updates a client record
func (r *agencyClientRepository) Save(ctx context.Context, client *model.AgencyClient) (*model.AgencyClient, error) {
	if err := r.db.WithContext(ctx).Save(client).Error; err != nil {
		return nil, err
	}
	return client, nil
}
