package model

import (
	"gorm.io/gorm"
	"time"
)

type AgencyClient struct {
	Email            string    `gorm:"primaryKey;column:EMAIL" json:"email"`
	CustomerCode     *string   `gorm:"column:CUSTOMER_CODE" json:"customerCode,omitempty"`
	Passport         *string   `gorm:"column:PASSPORT;uniqueIndex" json:"passport,omitempty"`
	FirstName        string    `gorm:"column:FIRST_NAME;not null" json:"firstName"`
	LastName         string    `gorm:"column:LAST_NAME;not null" json:"lastName"`
	CountryCode      *string   `gorm:"column:COUNTRY_CODE" json:"countryCode,omitempty"`
	Phone            *string   `gorm:"column:PHONE" json:"phone,omitempty"`
	ResidenceCountry *string   `gorm:"column:RESIDENCE_COUNTRY" json:"residenceCountry,omitempty"`
	ResidenceCity    *string   `gorm:"column:RESIDENCE_CITY" json:"residenceCity,omitempty"`
	Agency           *string   `gorm:"column:AGENCY" json:"agency,omitempty"`
	Status           *string   `gorm:"column:STATUS" json:"status,omitempty"`
	Stage            *string   `gorm:"column:STAGE" json:"stage,omitempty"`
	ClientType       *string   `gorm:"column:CLIENT_TYPE" json:"clientType,omitempty"`
	ClientCategory   *string   `gorm:"column:CLIENT_CATEGORY" json:"clientCategory,omitempty"`
	CreatedBy        *string   `gorm:"column:CREATED_BY" json:"createdBy,omitempty"`
	CreatedDate      time.Time `gorm:"column:CREATED_DATE" json:"createdDate"`
	LastModifiedBy   *string   `gorm:"column:LAST_MODIFIED_BY" json:"lastModifiedBy,omitempty"`
	LastModifiedDate time.Time `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate"`
	Deleted          bool      `gorm:"column:DELETED;not null" json:"deleted"`
	Eligible         bool      `gorm:"column:ELIGIBLE;not null" json:"eligible"`
	EmailValidated   bool      `gorm:"column:EMAIL_VALIDATED;not null" json:"emailValidated"`
	UtmSource        *string   `gorm:"column:UTM_SOURCE" json:"utmSource,omitempty"`
	UtmMedium        *string   `gorm:"column:UTM_MEDIUM" json:"utmMedium,omitempty"`
	UtmCampaign      *string   `gorm:"column:UTM_CAMPAIGN" json:"utmCampaign,omitempty"`
	UtmTerm          *string   `gorm:"column:UTM_TERM" json:"utmTerm,omitempty"`
	UtmContent       *string   `gorm:"column:UTM_CONTENT" json:"utmContent,omitempty"`
	ReferralRock     bool      `gorm:"column:UTM_REFERRAL_ROCK;not null" json:"referralRock"`
	ReferralRockType *string   `gorm:"column:UTM_REFERRAL_ROCK_TYPE" json:"referralRockType,omitempty"`
	ReferralRockCode *string   `gorm:"column:UTM_REFERRAL_ROCK_CODE" json:"referralRockCode,omitempty"`
	ReferrerName     *string   `gorm:"column:UTM_REFERFEROR_NAME" json:"referrerName,omitempty"`
	ReferrerEmail    *string   `gorm:"column:UTM_REFERFEROR_EMAIL" json:"referrerEmail,omitempty"`
	Landing          *string   `gorm:"column:LANDING" json:"landing,omitempty"`
	Sponsor          *string   `gorm:"column:SPONSOR" json:"sponsor,omitempty"`
	Vacancy          *string   `gorm:"column:VACANCY" json:"vacancy,omitempty"`
	EnglishLevel     *string   `gorm:"column:ENGLISH_LEVEL" json:"englishLevel,omitempty"`
	CvURL            *string   `gorm:"column:CV_URL" json:"cvUrl,omitempty"`
	SecondName       *string   `gorm:"column:SECOND_NAME" json:"secondName,omitempty"`
	SecondLastName   *string   `gorm:"column:SECOND_LAST_NAME" json:"secondLastName,omitempty"`
	CompletedName    *string   `gorm:"column:COMPLETED_NAME" json:"completedName,omitempty"`
	Round            *int      `gorm:"column:ROUND" json:"round,omitempty"`
	ContractType     *string   `gorm:"column:CONTRACT_TYPE" json:"contractType,omitempty"`
	Gender           *string   `gorm:"column:GENDER" json:"gender,omitempty"`
	Nationality      *string   `gorm:"column:NATIONALITY" json:"nationality,omitempty"`
	DealOwner        *string   `gorm:"column:DEAL_OWNER" json:"dealOwner,omitempty"`
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
