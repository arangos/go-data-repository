package repository

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"
)

// AgencyClientCustomRepository handles complex queries not covered by GORM
type AgencyClientCustomRepository struct {
	db    *sqlx.DB
	utils util.ClientUtils
}

// NewAgencyClientCustomRepository creates a new custom repository
func NewAgencyClientCustomRepository(db *sqlx.DB, utils util.ClientUtils) *AgencyClientCustomRepository {
	return &AgencyClientCustomRepository{db: db, utils: utils}
}

// GetClientsByAgencyEmailApproved fetches clients with pagination and filters
func (r *AgencyClientCustomRepository) GetClientsByAgencyEmailApproved(
	ctx context.Context,
	agency model.Agency,
	approved string,
	searchText string,
	limit, offset int,
) ([]model.ConsultantsClientsDetails, int, error) {
	// Load base SQL (from QueryLoader equivalent)
	sqlBase := util.LoadQuery("findAgencyClientsByAgencyEmail", agency.AgencyType)
	countBase := util.LoadQuery("findAgencyClientsByAgencyEmailCount", agency.AgencyType)

	// Prepare parameters
	params := map[string]interface{}{
		"email":  agency.Email,
		"limit":  limit,
		"offset": offset,
	}
	// Apply filters
	if approved != "" && !util.NotAllowedSearchStates[approved] {
		params["approved"] = approved
	}
	if searchText != "" {
		params["searchText"] = "%" + searchText + "%"
	}

	// Execute query
	rows, err := r.db.NamedQueryContext(ctx, sqlBase, params)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var results []model.ConsultantsClientsDetails
	for rows.Next() {
		var detail model.ConsultantsClientsDetails
		// TODO: scan required fields
		// rows.StructScan(&detail)
		results = append(results, detail)
	}

	// Count total
	var total int
	countRow := r.db.QueryRowxContext(ctx, countBase, params)
	if err := countRow.Scan(&total); err != nil {
		return nil, 0, fmt.Errorf("count query error: %w", err)
	}

	return results, total, nil
}
