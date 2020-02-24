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
	"strconv"
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
		result = map[string]string{"message": "Article was not found"}
		w.WriteHeader(http.StatusNotFound)
	} else {
		result = article
	}

	json.NewEncoder(w).Encode(result)
}

// GetArticles gets all articles from the database
// can filter results by category publisher created_at published_at
func GetArticles(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	var articles []models.Article
	db := store.GetConnection()
	defer db.Close()

	db = db.Preload("Category").Preload("Publisher")
	if len(params) > 0 {
		categoryId, publisherId, notFound := helpers.GetRelationsIds(r)

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
		columns, relations, _ := helpers.GetAllowedParams()

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
// updates title, body, published_at, category, publisher fields
func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	var fieldsToUpdate map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&fieldsToUpdate)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Error processing the request")
		return
	}
	if _, ok := fieldsToUpdate["id"]; !ok {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode("Enter the article ID")
		return
	}
	db := store.GetConnection()
	defer db.Close()

	var article models.Article
	id, _ := fieldsToUpdate["id"].(float64)
	db.First(&article, strconv.Itoa(int(id)))

	if article.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(fmt.Sprintf("The article with ID %s was not found", strconv.Itoa(int(id))))
		return
	}
	_, relations, columns := helpers.GetAllowedParams()
	updates := make(map[string]interface{})

	for field, value := range fieldsToUpdate {
		if _, exists := columns[field]; exists {
			switch {
			case "Title" == helpers.FieldsMapper[field]:
				updates["title"] = value
			case "Body" == helpers.FieldsMapper[field]:
				updates["body"] = value
			case "PublishedAt" == helpers.FieldsMapper[field]:
				newDate, err := helpers.GetTimeFromString(value.(string))
				if err {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode("Enter parsing published_at")
					return
				}
				updates["published_at"] = newDate
			}
		}
		if _, exists := relations[field]; exists {
			switch {
			case "Category" == helpers.FieldsMapper[field]:
				categoryId, _ := helpers.GetCategoryOrPublisherIdByName(value.(string), "categories")
				if categoryId == 0 {
					var newCategory models.Category
					db.Create(&models.Category{Name: value.(string)}).Scan(&newCategory)
					updates["category_id"] = newCategory.Id
				} else {
					updates["category_id"] = categoryId
				}
			case "Publisher" == helpers.FieldsMapper[field]:
				publisherId, _ := helpers.GetCategoryOrPublisherIdByName(value.(string), "publishers")
				if publisherId == 0 {
					var newPublisher models.Publisher
					db.Create(&models.Publisher{Name: value.(string)}).Scan(&newPublisher)
					updates["publisher_id"] = newPublisher.Id
				} else {
					updates["publisher_id"] = publisherId
				}
			}
		}
	}
	db.Model(&article).Updates(updates)
	json.NewEncoder(w).Encode(article)
}

// DeleteArticle deletes an article from the database given the id
func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	tokens := strings.Split(r.URL.String(), "/")
	id := tokens[len(tokens)-1]
	db := store.GetConnection()
	defer db.Close()
	db.Delete(models.Article{}, fmt.Sprintf("id = %s", id))

	json.NewEncoder(w).Encode(fmt.Sprintf("Article with the ID %s was deleted", id))
}

// NotFound returns a 404 error message for any unknown path
func NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode("Page Not Found")
}
