package model

import (
	"gorm.io/gorm"
	"time"
)

// JobApplication maps to the 'job_application' table
type JobApplication struct {
	ID                    uint           `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	SenderEmail           *string        `gorm:"column:SENDER_EMAIL;size:255" json:"senderEmail,omitempty"`
	RecipientEmail        *string        `gorm:"column:RECIPIENT_EMAIL;size:255" json:"recipientEmail,omitempty"`
	JobApplicationName    *string        `gorm:"column:JOB_APPLICATION_NAME;size:255" json:"jobApplicationName,omitempty"`
	JobApplicationID      *string        `gorm:"column:JOB_APPLICATION_ID;size:255" json:"jobApplicationId,omitempty"`
	CreationDate          time.Time      `gorm:"column:CREATION_DATE" json:"creationDate"`
	CreatedBy             *string        `gorm:"column:CREATED_BY;size:255" json:"createdBy,omitempty"`
	ApplicationURL        *string        `gorm:"column:APPLICATION_URL" json:"applicationUrl,omitempty"`
	ApplicationSponsor    *string        `gorm:"column:APPLICATION_SPONSOR" json:"applicationSponsor,omitempty"`
	ApplicationData       []JsonMetadata `gorm:"column:APPLICATION_DATA;type:json" json:"applicationData,omitempty"`
	Approved              *string        `gorm:"column:APPROVED" json:"approved,omitempty"`
	LastModifiedBy        *string        `gorm:"column:LAST_MODIFIED_BY;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate      *time.Time     `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate,omitempty"`
	AWSDocuments          []JsonMetadata `gorm:"column:AWS_DOCUMENTS;type:json" json:"awsDocuments,omitempty"`
	MccusaApprovedCheck   *bool          `gorm:"column:MCCUSA_APPROVED_CHECK" json:"mccusaApprovedCheck,omitempty"`
	JobApplicationAgency  *string        `gorm:"column:JOB_APPLICATION_AGENCY" json:"jobApplicationAgency,omitempty"`
	AgencyApprovedCheck   *bool          `gorm:"column:AGENCY_APPROVED_CHECK" json:"agencyApprovedCheck,omitempty"`
	JobApplicationVacancy *string        `gorm:"column:JOB_APPLICATION_VACANCY" json:"jobApplicationVacancy,omitempty"`
	CustomerCode          *string        `gorm:"column:CUSTOMER_CODE" json:"customerCode,omitempty"`
	ApprovalDate          *time.Time     `gorm:"column:APPROVAL_DATE" json:"approvalDate,omitempty"`
}

// BeforeCreate sets defaults before inserting a new record
func (j *JobApplication) BeforeCreate(tx *gorm.DB) (err error) {
	j.CreationDate = time.Now()
	// Defaults for approval checks
	approved := "PENDING"
	j.Approved = &approved
	agencyApproved := false
	j.AgencyApprovedCheck = &agencyApproved
	mccusaApproved := false
	j.MccusaApprovedCheck = &mccusaApproved
	return
}

// BeforeUpdate updates the last modified timestamp
func (j *JobApplication) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	j.LastModifiedDate = &now
	return
}
