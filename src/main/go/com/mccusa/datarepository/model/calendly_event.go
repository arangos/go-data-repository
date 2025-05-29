package model

import (
	"gorm.io/gorm"
	"time"
)

// CalendlyEvent maps to the 'calendly_event' table.
type CalendlyEvent struct {
	ID                   uint       `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	InviteeEmail         *string    `gorm:"column:INVITEE_EMAIL" json:"inviteeEmail,omitempty"`
	UserName             *string    `gorm:"column:USER_NAME" json:"userName,omitempty"`
	Team                 *string    `gorm:"column:TEAM" json:"team,omitempty"`
	InviteeName          *string    `gorm:"column:INVITEE_NAME" json:"inviteeName,omitempty"`
	InviteeFirstName     *string    `gorm:"column:INVITEE_FIRST_NAME" json:"inviteeFirstName,omitempty"`
	InviteeLastName      *string    `gorm:"column:INVITEE_LAST_NAME" json:"inviteeLastName,omitempty"`
	EventTypeName        *string    `gorm:"column:EVENT_TYPE_NAME" json:"eventTypeName,omitempty"`
	StartDateTime        *time.Time `gorm:"column:START_DATE_TIME" json:"startDateTime,omitempty"`
	EndDateTime          *time.Time `gorm:"column:END_DATE_TIME" json:"endDateTime,omitempty"`
	EventCreatedDateTime *time.Time `gorm:"column:EVENT_CREATED_DATE_TIME" json:"eventCreatedDateTime,omitempty"`
	Canceled             *bool      `gorm:"column:CANCELED" json:"canceled,omitempty"`
	CanceledBy           *string    `gorm:"column:CANCELED_BY" json:"canceledBy,omitempty"`
	CancellationReason   *string    `gorm:"column:CANCELLATION_REASON" json:"cancellationReason,omitempty"`
	MarkedAsNoShow       *bool      `gorm:"column:MARKED_AS_NO_SHOW" json:"markedAsNoShow,omitempty"`
	InviteeUUID          *string    `gorm:"column:INVITEE_UUID" json:"inviteeUuid,omitempty"`
	EventUUID            *string    `gorm:"column:EVENT_UUID" json:"eventUuid,omitempty"`
	LastModifiedBy       *string    `gorm:"column:LAST_MODIFIED_BY;size:255" json:"lastModifiedBy,omitempty"`
	LastModifiedDate     *time.Time `gorm:"column:LAST_MODIFIED_DATE" json:"lastModifiedDate,omitempty"`
}

// BeforeUpdate updates last modified timestamps before saving changes.
func (e *CalendlyEvent) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	e.LastModifiedDate = &now
	// LastModifiedBy should be set via context or a util if you have user info
	return
}
