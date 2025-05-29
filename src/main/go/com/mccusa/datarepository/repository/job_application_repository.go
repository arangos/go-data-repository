package repository

import (
	"context"
	"errors"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"gorm.io/gorm"
)

// JobApplicationRestRepository defines methods for JobApplication entity operations.
type JobApplicationRestRepository interface {
	FindAll(ctx context.Context, limit, offset int) ([]model.JobApplication, int64, error)
	FindBySenderEmail(ctx context.Context, senderEmail string) (*model.JobApplication, error)
	FindByRecipientEmail(ctx context.Context, recipientEmail string) ([]model.JobApplication, error)
	FindWithClientInfo(ctx context.Context, sponsorName string, limit, offset int) ([]struct {
		Job    model.JobApplication
		Client model.Client
	}, int64, error)
	FindHistoricalWithClientInfo(ctx context.Context, sponsorName string, limit, offset int) ([]struct {
		Job    model.JobApplication
		Client model.Client
	}, int64, error)
	DeleteDuplicate(ctx context.Context, customerCode, jobApplicationVacancy string, jobApplicationId uint) error
}

type jobAppRepo struct {
	db *gorm.DB
}

// NewJobApplicationRestRepository creates a new JobApplicationRestRepository.
func NewJobApplicationRestRepository(db *gorm.DB) JobApplicationRestRepository {
	return &jobAppRepo{db: db}
}

func (r *jobAppRepo) FindAll(ctx context.Context, limit, offset int) ([]model.JobApplication, int64, error) {
	var items []model.JobApplication
	var count int64
	if err := r.db.WithContext(ctx).Model(&model.JobApplication{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	if err := r.db.WithContext(ctx).Limit(limit).Offset(offset).Find(&items).Error; err != nil {
		return nil, 0, err
	}
	return items, count, nil
}

func (r *jobAppRepo) FindBySenderEmail(ctx context.Context, senderEmail string) (*model.JobApplication, error) {
	var ja model.JobApplication
	err := r.db.WithContext(ctx).Where("sender_email = ?", senderEmail).First(&ja).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &ja, err
}

func (r *jobAppRepo) FindByRecipientEmail(ctx context.Context, recipientEmail string) ([]model.JobApplication, error) {
	var list []model.JobApplication
	if err := r.db.WithContext(ctx).Where("recipient_email = ?", recipientEmail).Find(&list).Error; err != nil {
		return nil, err
	}
	return list, nil
}

func (r *jobAppRepo) FindWithClientInfo(ctx context.Context, sponsorName string, limit, offset int) ([]struct {
	Job    model.JobApplication
	Client model.Client
}, int64, error) {
	var rows []struct {
		Job    model.JobApplication
		Client model.Client
	}
	var count int64
	r.db.WithContext(ctx).Model(&model.JobApplication{}).
		Joins("inner join client c on lower(job_application.recipient_email)=lower(c.email)").
		Where("job_application.application_sponsor = ? AND job_application.approved = ?", sponsorName, "PENDING").
		Count(&count)
	if err := r.db.WithContext(ctx).
		Limit(limit).Offset(offset).
		Joins("inner join client c on lower(job_application.recipient_email)=lower(c.email)").
		Where("job_application.application_sponsor = ? AND job_application.approved = ?", sponsorName, "PENDING").
		Select("job_application.*, c.*").
		Scan(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func (r *jobAppRepo) FindHistoricalWithClientInfo(ctx context.Context, sponsorName string, limit, offset int) ([]struct {
	Job    model.JobApplication
	Client model.Client
}, int64, error) {
	var rows []struct {
		Job    model.JobApplication
		Client model.Client
	}
	var count int64
	r.db.WithContext(ctx).Model(&model.JobApplication{}).
		Joins("inner join client c on lower(job_application.recipient_email)=lower(c.email)").
		Where("job_application.application_sponsor = ? AND (job_application.approved = ? OR job_application.approved = ?)", sponsorName, "ACCEPT", "REJECT").
		Count(&count)
	if err := r.db.WithContext(ctx).
		Limit(limit).Offset(offset).
		Joins("inner join client c on lower(job_application.recipient_email)=lower(c.email)").
		Where("job_application.application_sponsor = ? AND (job_application.approved = ? OR job_application.approved = ?)", sponsorName, "ACCEPT", "REJECT").
		Select("job_application.*, c.*").
		Scan(&rows).Error; err != nil {
		return nil, 0, err
	}
	return rows, count, nil
}

func (r *jobAppRepo) DeleteDuplicate(ctx context.Context, customerCode, jobApplicationVacancy string, jobApplicationId uint) error {
	return r.db.WithContext(ctx).
		Where("customer_code = ? AND job_application_vacancy = ? AND approved = ? AND id <> ?", customerCode, jobApplicationVacancy, "PENDING", jobApplicationId).
		Delete(&model.JobApplication{}).Error
}
