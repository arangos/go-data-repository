package model

import (
	"gorm.io/gorm"
	"time"
)

// Sponsor maps to the 'sponsor' table.
type Sponsor struct {
	ID                               uint       `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	IDN                              *string    `gorm:"column:IDN;size:255" json:"idn,omitempty"`
	SponsorName                      *string    `gorm:"column:SPONSOR_NAME;size:255" json:"sponsorName,omitempty"`
	SponsorLocation                  *string    `gorm:"column:SPONSOR_LOCATION;size:255" json:"sponsorLocation,omitempty"`
	SponsorContact                   *string    `gorm:"column:SPONSOR_CONTACT;size:255" json:"sponsorContact,omitempty"`
	SponsorEmail                     string     `gorm:"column:SPONSOR_EMAIL" json:"sponsorEmail,omitempty"`
	SponsorLogin                     *string    `gorm:"column:SPONSOR_LOGIN" json:"sponsorLogin,omitempty"`
	SponsorPassword                  *string    `gorm:"column:SPONSOR_PASSWORD" json:"-"`
	ApplicationWindow                *time.Time `gorm:"column:APPLICATION_WINDOW" json:"applicationWindow,omitempty"`
	Deleted                          bool       `gorm:"column:DELETED;not null" json:"deleted"`
	CreatedBy                        *string    `gorm:"column:CREATED_BY;size:255" json:"createdBy,omitempty"`
	CreatedDate                      *time.Time `gorm:"column:CREATED_DATE" json:"createdDate,omitempty"`
	LastModifiedBy                   *string    `gorm:"column:LAST_MODIFIED_BY;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate                 *time.Time `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate,omitempty"`
	SponsorJobApplicationDescription *string    `gorm:"column:SPONSOR_JOB_APPLICATION_DESCRIPTION" json:"sponsorJobApplicationDescription,omitempty"`
}

// BeforeCreate sets CreatedDate just before inserting a new record.
func (s *Sponsor) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	s.CreatedDate = &now
	// s.CreatedBy should be set via context or a util if you have authenticated user info
	return
}

// BeforeUpdate sets LastModifiedDate just before updating.
func (s *Sponsor) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	s.LastModifiedDate = &now
	return
}
