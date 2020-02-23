package controllers

import (
	"encoding/json"
	npq "github.com/Knetic/go-namedParameterQuery"
	"github.com/ganamuhibudin/goapi/helpers"
	"github.com/ganamuhibudin/goapi/models"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"golang.org/x/crypto/bcrypt"
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

func (uc *UserController) CreateUser(ctx iris.Context) {
	user := &models.Users{}

	errInput := ctx.ReadJSON(&user)
	if errInput != nil {
		errMsg := "Failed to read user input. Error: " + errInput.Error()
		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	// Check if email already exists
	usr := uc.DB.Debug().Where("email = ?", user.Email).First(&user)
	if !usr.RecordNotFound() {
		respMsg := "Email " + user.Email + " has already been used."
		helpers.NewResponse(ctx, iris.StatusBadRequest, respMsg)
		return

	}

	user.Password, _ = HashPassword(user.Password)

	errCreate := uc.DB.Debug().Create(&user).Error
	if errCreate != nil {
		errMsg := "Failed to create user. Error: " + errCreate.Error()
		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	helpers.NewResponse(ctx, iris.StatusCreated, user)
	return
}

func (uc *UserController) UpdateUser(ctx iris.Context) {
	var userData map[string]interface{}
	user := &models.Users{}
	userID := ctx.Params().Get("id")

	err := ctx.ReadJSON(&userData)
	if err != nil {
		errMsg := "Failed to read data. Error: " + err.Error()
		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	// Validate data exist
	data := uc.DB.Debug().First(&user, userID)
	if data.RecordNotFound() {
		helpers.NewResponse(ctx, iris.StatusNotFound, "User doesn't exist")
		return
	}

	// Update data
	uc.DB.Model(&user).Updates(userData)

	helpers.NewResponse(ctx, iris.StatusOK, user)
	return
}

func (uc *UserController) DeleteUser(ctx iris.Context) {
	var user models.Users
	userID := ctx.Params().Get("id")

	// Get User
	usr := uc.DB.Debug().First(&user, userID)

	if usr.RecordNotFound() {
		helpers.NewResponse(ctx, iris.StatusNotFound, "User doesn't exist")
		return
	}

	// Delete user
	errDelUser := uc.DB.Debug().Delete(&user).Error
	if errDelUser != nil {
		errMsg := "Failed to delete user. Error: " + errDelUser.Error()
		helpers.NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	helpers.NewResponse(ctx, iris.StatusOK, "Delete User Successful")
	return
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
