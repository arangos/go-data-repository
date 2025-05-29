package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// SponsorRestRepository defines methods to manage Sponsor entities.
type SponsorRestRepository interface {
	// FindAll returns a slice of sponsors and the total count for pagination.
	FindAll(ctx context.Context, limit, offset int) ([]model.Sponsor, int64, error)
	// FindByEmail looks up a single Sponsor by sponsor_email.
	FindByEmail(ctx context.Context, email string) (*model.Sponsor, error)
	// FindByName looks up a single Sponsor by sponsor_name.
	FindByName(ctx context.Context, name string) (*model.Sponsor, error)
}

type sponsorRestRepo struct {
	db *gorm.DB
}

// NewSponsorRestRepository constructs a new SponsorRestRepository.
func NewSponsorRestRepository(db *gorm.DB) SponsorRestRepository {
	return &sponsorRestRepo{db: db}
}

func (r *sponsorRestRepo) FindAll(ctx context.Context, limit, offset int) ([]model.Sponsor, int64, error) {
	var items []model.Sponsor
	var count int64

	// Total count
	if err := r.db.WithContext(ctx).Model(&model.Sponsor{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// Paginated fetch
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (r *sponsorRestRepo) FindByEmail(ctx context.Context, email string) (*model.Sponsor, error) {
	var s model.Sponsor
	err := r.db.WithContext(ctx).
		Where("sponsor_email = ?", email).
		First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &s, err
}

func (r *sponsorRestRepo) FindByName(ctx context.Context, name string) (*model.Sponsor, error) {
	var s model.Sponsor
	err := r.db.WithContext(ctx).
		Where("sponsor_name = ?", name).
		First(&s).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &s, err
}
