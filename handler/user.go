package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"screening/models"
	"screening/service"
)

type UserHandler interface {
	CreateOrUpdateUser(http.ResponseWriter, *http.Request)
}

type userHandler struct {
	userService service.UserService
}

func NewUserHandler(userService service.UserService) *userHandler {
	return &userHandler{
		userService: userService,
	}
}

func (h *userHandler) CreateOrUpdateUser(w http.ResponseWriter, r *http.Request) {
	funcdesc := "CreateOrUpdateUser"
	log.Println("enter handeler" + funcdesc)

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		log.Println("error while decoding")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	userId, err := h.userService.CreateOrUpdateUser(user)

	if err != nil {
		log.Println("error while creating user: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	message := "User created successfully"
	if user.ID != 0 {
		message = "User updated successfully"
	}

	log.Println("exit " + funcdesc)
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id":%d,"message":"`+message+`"}`, userId)))
}
