package router

import (
	"github.com/grahamuk2018/GoWebserver/services"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")
	router.HandleFunc("/posts", services.CreatePost).Methods("POST")
	router.HandleFunc("/posts/{id}", services.CreatePost).Methods("GET")
	router.HandleFunc("/posts/{id}", services.UpdatePost).Methods("PUT")
	router.HandleFunc("/posts/{id}", services.DeletePost).Methods("DELETE")
	return router
}
