package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// CalendlyEventRestRepository provides CRUD and paging operations for CalendlyEvent entities.
type CalendlyEventRestRepository interface {
	// FindAll returns a slice of events and total count for pagination.
	FindAll(ctx context.Context, limit, offset int) ([]model.CalendlyEvent, int64, error)
	// FindByEventUUID finds an event by its UUID.
	FindByEventUUID(ctx context.Context, uuid string) (*model.CalendlyEvent, error)
	// FindByInviteeEmail finds an event by invitee email.
	FindByInviteeEmail(ctx context.Context, email string) (*model.CalendlyEvent, error)
}

type calendlyEventRepo struct {
	db *gorm.DB
}

// NewCalendlyEventRestRepository constructs a new CalendlyEventRestRepository.
func NewCalendlyEventRestRepository(db *gorm.DB) CalendlyEventRestRepository {
	return &calendlyEventRepo{db: db}
}

func (r *calendlyEventRepo) FindAll(ctx context.Context, limit, offset int) ([]model.CalendlyEvent, int64, error) {
	var events []model.CalendlyEvent
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.CalendlyEvent{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.WithContext(ctx).
		Limit(limit).
		Offset(offset).
		Find(&events).Error; err != nil {
		return nil, 0, err
	}
	return events, count, nil
}

func (r *calendlyEventRepo) FindByEventUUID(ctx context.Context, uuid string) (*model.CalendlyEvent, error) {
	var ev model.CalendlyEvent
	err := r.db.WithContext(ctx).
		Where("event_uuid = ?", uuid).
		First(&ev).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ev, err
}

func (r *calendlyEventRepo) FindByInviteeEmail(ctx context.Context, email string) (*model.CalendlyEvent, error) {
	var ev model.CalendlyEvent
	err := r.db.WithContext(ctx).
		Where("invitee_email = ?", email).
		First(&ev).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ev, err
}
