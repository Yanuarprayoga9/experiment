package services

import (
	"fmt"
	"strconv"
	"strings"
)

// QueryBuilder helps build SQL queries with filtering, sorting, and pagination
type QueryBuilder struct {
	TableName     string
	SelectFields  []string
	AllowedSorts  []string
	Filters       map[string]interface{}
	Sort          string
	Order         string
	Page          int
	Limit         int
	Search        string
	SearchFields  []string
}

// NewQueryBuilder creates a new QueryBuilder instance
func NewQueryBuilder(tableName string, selectFields ...string) *QueryBuilder {
	return &QueryBuilder{
		TableName:    tableName,
		SelectFields: selectFields,
		Filters:      make(map[string]interface{}),
		Page:         1,
		Limit:        10,
		Order:        "ASC",
	}
}

// AllowSort sets the allowed sort fields
func (qb *QueryBuilder) AllowSort(fields ...string) *QueryBuilder {
	qb.AllowedSorts = fields
	return qb
}

// AddFilter adds a filter condition
func (qb *QueryBuilder) AddFilter(field string, value interface{}) *QueryBuilder {
	qb.Filters[field] = value
	return qb
}

// SetSearch sets search term and searchable fields
func (qb *QueryBuilder) SetSearch(term string, fields ...string) *QueryBuilder {
	qb.Search = term
	qb.SearchFields = fields
	return qb
}

// ApplyFromQuery applies query parameters from HTTP request
func (qb *QueryBuilder) ApplyFromQuery(queries map[string]string) *QueryBuilder {
	// Handle pagination
	if page, ok := queries["page"]; ok {
		if p, err := strconv.Atoi(page); err == nil && p > 0 {
			qb.Page = p
		}
	}

	if limit, ok := queries["limit"]; ok {
		if l, err := strconv.Atoi(limit); err == nil && l > 0 && l <= 100 {
			qb.Limit = l
		}
	}

	// Handle sorting
	if sort, ok := queries["sort"]; ok {
		qb.Sort = sort
	}

	if order, ok := queries["order"]; ok {
		if strings.ToUpper(order) == "DESC" {
			qb.Order = "DESC"
		}
	}

	// Handle search
	if search, ok := queries["search"]; ok {
		qb.Search = search
	}

	// Handle specific filters
	for key, value := range queries {
		if !isReservedParam(key) {
			qb.AddFilter(key, value)
		}
	}

	return qb
}

// BuildSelectQuery builds the SELECT query
func (qb *QueryBuilder) BuildSelectQuery() (string, []interface{}) {
	var args []interface{}
	argIndex := 1

	// SELECT clause
	selectClause := "*"
	if len(qb.SelectFields) > 0 {
		selectClause = strings.Join(qb.SelectFields, ", ")
	}

	query := fmt.Sprintf("SELECT %s FROM %s", selectClause, qb.TableName)

	// WHERE clause
	var whereConditions []string

	// Add filters
	for field, value := range qb.Filters {
		whereConditions = append(whereConditions, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	// Add search conditions
	if qb.Search != "" && len(qb.SearchFields) > 0 {
		var searchConditions []string
		searchTerm := "%" + qb.Search + "%"
		for _, field := range qb.SearchFields {
			searchConditions = append(searchConditions, fmt.Sprintf("%s LIKE $%d", field, argIndex))
			args = append(args, searchTerm)
			argIndex++
		}
		if len(searchConditions) > 0 {
			whereConditions = append(whereConditions, "("+strings.Join(searchConditions, " OR ")+")")
		}
	}

	// Add soft delete filter
	whereConditions = append(whereConditions, "DeletedAt IS NULL")

	if len(whereConditions) > 0 {
		query += " WHERE " + strings.Join(whereConditions, " AND ")
	}

	// ORDER BY clause
	if qb.Sort != "" && qb.isAllowedSort(qb.Sort) {
		query += fmt.Sprintf(" ORDER BY %s %s", qb.Sort, qb.Order)
	}

	// OFFSET and LIMIT for pagination
	offset := (qb.Page - 1) * qb.Limit
	query += fmt.Sprintf(" OFFSET %d ROWS FETCH NEXT %d ROWS ONLY", offset, qb.Limit)

	return query, args
}

// BuildCountQuery builds the COUNT query for pagination
func (qb *QueryBuilder) BuildCountQuery() (string, []interface{}) {
	var args []interface{}
	argIndex := 1

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", qb.TableName)

	// WHERE clause (same as select query but without pagination)
	var whereConditions []string

	// Add filters
	for field, value := range qb.Filters {
		whereConditions = append(whereConditions, fmt.Sprintf("%s = $%d", field, argIndex))
		args = append(args, value)
		argIndex++
	}

	// Add search conditions
	if qb.Search != "" && len(qb.SearchFields) > 0 {
		var searchConditions []string
		searchTerm := "%" + qb.Search + "%"
		for _, field := range qb.SearchFields {
			searchConditions = append(searchConditions, fmt.Sprintf("%s LIKE $%d", field, argIndex))
			args = append(args, searchTerm)
			argIndex++
		}
		if len(searchConditions) > 0 {
			whereConditions = append(whereConditions, "("+strings.Join(searchConditions, " OR ")+")")
		}
	}

	// Add soft delete filter
	whereConditions = append(whereConditions, "DeletedAt IS NULL")

	if len(whereConditions) > 0 {
		query += " WHERE " + strings.Join(whereConditions, " AND ")
	}

	return query, args
}

// isAllowedSort checks if the sort field is allowed
func (qb *QueryBuilder) isAllowedSort(field string) bool {
	for _, allowed := range qb.AllowedSorts {
		if strings.ToLower(allowed) == strings.ToLower(field) {
			return true
		}
	}
	return false
}

// isReservedParam checks if the parameter is reserved for query building
func isReservedParam(param string) bool {
	reserved := []string{"page", "limit", "sort", "order", "search"}
	for _, r := range reserved {
		if r == param {
			return true
		}
	}
	return false
}