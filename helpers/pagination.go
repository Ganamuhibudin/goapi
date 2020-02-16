package helpers

import (
	"errors"
	"fmt"
	npq "github.com/Knetic/go-namedParameterQuery"
	"github.com/ganamuhibudin/goapi/models"
	"github.com/jinzhu/gorm"
	"math"
)

type Filter struct {
	Key   string
	Value string
}

func GetPagination(db *gorm.DB, model interface{}, conditions *npq.NamedParameterQuery, limit int, page int) (map[string]interface{}, error) {
	count := 0

	switch model.(type) {
	case models.Users:
		db.Model(&models.Users{}).Where(conditions.GetParsedQuery(), (conditions.GetParsedParameters())...).Count(&count)
	default:
		err := "Model type is not supported"
		return nil, errors.New(err)

	}

	// totalPage := math.Ceil(float64(count) / float64(pageConfig.Limit))
	totalPage := math.Ceil(float64(count) / float64(limit))
	hasNext := false

	if float64(page) < totalPage {
		hasNext = true
	}

	pagination := map[string]interface{}{
		// "page":         pageConfig.Page,
		"page":        page,
		"total_pages": totalPage,
		"total_items": count,
		// "per_page":     pageConfig.Limit,
		"per_page": limit,
		"has_next": hasNext,
		// "has_previous": pageConfig.Page > 1,
		"has_previous": page > 1,
	}

	fmt.Printf("Pagination: %+v\n", pagination)
	return pagination, nil
}
