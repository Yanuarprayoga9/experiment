package services

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	sq "github.com/Masterminds/squirrel"
)

type QueryBuilder struct {
	builder         sq.SelectBuilder
	selectColumns   []string
	allowedSorts    map[string]bool
	allowedFilters  map[string]bool
	sortBy          string
	order           string
	limit           uint64
	offset          uint64
	fromTable       string
	whereConditions []sq.Sqlizer
	debug           bool
	customColumn    map[string]sq.SelectBuilder
}

// Regex to validate safe SQL identifiers
var safeAliasRegex = regexp.MustCompile(`^[a-zA-Z_][a-zA-Z0-9_]*$`)

func NewQueryBuilder(table string, columns ...string) *QueryBuilder {
	var builder sq.SelectBuilder
	if len(columns) == 0 {
		builder = sq.Select("*").From(table).PlaceholderFormat(sq.AtP)
	} else {
		builder = sq.Select(columns...).From(table).PlaceholderFormat(sq.AtP)
	}

	return &QueryBuilder{
		builder:        builder,
		selectColumns:  columns,
		allowedSorts:   map[string]bool{},
		allowedFilters: map[string]bool{},
		fromTable:      table,
		limit:          50,
		offset:         0,
		customColumn:   map[string]sq.SelectBuilder{},
	}
}

// AllowSort defines safe columns that can be used in ORDER BY
func (qb *QueryBuilder) AllowSort(fields ...string) {
	for _, f := range fields {
		qb.allowedSorts[f] = true
	}
}

// AllowFilter defines safe columns that can be used in WHERE
func (qb *QueryBuilder) AllowFilter(fields ...string) {
	for _, f := range fields {
		qb.allowedFilters[f] = true
	}
}

// AddCustomColumn adds a safe subquery column to the SELECT list
func (qb *QueryBuilder) AddCustomColumn(alias string, subquery sq.SelectBuilder) {
	qb.customColumn[alias] = subquery
}

// Debug enables debug output during Build and Count
func (qb *QueryBuilder) Debug() {
	qb.debug = true
}

// Join adds a JOIN clause
func (qb *QueryBuilder) Join(table string, args ...interface{}) {
	qb.builder = qb.builder.Join(table, args...)
}

// Where adds a custom condition manually
func (qb *QueryBuilder) Where(cond sq.Sqlizer) {
	qb.whereConditions = append(qb.whereConditions, cond)
	qb.builder = qb.builder.Where(cond)
}

// SortBy sets a sort field (must be whitelisted)
func (qb *QueryBuilder) SortBy(field string) {
	qb.sortBy = field
}

