package config

import (
	_ "github.com/ganamuhibudin/goapi/routers"
	"github.com/ganamuhibudin/goapi/utils"
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris"
	"os"
)

type Routes struct {
	DB *gorm.DB
}

// Setup sets routes
func (r *Routes) Setup(host string, port string) {

	// setup routes
	app := utils.GetIrisApplication()

	// Set loggin level
	app.Logger().SetLevel(os.Getenv("LOG_LEVEL"))

	// start server
	app.Run(iris.Addr(host+":"+port), iris.WithoutServerError(iris.ErrServerClosed))
}
