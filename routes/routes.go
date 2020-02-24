// Package routes defines the routing for the API
package routes

import (
	"articlemaker/handlers"
	"github.com/gorilla/mux"
)

// prefix is the route prefix for this version of the API
const prefix = "/api/v1"

// HandleRequests sets up routes and handles all incoming http requests to the server
func HandleRequests() (router *mux.Router) {
	router = mux.NewRouter().StrictSlash(true)

	router.HandleFunc(prefix+"/article", handlers.CreateArticle).Methods("POST")
	router.HandleFunc(prefix+"/article", handlers.UpdateArticle).Methods("PUT")
	router.HandleFunc(prefix+"/article", handlers.GetArticles)
	router.HandleFunc(prefix+"/article/{id:[0-9]+}", handlers.DeleteArticle).Methods("DELETE")
	router.HandleFunc(prefix+"/article/{id:[0-9]+}", handlers.GetArticle)
	router.HandleFunc(prefix+"/", handlers.NotFound)

	return router
}