// ApplyFromQuery parses a map of query params and updates filters, sorts, limit/offset
func (qb *QueryBuilder) ApplyFromQuery(params map[string]string) {
	for key, val := range params {
		if key == "sortBy" || key == "order" || key == "limit" || key == "offset" || key == "page" {
			continue
		}

		// if !validateColumnName(key, qb.allowedFilters) {
		// 	continue
		// }

		switch {
		case strings.HasPrefix(val, "like:"):
			cond := sq.Like{key: "%" + strings.TrimPrefix(val, "like:") + "%"}
			qb.whereConditions = append(qb.whereConditions, cond)
			qb.builder = qb.builder.Where(cond)

		case strings.HasPrefix(val, "ilike:"):
			keyValues := strings.Split(key, ",")
			if len(keyValues) > 1 {
				orConds := make(sq.Or, 0)
				for _, k := range keyValues {
					orConds = append(orConds, sq.Expr(fmt.Sprintf("%s LIKE ? COLLATE SQL_Latin1_General_CP1_CI_AS", k), "%"+strings.TrimPrefix(val, "ilike:")+"%"))
				}
				qb.whereConditions = append(qb.whereConditions, orConds)
				qb.builder = qb.builder.Where(orConds)
			} else {
				cond := sq.Expr(fmt.Sprintf("%s LIKE ? COLLATE SQL_Latin1_General_CP1_CI_AS", key),
					"%"+strings.TrimPrefix(val, "ilike:")+"%")
				qb.whereConditions = append(qb.whereConditions, cond)
				qb.builder = qb.builder.Where(cond)
			}

		case strings.HasPrefix(val, "or:"):
			orValues := strings.Split(strings.TrimPrefix(val, "or:"), ",")
			orConds := make(sq.Or, 0, len(orValues))
			for _, v := range orValues {
				orConds = append(orConds, sq.Eq{key: v})
			}
			qb.whereConditions = append(qb.whereConditions, orConds)
			qb.builder = qb.builder.Where(orConds)
		case strings.HasPrefix(val, "not:"):
			cond := sq.NotEq{key: strings.TrimPrefix(val, "not:")}
			qb.whereConditions = append(qb.whereConditions, cond)
			qb.builder = qb.builder.Where(cond)
		default:
			if qb.fromTable == "users" && key == "attributes" {
				KeyValues := strings.Split(val, ";")
				qb.builder = qb.builder.Join("UsersAttributes ua ON ua.UserId = users.Id")
				qb.builder = qb.builder.Where(sq.Eq{"ua.AttributeKey": KeyValues[0], "ua.AttributeValue": KeyValues[1]})
			} else {

				cond := sq.Eq{key: val}
				qb.whereConditions = append(qb.whereConditions, cond)
				qb.builder = qb.builder.Where(cond)
			}
		}

	}
	var limit int
	qb.sortBy = params["sortBy"]
	qb.order = strings.ToUpper(params["order"])

	// if !validateColumnName(qb.sortBy, qb.allowedSorts) {
	// 	qb.sortBy = ""
	// }

	if qb.order != "DESC" {
		qb.order = "ASC"
	}

	if l := params["limit"]; l != "" {
		if n, err := strconv.Atoi(l); err == nil {
			limit = n
			qb.limit = uint64(n)
		}
	}
	if o := params["offset"]; o != "" {
		if n, err := strconv.Atoi(o); err == nil {
			qb.offset = uint64(n)
		}
	}
	if o := params["page"]; o != "" {
		if limit != 0 {
			if n, err := strconv.Atoi(o); err == nil {
				offset := limit*n - limit
				qb.offset = uint64(offset)
			}
		}
	}
}

// Build compiles the final SQL query with subqueries, sort, pagination
func (qb *QueryBuilder) Build() (string, []interface{}, error) {
    for alias, subquery := range qb.customColumn {
        if !validateAliasName(alias) {
            return "", nil, fmt.Errorf("unsafe alias name: %s", alias)
        }

        subquerySQL, subqueryArgs, err := subquery.ToSql()
        if err != nil {
            return "", nil, err
        }

        qb.builder = qb.builder.Column(fmt.Sprintf("(%s) AS %s", subquerySQL, alias), subqueryArgs...)
    }

    if qb.sortBy != "" {
        qb.builder = qb.builder.OrderBy(qb.sortBy + " " + qb.order)
    } else {
        qb.builder = qb.builder.OrderBy("id ASC")
    }

    sql, args, err := qb.builder.ToSql()
    if err != nil {
        return "", nil, err
    }

    // Ganti syntax paging agar valid untuk MySQL
    sql += " LIMIT " + strconv.FormatUint(qb.limit, 10) + " OFFSET " + strconv.FormatUint(qb.offset, 10)

    if qb.debug {
        fmt.Println("DEBUG SQL (Build):", sql)
        fmt.Println("DEBUG ARGS:", args)
    }

    return sql, args, nil
}

// Count builds a COUNT(*) query using the same WHERE clauses
func (qb *QueryBuilder) Count() (string, []interface{}, error) {
	countBuilder := sq.Select("COUNT(*)").From(qb.fromTable).PlaceholderFormat(sq.AtP)

	for _, cond := range qb.whereConditions {
		countBuilder = countBuilder.Where(cond)
	}

	sql, args, err := countBuilder.ToSql()
	if err != nil {
		return "", nil, err
	}

	if qb.debug {
		fmt.Println("DEBUG SQL (Count):", sql)
		fmt.Println("DEBUG ARGS:", args)
	}

	return sql, args, nil
}

// -- Helpers --

func validateColumnName(name string, allowed map[string]bool) bool {
	val := allowed[name]
	return val
}

func validateAliasName(alias string) bool {
	return safeAliasRegex.MatchString(alias)
}
