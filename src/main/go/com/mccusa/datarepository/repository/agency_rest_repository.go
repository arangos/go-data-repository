package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// AgencyRestRepository provides CRUD and paging operations for Agency entities.
type AgencyRestRepository interface {
	// FindAll returns a slice of agencies and total count for pagination.
	FindAll(ctx context.Context, limit int, offset int) ([]model.Agency, int64, error)
	// FindByShortName returns agencies matching the given short name.
	FindByShortName(ctx context.Context, shortName string) ([]model.Agency, error)
	// FindByEmail returns the Agency with the given email.
	FindByEmail(ctx context.Context, email string) (*model.Agency, error)
}

type agencyRestRepo struct {
	db *gorm.DB
}

// NewAgencyRestRepository constructs a new AgencyRestRepository.
func NewAgencyRestRepository(db *gorm.DB) AgencyRestRepository {
	return &agencyRestRepo{db: db}
}

func (r *agencyRestRepo) FindAll(ctx context.Context, limit, offset int) ([]model.Agency, int64, error) {
	var agencies []model.Agency
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.Agency{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&agencies).Error; err != nil {
		return nil, 0, err
	}
	return agencies, count, nil
}

func (r *agencyRestRepo) FindByShortName(ctx context.Context, shortName string) ([]model.Agency, error) {
	var agencies []model.Agency
	err := r.db.WithContext(ctx).
		Where("agency_short_name = ?", shortName).
		Find(&agencies).Error
	return agencies, err
}

func (r *agencyRestRepo) FindByEmail(ctx context.Context, email string) (*model.Agency, error) {
	var agency model.Agency
	err := r.db.WithContext(ctx).
		Where("email = ?", email).
		First(&agency).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &agency, err
}
