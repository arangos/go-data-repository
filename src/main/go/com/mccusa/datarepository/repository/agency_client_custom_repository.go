package repository

import (
	"context"
	"fmt"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model"
	"go-data-repository/src/main/go/com/mccusa/datarepository/model/dto"
	"go-data-repository/src/main/go/com/mccusa/datarepository/util"

	"github.com/jmoiron/sqlx"
)

// notAllowedSearchStates holds the list of forbidden search states.
var notAllowedSearchStates = []string{
	"PENDING_CLIENTS",
	"COMPLETED_CLIENTS",
	"INVOICED_CLIENTS",
	"ENGAGED_CLIENTS",
}

// AgencyClientCustomRepository handles complex queries not covered by GORM
type AgencyClientCustomRepository struct {
	db    *sqlx.DB
	utils util.ClientUtils
}

// NewAgencyClientCustomRepository creates a new custom repository
func NewAgencyClientCustomRepository(db *sqlx.DB, utils util.ClientUtils) AgencyClientCustomRepository {
	return AgencyClientCustomRepository{db: db, utils: utils}
}

// GetClientsByAgencyEmailApproved fetches clients with pagination and filters
func (r *AgencyClientCustomRepository) GetClientsByAgencyEmailApproved(
	ctx context.Context,
	agency model.Agency,
	approved string,
	searchText string,
	limit, offset int,
) ([]dto.ConsultantsClientsDetails, int, error) {
	// Load base SQL (from QueryLoader equivalent)
	sqlBase := util.GetQuery("findAgencyClientsByAgencyEmail")
	countBase := util.GetQuery("findAgencyClientsByAgencyEmailCount")

	// Prepare parameters
	params := map[string]interface{}{
		"email":  agency.Email,
		"limit":  limit,
		"offset": offset,
	}
	// Apply filters
	if approved != "" && !isAllowed(approved) {
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

	var results []dto.ConsultantsClientsDetails
	for rows.Next() {
		var detail dto.ConsultantsClientsDetails
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

var notAllowedSearchStatesMap = func() map[string]struct{} {
	m := make(map[string]struct{}, len(notAllowedSearchStates))
	for _, s := range notAllowedSearchStates {
		m[s] = struct{}{}
	}
	return m
}()

// Usage:
func isAllowed(state string) bool {
	_, forbidden := notAllowedSearchStatesMap[state]
	return !forbidden
}

// create new method to save the client entity
func (r *AgencyClientCustomRepository) SaveClient(ctx context.Context, client *model.AgencyClient) error {
	if client == nil {
		return fmt.Errorf("client cannot be nil")
	}
	if client.Email == "" {
		return fmt.Errorf("client email is required")
	}

	// Insert the client in the database
	query := `INSERT INTO postgres.client (
		FIRST_NAME, LAST_NAME, SECOND_NAME, SECOND_LAST_NAME, COMPLETED_NAME,
		EMAIL, GENDER, AGENCY, SPONSOR, VACANCY, DELETED, ROUND,
		CONTRACT_TYPE, ELIGIBLE, CLIENT_TYPE, CREATED_DATE, CREATED_BY,
		DEAL_OWNER, CUSTOMER_CODE, PASSPORT
	) VALUES (
		:firstName, :lastName, :secondName, :secondLastName, :completedName,
		:email, :gender, :agency, :sponsor, :vacancy, :deleted, :round,
		:contractType, :eligible, :clientType, now(), 'paginamccusa',
		:dealOwner, :customerCode, :passport
	)`

	_, err := r.db.NamedExecContext(ctx, query, client)
	return err
}
