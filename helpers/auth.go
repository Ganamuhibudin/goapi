package helpers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"strings"
)

func Authentication(ctx iris.Context) {
	authHeader := ctx.GetHeader("Authorization")
	if authHeader == "" {
		errMsg := "Unauthorized"
		NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}

	tokenString := strings.Split(authHeader, " ")
	_, err := jwt.Parse(tokenString[1], func(token *jwt.Token) (interface{}, error) {
		if jwt.GetSigningMethod("HS256") != token.Method {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("secret"), nil
	})

	if err != nil {
		errMsg := "Unauthorized. Error: " + err.Error()
		NewResponse(ctx, iris.StatusInternalServerError, errMsg)
		return
	}
	ctx.Next()
}
