package routes

import (
	"screening/database"
	"screening/handler"
	"screening/service"
)

// database
var (
	userDB database.UserDB
)

// service
var (
	userService service.UserService
)

// handler
var (
	userHandler handler.UserHandler
)

// database
func buildDB() {
	userDB = database.NewUserDB()
}

// service
func buildService() {
	userService = service.NewUserService(userDB)
}

func buildHandler() {
	buildDB()
	buildService()
	userHandler = handler.NewUserHandler(userService)
}
