package model

import (
	"gorm.io/gorm"
	"time"
)

func (DocusignEnvelope) TableName() string {
	return "postgres.docusign_envelope"
}

type DocusignEnvelope struct {
	EnvelopeID       string    `json:"envelope_id" gorm:"primaryKey;column:envelope_id"`
	EnvelopeStatus   *string   `json:"envelope_status" gorm:"column:envelope_status"`
	EnvelopeName     *string   `json:"envelope_name" gorm:"column:envelope_name"`
	ClientEmail      *string   `json:"client_email" gorm:"column:client_email"`
	CreatedDate      time.Time `json:"created_date" gorm:"column:created_date"`
	CreatedBy        *string   `json:"created_by" gorm:"column:created_by"`
	LastModifiedBy   *string   `json:"last_modified_by" gorm:"column:last_modified_by"`
	LastModifiedDate time.Time `json:"last_modified_date" gorm:"column:last_modified_date"`
}

// BeforeCreate will set CreatedDate before inserting
func (docusignEnvelope *DocusignEnvelope) BeforeCreate(tx *gorm.DB) (err error) {
	docusignEnvelope.CreatedDate = time.Now()
	user := "paginamccusa"
	docusignEnvelope.CreatedBy = &user
	return
}

// BeforeUpdate will set LastModifiedDate before updating
func (docusignEnvelope *DocusignEnvelope) BeforeUpdate(tx *gorm.DB) (err error) {
	docusignEnvelope.LastModifiedDate = time.Now()
	user := "paginamccusa"
	docusignEnvelope.LastModifiedBy = &user
	return
}
