package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

type ConsultantsRestRepository interface {
	// FindAll returns a page of consultants and the total count
	FindAll(ctx context.Context, limit, offset int) ([]model.Consultant, int64, error)
	// FindByMail looks up a single consultant by email
	FindByMail(ctx context.Context, mail string) (*model.Consultant, error)
}

type consultantsRestRepo struct {
	db *gorm.DB
}

// NewConsultantsRestRepository constructs the repository
func NewConsultantsRestRepository(db *gorm.DB) ConsultantsRestRepository {
	return &consultantsRestRepo{db: db}
}

func (r *consultantsRestRepo) FindAll(ctx context.Context, limit, offset int) ([]model.Consultant, int64, error) {
	var consultants []model.Consultant
	var count int64

	// total count
	if err := r.db.WithContext(ctx).Model(&model.Consultant{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// paginated fetch
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&consultants).Error; err != nil {
		return nil, 0, err
	}
	return consultants, count, nil
}

func (r *consultantsRestRepo) FindByMail(ctx context.Context, mail string) (*model.Consultant, error) {
	var c model.Consultant
	err := r.db.WithContext(ctx).
		Where("mail = ?", mail).
		First(&c).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &c, err
}
