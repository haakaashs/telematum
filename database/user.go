package database

import (
	"log"
	"screening/config"
	"screening/models"

	"gorm.io/gorm"
)

type UserDB interface {
	CreateOrUpdateUser(models.User) (uint, error)
}

type userDB struct {
	db *gorm.DB
}

func NewUserDB() *userDB {
	return &userDB{
		db: config.Conn,
	}
}

func (d *userDB) CreateOrUpdateUser(user models.User) (uint, error) {
	funcdesc := "CreateOrUpdateUser"
	log.Println("enter DB" + funcdesc)

	result := d.db.Debug().Save(&user)
	if err := result.Error; err != nil {
		log.Println("error in DB query: ", err.Error())
		return user.ID, err
	}

	log.Println("exit " + funcdesc)
	return user.ID, nil
}
