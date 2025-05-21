package model

import (
	"gorm.io/gorm"
	"time"
)

// Contract maps to the 'contract' table
type Contract struct {
	ID                  uint           `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	SenderEmail         *string        `gorm:"column:SENDER_EMAIL;size:255" json:"senderEmail,omitempty"`
	RecipientEmail      *string        `gorm:"column:RECIPIENT_EMAIL;size:255" json:"recipientEmail,omitempty"`
	ContractName        *string        `gorm:"column:CONTRACT_NAME;size:255" json:"contractName,omitempty"`
	ContractID          *string        `gorm:"column:CONTRACT_ID;size:255" json:"contractId,omitempty"`
	CreationDate        time.Time      `gorm:"column:CREATION_DATE" json:"creationDate"`
	CreatedBy           *string        `gorm:"column:CREATED_BY;size:255" json:"createdBy,omitempty"`
	ContractURL         *string        `gorm:"column:CONTRACT_URL" json:"contractUrl,omitempty"`
	ContractSponsor     *string        `gorm:"column:CONTRACT_SPONSOR" json:"contractSponsor,omitempty"`
	ContractVacancy     *string        `gorm:"column:CONTRACT_VACANCY" json:"contractVacancy,omitempty"`
	ContractData        []JsonMetadata `gorm:"column:CONTRACT_DATA;type:json" json:"contractData,omitempty"`
	AWSDocuments        []JsonMetadata `gorm:"column:AWS_DOCUMENTS;type:json" json:"awsDocuments,omitempty"`
	LastModifiedBy      *string        `gorm:"column:LAST_MODIFIED_BY;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate    *time.Time     `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate,omitempty"`
	MccusaApprovedCheck *bool          `gorm:"column:MCCUSA_APPROVED_CHECK" json:"mccusaApprovedCheck,omitempty"`
	ContractAgency      *string        `gorm:"column:CONTRACT_AGENCY" json:"contractAgency,omitempty"`
	AgencyApprovedCheck *bool          `gorm:"column:AGENCY_APPROVED_CHECK" json:"agencyApprovedCheck,omitempty"`
	CustomerCode        *string        `gorm:"column:CUSTOMER_CODE" json:"customerCode,omitempty"`
	JobApplicationID    *int           `gorm:"column:JOB_APPLICATION_ID" json:"jobApplicationId,omitempty"`
	ContractType        *string        `gorm:"column:CONTRACT_TYPE" json:"contractType,omitempty"`
}

// BeforeCreate sets defaults before inserting a new contract
func (c *Contract) BeforeCreate(tx *gorm.DB) (err error) {
	c.CreationDate = time.Now()
	// Defaults for approval checks
	defaultBool := false
	c.MccusaApprovedCheck = &defaultBool
	c.AgencyApprovedCheck = &defaultBool
	return
}

// BeforeUpdate updates last modified timestamp
func (c *Contract) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	c.LastModifiedDate = &now
	return
}
