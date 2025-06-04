package user

import (
	"context"
	"blogging/base"
	"blogging/services"

	"github.com/georgysavva/scany/v2/sqlscan"
	"github.com/rs/zerolog/log"
)

func GetAllUsers(ctx context.Context, db *base.DB, qb *services.QueryBuilder) (GetAllUsersResponse, error) {
	var response GetAllUsersResponse

	query, args, err := qb.Build()
	if err != nil {
		log.Err(err).Msg("Failed to build query: [GetAllUsers] [Data]")
		return response, err
	}

	// Fetch user list
	if err := sqlscan.Select(ctx, db, &response.Users, query, args...); err != nil {
		log.Err(err).Msg("Failed to execute user list query")
		return response, err
	}

	// Count total
	queryCount, argsCount, err := qb.Count()
	if err != nil {
		log.Err(err).Msg("Failed to build count query: [GetAllUsers]")
		return response, err
	}

	if err := sqlscan.Get(ctx, db, &response.Total, queryCount, argsCount...); err != nil {
		log.Err(err).Msg("Failed to get total user count")
		return response, err
	}

	return response, nil
}
