package routers

import (
	"github.com/ganamuhibudin/goapi/utils"
	"github.com/kataras/iris"
	"log"
)

func init() {
	log.Println("Initialize default router...")

	app := utils.GetIrisApplication()

	// Default route
	app.Get("/", func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusOK)
		ctx.JSON(iris.Map{
			"message": "Winter is coming !",
		})
	})

	// Error route 404
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context) {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{
			"message": "You just got lost in shiva. Nothing to see here...",
		})
	})
}
