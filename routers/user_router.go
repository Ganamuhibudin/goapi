package routers

import (
	"github.com/ganamuhibudin/goapi/controllers"
	"github.com/ganamuhibudin/goapi/helpers"
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

		authentication := helpers.Authentication

		// User login
		users.Post("/auth", userController.Login)

		// Get all user
		users.Get("/", authentication, userController.GetAll)

		// Get user by id
		users.Get("/{id:int}", authentication, userController.GetUser)

		// Create user
		users.Post("/create", authentication, userController.CreateUser)

		// Update user by id
		users.Put("/{id:int}", authentication, userController.UpdateUser)

		// Delete user by id
		users.Delete("/{id:int}", authentication, userController.DeleteUser)
	})
}
