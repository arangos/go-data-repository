package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// ContractTypeRestRepository provides CRUD and paging operations for ContractType entities.
type ContractTypeRestRepository interface {
	// FindAll returns a slice of contract types and total count for pagination.
	FindAll(ctx context.Context, limit, offset int) ([]model.ContractType, int64, error)
	// FindByType retrieves a ContractType by its name.
	FindByType(ctx context.Context, name string) (*model.ContractType, error)
	// FindByTemplateID retrieves a ContractType by its DocusignTemplateId.
	FindByTemplateID(ctx context.Context, templateID string) (*model.ContractType, error)
}

type contractTypeRepo struct {
	db *gorm.DB
}

// NewContractTypeRestRepository constructs a new ContractTypeRestRepository.
func NewContractTypeRestRepository(db *gorm.DB) ContractTypeRestRepository {
	return &contractTypeRepo{db: db}
}

func (r *contractTypeRepo) FindAll(ctx context.Context, limit, offset int) ([]model.ContractType, int64, error) {
	var items []model.ContractType
	var count int64

	// Count total records
	if err := r.db.WithContext(ctx).
		Model(&model.ContractType{}).
		Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// Fetch paginated slice
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (r *contractTypeRepo) FindByType(ctx context.Context, name string) (*model.ContractType, error) {
	var ct model.ContractType
	err := r.db.WithContext(ctx).
		Where("contract_type = ?", name).
		First(&ct).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ct, err
}

func (r *contractTypeRepo) FindByTemplateID(ctx context.Context, templateID string) (*model.ContractType, error) {
	var ct model.ContractType
	err := r.db.WithContext(ctx).
		Where("docusign_template_id = ?", templateID).
		First(&ct).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ct, err
}
