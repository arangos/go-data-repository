package dto

import "time"

// ClientFormConsultantResponse mirrors the Java DTO for client form survey results
type ClientFormConsultantResponse struct {
	ClientFormID          int64     `json:"client_form_id"`
	ClientFormCreatedBy   string    `json:"client_form_created_by"`
	ClientFormCreatedDate time.Time `json:"client_form_created_date"`
	MccusaApprovedCheck   bool      `json:"mccusaApprovedCheck"`
	AgencyApprovedCheck   bool      `json:"agencyApprovedCheck"`
}
