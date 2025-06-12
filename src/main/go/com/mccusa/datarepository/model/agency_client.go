package model

import (
	"gorm.io/gorm"
	"time"
)

// TableName specifies the table name for GORM
func (AgencyClient) TableName() string {
	return "postgres.client"
}

type AgencyClient struct {
	Email            string    `gorm:"primaryKey;column:email" json:"email"`
	CustomerCode     *string   `gorm:"column:customer_code" json:"customerCode,omitempty"`
	Passport         *string   `gorm:"column:passport;uniqueIndex" json:"passport,omitempty"`
	FirstName        string    `gorm:"column:first_name;not null" json:"firstName"`
	LastName         string    `gorm:"column:last_name;not null" json:"lastName"`
	CountryCode      *string   `gorm:"column:country_code" json:"countryCode,omitempty"`
	Phone            *string   `gorm:"column:phone" json:"phone,omitempty"`
	ResidenceCountry *string   `gorm:"column:residence_country" json:"residenceCountry,omitempty"`
	ResidenceCity    *string   `gorm:"column:residence_city" json:"residenceCity,omitempty"`
	Agency           *string   `gorm:"column:agency" json:"agency,omitempty"`
	Status           *string   `gorm:"column:status" json:"status,omitempty"`
	Stage            *string   `gorm:"column:stage" json:"stage,omitempty"`
	ClientType       *string   `gorm:"column:client_type" json:"clientType,omitempty"`
	ClientCategory   *string   `gorm:"column:client_category" json:"clientCategory,omitempty"`
	CreatedBy        *string   `gorm:"column:created_by" json:"createdBy,omitempty"`
	CreatedDate      time.Time `gorm:"column:created_date" json:"createdDate"`
	LastModifiedBy   *string   `gorm:"column:last_modified_by" json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time `gorm:"column:last_modified_date" json:"lastModifiedDate"`
	Deleted          bool      `gorm:"column:deleted;not null" json:"deleted"`
	Eligible         bool      `gorm:"column:eligible;not null" json:"eligible"`
	EmailValidated   bool      `gorm:"column:email_validated;not null" json:"emailValidated"`
	UtmSource        *string   `gorm:"column:utm_source" json:"utmSource,omitempty"`
	UtmMedium        *string   `gorm:"column:utm_medium" json:"utmMedium,omitempty"`
	UtmCampaign      *string   `gorm:"column:utm_campaign" json:"utmCampaign,omitempty"`
	UtmTerm          *string   `gorm:"column:utm_term" json:"utmTerm,omitempty"`
	UtmContent       *string   `gorm:"column:utm_content" json:"utmContent,omitempty"`
	ReferralRock     bool      `gorm:"column:utm_referral_rock;not null" json:"referralRock"`
	ReferralRockType *string   `gorm:"column:utm_referral_rock_type" json:"referralRockType,omitempty"`
	ReferralRockCode *string   `gorm:"column:utm_referral_rock_code" json:"referralRockCode,omitempty"`
	ReferrerName     *string   `gorm:"column:utm_referferor_name" json:"referrerName,omitempty"`
	ReferrerEmail    *string   `gorm:"column:utm_referferor_email" json:"referrerEmail,omitempty"`
	Landing          *string   `gorm:"column:landing" json:"landing,omitempty"`
	Sponsor          *string   `gorm:"column:sponsor" json:"sponsor,omitempty"`
	Vacancy          *string   `gorm:"column:vacancy" json:"vacancy,omitempty"`
	EnglishLevel     *string   `gorm:"column:english_level" json:"englishLevel,omitempty"`
	CvURL            *string   `gorm:"column:cv_url" json:"cvUrl,omitempty"`
	SecondName       *string   `gorm:"column:second_name" json:"secondName,omitempty"`
	SecondLastName   *string   `gorm:"column:second_last_name" json:"secondLastName,omitempty"`
	CompletedName    *string   `gorm:"column:completed_name" json:"completedName,omitempty"`
	Round            *int      `gorm:"column:round" json:"round,omitempty"`
	ContractType     *string   `gorm:"column:contract_type" json:"contractType,omitempty"`
	Gender           *string   `gorm:"column:gender" json:"gender,omitempty"`
	Nationality      *string   `gorm:"column:nationality" json:"nationality,omitempty"`
	DealOwner        *string   `gorm:"column:deal_owner" json:"dealOwner,omitempty"`
}

// BeforeCreate sets CreatedDate before inserting
func (a *AgencyClient) BeforeCreate(tx *gorm.DB) (err error) {
	a.CreatedDate = time.Now()
	return
}

// BeforeUpdate sets LastModifiedDate before updating
func (a *AgencyClient) BeforeUpdate(tx *gorm.DB) (err error) {
	a.LastModifiedDate = time.Now()
	return
}
