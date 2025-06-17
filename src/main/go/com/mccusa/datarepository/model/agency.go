package model

import (
	"gorm.io/gorm"
	"time"
)

func (Agency) TableName() string {
	return "postgres.agency"
}

type Agency struct {
	Email                        string    `gorm:"column:email;primaryKey" json:"email"`
	Dni                          *string   `gorm:"column:dni" json:"dni,omitempty"`
	AgencyName                   *string   `gorm:"column:agency_name" json:"agencyName,omitempty"`
	AgencyShortName              *string   `gorm:"column:agency_short_name" json:"agencyShortName,omitempty"`
	AgencyType                   *string   `gorm:"column:agency_type" json:"agencyType,omitempty"`
	Phone                        *string   `gorm:"column:phone" json:"phone,omitempty"`
	Contact                      *string   `gorm:"column:contact" json:"contact,omitempty"`
	Representative               *string   `gorm:"column:representative" json:"representative,omitempty"`
	Location                     *string   `gorm:"column:location" json:"location,omitempty"`
	Tier                         *string   `gorm:"column:tier" json:"tier,omitempty"`
	Status                       *string   `gorm:"column:status" json:"status,omitempty"`
	Country                      *string   `gorm:"column:country" json:"country,omitempty"`
	ClientProfile                *string   `gorm:"column:client_profile" json:"clientProfile,omitempty"`
	Commission                   *string   `gorm:"column:commission" json:"commission,omitempty"`
	AgencySize                   *string   `gorm:"column:agency_size" json:"agencySize,omitempty"`
	MainMarketsWithoutLimitation *string   `gorm:"column:main_markets_without_limitation" json:"mainMarketsWithoutLimitation,omitempty"`
	HasLinkedClients             bool      `gorm:"column:has_linked_clients;not null" json:"hasLinkedClients"`
	ContractType                 *string   `gorm:"column:contract_type" json:"contractType,omitempty"`
	ContractSignDate             time.Time `gorm:"column:contract_sign_date" json:"contractSignDate,omitempty"`
	LastRenovationDate           time.Time `gorm:"column:last_renovation_date" json:"lastRenovationDate,omitempty"`
	AgencyVersion                *string   `gorm:"column:agency_version" json:"agencyVersion,omitempty"`
	CreatedBy                    *string   `gorm:"column:created_by;size:255" json:"createdBy,omitempty"`
	CreatedDate                  time.Time `gorm:"column:created_date" json:"createdDate"`
	LastModifiedBy               *string   `gorm:"column:last_modified_by;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate             time.Time `gorm:"column:last_modified_date" json:"lastModifiedDate"`
	Deleted                      *bool     `gorm:"column:deleted;not null" json:"deleted"`
	Password                     *string   `gorm:"column:password;size:100" json:"password,omitempty"`
	ExtraEmails                  *string   `gorm:"column:extra_emails;size:255" json:"extraEmails,omitempty"`
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
