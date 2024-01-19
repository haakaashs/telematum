package service

import (
	"log"
	"screening/database"
	"screening/models"
)

type UserService interface {
	CreateOrUpdateUser(models.User) (uint, error)
}

type userService struct {
	userDB database.UserDB
}

func NewUserService(userDB database.UserDB) *userService {
	return &userService{
		userDB: userDB,
	}
}

func (s *userService) CreateOrUpdateUser(user models.User) (uint, error) {
	funcdesc := "CreateOrUpdateUser"
	log.Println("enter service" + funcdesc)

	userId, err := s.userDB.CreateOrUpdateUser(user)
	if err != nil {
		return userId, err
	}

	log.Println("exit " + funcdesc)
	return userId, nil
}
