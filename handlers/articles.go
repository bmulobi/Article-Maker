// Package handlers implements functions for processing API requests
package handlers

import (
	"articlemaker/handlers/helpers"
	"articlemaker/models"
	"articlemaker/store"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
)

// CreateArticle creates a new article
func CreateArticle(w http.ResponseWriter, r *http.Request) {
	// todo : add request validation
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

	result := db.Set("gorm:association_autoupdate", false).Create(&article)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)
}

// GetArticle gets an article from the database given the id
func GetArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	db := store.GetConnection()
	defer db.Close()

	var article models.Article
	db.Preload("Category").Preload("Publisher").First(&article, vars["id"])

	var result interface{}
	if article.Id == 0 {
		result =  map[string]string{"message": "Article was not found"}
		w.WriteHeader(http.StatusNotFound)
	} else {
		result = article
	}

	json.NewEncoder(w).Encode(result)
}

// GetArticles gets all articles from the database
func GetArticles(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var articles []models.Article
	db := store.GetConnection()
	defer db.Close()

	db = db.Preload("Category").Preload("Publisher")
	if len(params) > 0 {
		categoryId, publisherId, notFound := helpers.GetCategoryAndPublisherIds(r)

		if notFound {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("No results for the given parameters")
			return
		}
		if categoryId != 0 {
			db = db.Where("category_id = ?", categoryId)
		}
		if publisherId != 0 {
			db = db.Where("publisher_id = ?", publisherId)
		}
		columns, relations := helpers.GetAllowedSearchParams()

		isColumn := false
		isRelation := false

		for param, value := range params {
			if _, isColumn = columns[param]; isColumn {
				db = db.Where(fmt.Sprintf("%s = ?", param), fmt.Sprintf("%s", strings.Join(value, " ")))
			} else if _, isRelation = relations[param]; isRelation {
				continue
			} else {
				w.WriteHeader(http.StatusNotFound)
				response := fmt.Sprintf("%s is not a valid parameter", param)
				response += fmt.Sprintf(" use one of : created_at published_at category publisher")
				json.NewEncoder(w).Encode(response)
				return
			}
		}
		if !strings.Contains(fmt.Sprintf("%s", db.QueryExpr()), "?") {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("No results for the given parameters")
			return
		}
	}

	db.Find(&articles)
	if len(articles) == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("No results for the given parameters")
	} else {
		json.NewEncoder(w).Encode(articles)
	}
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

// DeleteArticle deletes an article from the database given the id
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
