package repository

import (
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/constants"
	"gorm.io/gorm"
)

type DocusignEnvelopesRepository interface {
	// FindAllByEnvelopeStatusAndClientEmailIsNotNull
	FindAllByEnvelopeStatusAndClientEmailIsNotNull(envelopeStatus constants.EnvelopeStatus) ([]model.DocusignEnvelope, error)
}

type docusignEnvelopesRepo struct {
	db *gorm.DB
}

// NewDocusignEnvelopesRepository constructs a new DocusignEnvelopesRepository.
func NewDocusignEnvelopesRepository(db *gorm.DB) DocusignEnvelopesRepository {
	return &docusignEnvelopesRepo{db: db}
}

func (r *docusignEnvelopesRepo) FindAllByEnvelopeStatusAndClientEmailIsNotNull(envelopeStatus constants.EnvelopeStatus) ([]model.DocusignEnvelope, error) {
	var envelopes []model.DocusignEnvelope
	err := r.db.Model(&model.DocusignEnvelope{}).
		Where("envelope_status = ? AND client_email IS NOT NULL", envelopeStatus).
		Find(&envelopes).Error
	if err != nil {
		return nil, err
	}
	return envelopes, nil
}
