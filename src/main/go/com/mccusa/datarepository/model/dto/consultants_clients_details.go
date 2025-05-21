package dto

import (
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
)

// ConsultantsClientsDetails aggregates client details, job applications, contracts, form data, and invoice info.
type ConsultantsClientsDetails struct {
	Client          model.AgencyClient                 `json:"client"`
	JobApplications []model.JobApplication             `json:"jobApplications"`
	Contracts       []model.Contract                   `json:"contracts"`
	ClientForm      model.ClientFormConsultantResponse `json:"clientForm"`
	Invoice         model.Invoice                      `json:"invoice"`
}
