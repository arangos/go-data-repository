package util

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// StatusTitles maps ActiveCampaign status codes to user-friendly strings.
var StatusTitles = map[int]string{
	0: "OPEN",
	1: "WON",
	2: "LOST",
}

// StageTitles maps ActiveCampaign stage IDs to descriptive titles.
var StageTitles = map[int]string{
	53:  "New Lead",
	54:  "Automatic lead validation",
	6:   "Discovery",
	142: "Getting the Money",
	52:  "Rescheduling",
	4:   "Meeting",
	5:   "Ready to Apply",
	3:   "Received Job Application",
	27:  "Approved by Sponsor",
	112: "Ready to Sign Contract",
	115: "Automatic lead validation",
	49:  "PRE ETA 9089",
	40:  "ETA 9089 IN PROGRESS",
	43:  "ETA 9089 AUDITS",
	39:  "ETA 9089 APPROVALS",
	50:  "I-140 IN PROGRESS",
	42:  "I-140 AUDITS",
	44:  "I-140 APPROVALS",
	45:  "DS-260 IN PROGRESS",
	86:  "DS-260 D Complete",
	46:  "I-485 IN PROGRESS",
	47:  "I-765 APPROVALS",
	48:  "CHAMPION",
	110: "Future Interest",
	109: "Followup",
	28:  "Ready to Sign Contract",
	185: "Contract signed and waiting for invoice",
	186: "Contract invoiced and waiting payment",
}

// PopulateField creates the field value object for ActiveCampaign API
func PopulateField(contactID int64, url string, fieldID int) map[string]any {
	return map[string]any{
		"contact": contactID,
		"field":   fieldID,
		"value":   url,
	}
}

// PrepareRequestBody creates the complete request body with field value and defaults
func PrepareRequestBody(fieldValue map[string]any) []byte {
	data := map[string]any{
		"fieldValue":  fieldValue,
		"useDefaults": true,
	}
	body, _ := json.Marshal(data)
	return body
}

// PrepareRequest creates an HTTP request for ActiveCampaign API
func PrepareRequest(body []byte) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, "https://mccusa.api-us1.com/api/3/fieldValues", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	return req
}
