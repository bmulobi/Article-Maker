package routes

import (
	"articlemaker/handlers"
	"github.com/gorilla/mux"
)

// prefix is the route prefix for this version of the API
const prefix = "/api/v1" // put in env

// HandleRequests sets up routes and handles all incoming http requests to the server
func HandleRequests() (router *mux.Router)  {
	router = mux.NewRouter().StrictSlash(true) // todo: add /api/v1 prefix

	router.HandleFunc(prefix + "/article", handlers.CreateArticle).Methods("POST")  // POST
	router.HandleFunc(prefix + "/article", handlers.GetAllArticles) // GET
	router.HandleFunc(prefix + "/article/{id:[0-9]+}", handlers.UpdateArticle).Methods("PUT")  // PUT
	router.HandleFunc(prefix + "/article/{id:[0-9]+}", handlers.DeleteArticle).Methods("DELETE") // DELETE
	router.HandleFunc(prefix + "/article/{id:[0-9]+}", handlers.GetArticle) // GET
	router.HandleFunc(prefix + "/", handlers.NotFound)

	return router
}
