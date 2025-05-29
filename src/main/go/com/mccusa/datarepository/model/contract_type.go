package model

import (
	"gorm.io/gorm"
	"time"
)

// ContractType maps to the 'contract_types' table.
type ContractType struct {
	ID                     uint      `gorm:"column:id;primaryKey;autoIncrement" json:"id"`
	ContractType           string    `gorm:"column:contract_type;unique;not null" json:"contractType"`
	ContractName           string    `gorm:"column:contract_name;not null" json:"contractName"`
	DocusignTemplateID     string    `gorm:"column:docusign_template_id;not null" json:"docusignTemplateId"`
	ContractPrice          *int64    `gorm:"column:contract_price" json:"contractPrice,omitempty"`
	CreatedDate            time.Time `gorm:"column:created_date;not null" json:"createdDate"`
	CreatedBy              string    `gorm:"column:created_by;not null" json:"createdBy"`
	Comments               *string   `gorm:"column:comments" json:"comments,omitempty"`
	Installment1           *int64    `gorm:"column:installment1" json:"installment1,omitempty"`
	Installment2           *int64    `gorm:"column:installment2" json:"installment2,omitempty"`
	Installment3           *int64    `gorm:"column:installment3" json:"installment3,omitempty"`
	Installment4           *int64    `gorm:"column:installment4" json:"installment4,omitempty"`
	Status                 *string   `gorm:"column:status" json:"status,omitempty"`
	ActiveForAgencies      *bool     `gorm:"column:active_for_agencies" json:"activeForAgencies,omitempty"`
	ActiveForDirectClients *bool     `gorm:"column:active_for_direct_clients" json:"activeForDirectClients,omitempty"`
}

// BeforeCreate sets the created timestamp and creator before insert.
func (ct *ContractType) BeforeCreate(tx *gorm.DB) (err error) {
	ct.CreatedDate = time.Now()
	// ct.CreatedBy should be set via context or util (e.g., current user)
	return
}
