package routers

import (
	"github.com/ganamuhibudin/goapi/controllers"
	"github.com/ganamuhibudin/goapi/utils"
	"github.com/kataras/iris"
	"log"
)

func init() {
	log.Println("Initialize Notification Template Router")
	app := utils.GetIrisApplication()
	db := utils.GetDBConnection()

	// Email template collection
	app.PartyFunc("/users", func(users iris.Party) {
		userController := &controllers.UserController{DB: db}

		// Get all user
		users.Get("/", userController.GetAll)

		// Get user by id
		users.Get("/{id:int}", userController.GetUser)

		// Create user
		users.Post("/create", userController.CreateUser)

		// Update user by id
		users.Put("/{id:int}", userController.UpdateUser)

		// Delete user by id
		users.Delete("/{id:int}", userController.DeleteUser)
	})
}
