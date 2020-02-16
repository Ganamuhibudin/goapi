package controllers

import (
	"encoding/json"
	npq "github.com/Knetic/go-namedParameterQuery"
	"github.com/ganamuhibudin/goapi/helpers"
	"github.com/ganamuhibudin/goapi/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
)

type UserController struct {
	DB *gorm.DB
}

func (uc *UserController) GetAll(ctx iris.Context) {
	var users []models.Users
	var bindVals map[string]interface{}

	limit := 10
	db := uc.DB

	// Get query conditions
	tempQuery := ""

	// Get Order By
	orderBy := ctx.URLParamDefault("order", "id ASC")

	// Get page number
	page := ctx.URLParamIntDefault("page", 1)

	// Get query bind values
	bind := ctx.URLParamDefault("bind", "")
	if bind != "" {
		json.Unmarshal([]byte(bind), &bindVals)
	}

	// Set offset
	offset := 0
	if page > 1 {
		offset = (page - 1) * limit
	}

	dbWithConfig := db.Debug().Limit(limit).Offset(offset).Order(orderBy)

	// Process query condition and its bind parameters
	conditions := npq.NewNamedParameterQuery(tempQuery)

	// Set query bind values in query condition
	conditions.SetValuesFromMap(bindVals)

	err := dbWithConfig.
		Where(conditions.GetParsedQuery(), ((conditions).GetParsedParameters())...).
		Find(&users).Error

	if err != nil {
		errMsg := "Query Error. Error: " + err.Error()

		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	// Get Pagination
	pagination, errPage := helpers.GetPagination(db, models.Users{}, conditions, limit, page)
	if errPage != nil {
		errMsg := "Pagination error. Error: " + errPage.Error()

		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	resp := map[string]interface{}{
		"items":      users,
		"pagination": pagination,
	}

	helpers.NewResponse(ctx, iris.StatusOK, resp)
}

func (uc *UserController) GetUser(ctx iris.Context) {
	users := &models.Users{}
	id := ctx.Params().Get("id")

	if uc.DB.Debug().First(users, id).RecordNotFound() {
		users = nil
	}

	helpers.NewResponse(ctx, iris.StatusOK, users)

	return
}
