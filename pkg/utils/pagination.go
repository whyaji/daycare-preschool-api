package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/whyaji/daycare-preschool-api/pkg/types"
	"gorm.io/gorm"
)

func GetPageAndLimitFromQuery(c *fiber.Ctx) (page, limit int) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		page = 1
	}

	limit, err = strconv.Atoi(c.Query("limit"))
	if err != nil {
		limit = 10
	}

	return page, limit
}

func toSnakeCase(str string) string {
	var snakeCase string
	for i, char := range str {
		if i > 0 && char >= 'A' && char <= 'Z' {
			snakeCase += "_" + strings.ToLower(string(char))
		} else {
			snakeCase += strings.ToLower(string(char))
		}
	}
	return snakeCase
}

// c.Query("filter") => "name:john,doe:in;age:20:gt;birthplace:usa"
// first split by ":" is key, second split is value, third split is operator
// value can be split by "," to get multiple value
// operator can be "in", "gt", "lt", "gte", "lte", "notin", "like", "null", "notnull", default is "in"
func GetFilterConditionFromQuery(c *fiber.Ctx) map[string]any {
	filter := c.Query("filter")
	if filter == "" {
		return nil
	}

	conditions := make(map[string]any)
	filterParts := strings.SplitSeq(filter, ";")
	for filterPart := range filterParts {
		filterPartParts := strings.Split(filterPart, ":")
		if len(filterPartParts) > 3 || len(filterPartParts) < 2 {
			continue
		} else if len(filterPartParts) == 2 {
			filterPartParts = append(filterPartParts, "in")
		}

		key := toSnakeCase(filterPartParts[0])
		valueParts := strings.Split(filterPartParts[1], ",")
		operator := filterPartParts[2]

		conditions[key] = map[string]any{operator: valueParts}
	}

	return conditions
}

func GetPaginationFilterFromQuery(c *fiber.Ctx) types.PaginationFilter {
	page, limit := GetPageAndLimitFromQuery(c)
	filters := GetFilterConditionFromQuery(c)
	return types.PaginationFilter{
		Page:    page,
		Limit:   limit,
		Filters: filters,
	}
}

// ApplyFilters dynamically applies filters to a GORM query
func ApplyFilters(query *gorm.DB, filters map[string]any) *gorm.DB {
	for key, condition := range filters {
		for operator, values := range condition.(map[string]any) {
			switch operator {
			case "in":
				query = query.Where(fmt.Sprintf("%s IN (?)", key), values)
			case "notin":
				query = query.Where(fmt.Sprintf("%s NOT IN (?)", key), values)
			case "gt":
				query = query.Where(fmt.Sprintf("%s > ?", key), values.([]string)[0])
			case "lt":
				query = query.Where(fmt.Sprintf("%s < ?", key), values.([]string)[0])
			case "gte":
				query = query.Where(fmt.Sprintf("%s >= ?", key), values.([]string)[0])
			case "lte":
				query = query.Where(fmt.Sprintf("%s <= ?", key), values.([]string)[0])
			case "like":
				query = query.Where(fmt.Sprintf("%s LIKE ?", key), "%"+values.([]string)[0]+"%")
			case "null":
				query = query.Where(fmt.Sprintf("%s IS NULL", key))
			case "notnull":
				query = query.Where(fmt.Sprintf("%s IS NOT NULL", key))
			default:
				query = query.Where(fmt.Sprintf("%s IN (?)", key), values)
			}
		}
	}
	return query
}

func ApplyYearMonthFilter(query *gorm.DB, filters map[string]any, filterBy string) *gorm.DB {
	yearFilter, yearExists := filters["year"]
	monthFilter, monthExists := filters["month"]

	if yearExists {
		yearConditions := yearFilter.(map[string]any)
		for operator, values := range yearConditions {
			switch operator {
			case "in":
				query = query.Where(fmt.Sprintf("YEAR(%s) IN (?)", filterBy), values)
			case "notin":
				query = query.Where(fmt.Sprintf("YEAR(%s) NOT IN (?)", filterBy), values)
			case "gt":
				query = query.Where(fmt.Sprintf("YEAR(%s) > ?", filterBy), values.([]string)[0])
			case "lt":
				query = query.Where(fmt.Sprintf("YEAR(%s) < ?", filterBy), values.([]string)[0])
			case "gte":
				query = query.Where(fmt.Sprintf("YEAR(%s) >= ?", filterBy), values.([]string)[0])
			case "lte":
				query = query.Where(fmt.Sprintf("YEAR(%s) <= ?", filterBy), values.([]string)[0])
			default:
				query = query.Where(fmt.Sprintf("YEAR(%s) IN (?)", filterBy), values)
			}
		}
	}

	if monthExists {
		monthConditions := monthFilter.(map[string]any)
		for operator, values := range monthConditions {
			switch operator {
			case "in":
				query = query.Where(fmt.Sprintf("MONTH(%s) IN (?)", filterBy), values)
			case "notin":
				query = query.Where(fmt.Sprintf("MONTH(%s) NOT IN (?)", filterBy), values)
			case "gt":
				query = query.Where(fmt.Sprintf("MONTH(%s) > ?", filterBy), values.([]string)[0])
			case "lt":
				query = query.Where(fmt.Sprintf("MONTH(%s) < ?", filterBy), values.([]string)[0])
			case "gte":
				query = query.Where(fmt.Sprintf("MONTH(%s) >= ?", filterBy), values.([]string)[0])
			case "lte":
				query = query.Where(fmt.Sprintf("MONTH(%s) <= ?", filterBy), values.([]string)[0])
			default:
				query = query.Where(fmt.Sprintf("MONTH(%s) IN (?)", filterBy), values)
			}
		}
	}

	return query
}
