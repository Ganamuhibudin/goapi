package helpers

import (
	"encoding/json"

	"github.com/kataras/iris"
)

type Detail struct {
	Error interface{} `json:"errors,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

// Response struct for message json
type Messages struct {
	Status  int  `json:"status"`
	Success bool `json:"success"`
	*Detail
}

// Validation HTTP Status Code
// Retrun true or false
func IsSuccessCode(c int) bool {
	return c > 199 && c < 300

}

// Check JSON type
func IsJSON(s string) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil

}

// Simple json response
// Handler from kataras/ iris
// code of HTTP status code
// Message Body
func NewResponse(ctx iris.Context, code int, content interface{}) {
	// check status code success
	status := IsSuccessCode(code)

	// Set value to struct
	var res Messages
	res.Status = code
	res.Success = status

	// Init body response
	body := make(map[string]interface{})
	var returns interface{}
	returns = nil

	if _, ok := content.(string); ok {
		// check if content as JSON
		if IsJSON(content.(string)) == false {
			body["message"] = content
		} else {
			json.Unmarshal([]byte(content.(string)), &body)
		}
	} else {
		returns = content
	}

	// TO DO
	// WILL REFACTOR THIS LINE CODE
	if status == true {
		if returns != nil {
			res.Detail = &Detail{Data: returns}
		} else {
			res.Detail = &Detail{Data: body}
		}
	} else {
		if returns != nil {
			res.Detail = &Detail{Error: returns}
		} else {
			res.Detail = &Detail{Error: body}
		}
	}

	ctx.StatusCode(code)
	ctx.JSON(res)
}