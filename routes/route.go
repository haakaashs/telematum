package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SetupJsonApi() {
	buildHandler()

	server := mux.NewRouter()

	server.HandleFunc("/updateUser", userHandler.CreateOrUpdateUser).Methods("PUT")

	log.Println("started")
	if err := http.ListenAndServe(":8080", server); err != nil {
		log.Println("error while starting server: ", err)
	}
	// http.HandleFunc("/updateUser")
}
