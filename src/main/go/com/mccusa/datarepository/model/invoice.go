package model

import (
	"gorm.io/gorm"
	"time"
)

// Invoice maps to the 'invoice' table
// includes payment details and audit fields

type Invoice struct {
	ID                    uint       `gorm:"column:ID;primaryKey;autoIncrement" json:"id"`
	InvoicePaymentDate    *time.Time `gorm:"column:invoice_payment_date" json:"invoicePaymentDate,omitempty"`
	PaymentInvoice        *string    `gorm:"column:payment_invoice;uniqueIndex" json:"paymentInvoice,omitempty"`
	InvoicePaymentType    *string    `gorm:"column:invoice_payment_type" json:"invoicePaymentType,omitempty"`
	InvoicePaymentComment *string    `gorm:"column:invoice_payment_comment" json:"invoicePaymentComment,omitempty"`
	InvoiceAmountPaid     *int64     `gorm:"column:invoice_amount_paid" json:"invoiceAmountPaid,omitempty"`
	CustomerCode          *string    `gorm:"column:customer_code" json:"customerCode,omitempty"`
	JobApplicationID      *int       `gorm:"column:job_application_id" json:"jobApplicationId,omitempty"`
	InvoiceCreatedDate    *time.Time `gorm:"column:invoice_created_date" json:"invoiceCreatedDate,omitempty"`
	InvoiceCreatedBy      *string    `gorm:"column:invoice_created_by" json:"invoiceCreatedBy,omitempty"`
	LastModifiedBy        *string    `gorm:"column:last_modified_by" json:"lastModifiedBy,omitempty"`
	LastModifiedDate      *time.Time `gorm:"column:last_modified_date" json:"lastModifiedDate,omitempty"`
}

// BeforeCreate sets the createdBy and createdDate before insert
func (i *Invoice) BeforeCreate(tx *gorm.DB) (err error) {
	now := time.Now()
	i.InvoiceCreatedDate = &now
	// createdBy should be injected via context or util
	return
}

// BeforeUpdate updates lastModifiedDate before update
func (i *Invoice) BeforeUpdate(tx *gorm.DB) (err error) {
	now := time.Now()
	i.LastModifiedDate = &now
	return
}
