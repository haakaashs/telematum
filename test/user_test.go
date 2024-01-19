package test

import (
	"screening/database"
	"screening/service"
	"testing"
)

var (
	userDB      = database.NewUserDB()
	userService = service.NewUserService(userDB)
)

func TestCreateUser(t *testing.T) {
	// we can have all the test cases in the tt test table here
}
