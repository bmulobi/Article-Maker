// Package handlers implements functions for processing API requests
package handlers

import (
	"articlemaker/models"
	"articlemaker/store"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

// CreateArticle creates a new article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	var article models.Article
	json.Unmarshal(body, &article)

	db := store.GetConnection()
	defer db.Close()

	//json.Marshal(&article) todo : for validation only

	result := db.Create(&article)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetArticle gets an article from the database given the id
func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := store.GetConnection()
	defer db.Close()

	var article models.Article
	db.First(&article, vars["id"])

	var result interface{}
	if article.Id == 0 {
		result =  map[string]string{"message": "Article was not found"}
		w.WriteHeader(http.StatusNotFound)
	} else {
		result = article
	}

	json.NewEncoder(w).Encode(result)
}

// GetAllArticles gets all articles from the database
func GetAllArticles(w http.ResponseWriter, r *http.Request) {
	var articles []models.Article
	db := store.GetConnection()
	defer db.Close()

	db.Find(&articles)
	json.NewEncoder(w).Encode(articles)
}

// UpdateArticle updates an article given the id
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Internal server error")
		return
	}

	// get only needed fields
	// save to db (update)

	// set status code 200
	var article models.Article
	json.Unmarshal(body, &article)

	// set status code
	json.NewEncoder(w).Encode(article)
}

// DeleteArticle deletes an article from the database
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	db := store.GetConnection()
	defer db.Close()

	db.Delete(models.Article{}, fmt.Sprintf("id = %s", id))

	json.NewEncoder(w).Encode(fmt.Sprintf("Article with the ID %s was deleted", id))
}

// NotFound returns a 404 error message for any unknown path
func NotFound(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Page Not Found")
}
