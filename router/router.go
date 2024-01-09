package router

import (
	"GoWebserver/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	return router
}
