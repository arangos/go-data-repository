package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// ClientRestRepository provides CRUD and paging operations for Client entities.
type ClientRestRepository interface {
	// FindAll returns a slice of clients and total count for pagination.
	FindAll(ctx context.Context, limit, offset int) ([]model.Client, int64, error)
	// FindByEmail returns a Client matching the given email (case-insensitive).
	FindByEmail(ctx context.Context, email string) (*model.Client, error)
	// Save Saves a new client to the database.
	Save(ctx context.Context, client *model.Client) error
}

type clientRestRepo struct {
	db *gorm.DB
}

func (r *clientRestRepo) Save(ctx context.Context, client *model.Client) error {
	if err := r.db.WithContext(ctx).Create(client).Error; err != nil {
		return err
	}
	return nil
}

// NewClientRestRepository constructs a new ClientRestRepository.
func NewClientRestRepository(db *gorm.DB) ClientRestRepository {
	return &clientRestRepo{db: db}
}

func (r *clientRestRepo) FindAll(ctx context.Context, limit, offset int) ([]model.Client, int64, error) {
	var clients []model.Client
	var count int64
	// Count total
	if err := r.db.WithContext(ctx).Model(&model.Client{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	// Fetch paginated
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&clients).Error; err != nil {
		return nil, 0, err
	}
	return clients, count, nil
}

func (r *clientRestRepo) FindByEmail(ctx context.Context, email string) (*model.Client, error) {
	var client model.Client
	err := r.db.WithContext(ctx).
		Where("LOWER(email) = LOWER(?)", email).
		First(&client).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &client, err
}
