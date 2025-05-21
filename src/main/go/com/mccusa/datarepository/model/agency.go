package model

import (
	"gorm.io/gorm"
	"time"
)

// Agency maps to the 'agency' table
type Agency struct {
	Email                        string     `gorm:"column:EMAIL;primaryKey" json:"email"`
	Dni                          *string    `gorm:"column:DNI" json:"dni,omitempty"`
	AgencyName                   *string    `gorm:"column:AGENCY_NAME" json:"agencyName,omitempty"`
	AgencyShortName              *string    `gorm:"column:AGENCY_SHORT_NAME" json:"agencyShortName,omitempty"`
	AgencyType                   *string    `gorm:"column:AGENCY_TYPE" json:"agencyType,omitempty"`
	Phone                        *string    `gorm:"column:PHONE" json:"phone,omitempty"`
	Contact                      *string    `gorm:"column:CONTACT" json:"contact,omitempty"`
	Representative               *string    `gorm:"column:REPRESENTATIVE" json:"representative,omitempty"`
	Location                     *string    `gorm:"column:LOCATION" json:"location,omitempty"`
	Tier                         *string    `gorm:"column:TIER" json:"tier,omitempty"`
	Status                       *string    `gorm:"column:STATUS" json:"status,omitempty"`
	Country                      *string    `gorm:"column:COUNTRY" json:"country,omitempty"`
	ClientProfile                *string    `gorm:"column:CLIENT_PROFILE" json:"clientProfile,omitempty"`
	Commission                   *string    `gorm:"column:COMMISSION" json:"commission,omitempty"`
	AgencySize                   *string    `gorm:"column:AGENCY_SIZE" json:"agencySize,omitempty"`
	MainMarketsWithoutLimitation *string    `gorm:"column:MAIN_MARKETS_WITHOUT_LIMITATION" json:"mainMarketsWithoutLimitation,omitempty"`
	HasLinkedClients             bool       `gorm:"column:HAS_LINKED_CLIENTS;not null" json:"hasLinkedClients"`
	ContractType                 *string    `gorm:"column:CONTRACT_TYPE" json:"contractType,omitempty"`
	ContractSignDate             *time.Time `gorm:"column:CONTRACT_SIGN_DATE" json:"contractSignDate,omitempty"`
	LastRenovationDate           *time.Time `gorm:"column:LAST_RENOVATION_DATE" json:"lastRenovationDate,omitempty"`
	AgencyVersion                *string    `gorm:"column:AGENCY_VERSION" json:"agencyVersion,omitempty"`
	CreatedBy                    *string    `gorm:"column:CREATED_BY;size:255" json:"createdBy,omitempty"`
	CreatedDate                  time.Time  `gorm:"column:CREATED_DATE" json:"createdDate"`
	LastModifiedBy               *string    `gorm:"column:LAST_MODIFIED_BY;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate             time.Time  `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate"`
	Deleted                      bool       `gorm:"column:DELETED;not null" json:"deleted"`
	Password                     *string    `gorm:"column:PASSWORD;size:100" json:"password,omitempty"`
	ExtraEmails                  *string    `gorm:"column:EXTRA_EMAILS;size:255" json:"extraEmails,omitempty"`
}

// BeforeCreate will set CreatedDate before inserting
func (a *Agency) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreatedDate = time.Now()
	return
}

// BeforeUpdate will set LastModifiedDate before updating
func (a *Agency) BeforeUpdate(tx *gorm.DB) (err error) {
	a.LastModifiedDate = time.Now()
	return
}
