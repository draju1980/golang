package router

import (
	"github.com/draju1980/golang/services"
	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Route {
	router := mux.NewRouter()
	router.HandleFunc("/posts", services.GetAllPosts).Methods("GET")

	return router
}
