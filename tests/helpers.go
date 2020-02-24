package tests

import (
	"articlemaker/handlers"
	"articlemaker/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// UpdateData defines data to update an article
type UpdateData struct {
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Publisher string `json:"publisher"`
	Category string `json:"category"`
	PublishedAt string `json:"published_at"`
}

// GetArticleForUpdate returns an article for update
func GetArticleForUpdate() UpdateData {
	data := UpdateData{
		Id:        1,
		Title:     "Updated title",
		Body:      "Updated body",
		Publisher: "Updated publisher",
		Category:  "Updated category",
		PublishedAt: "2020-02-24 21:59:31",
	}

	return data
}

// GetArticleForUpdate returns an article to create
func GetArticleForCreate(data map[string]string) models.Article  {
	categoryName := "test category one"
	publisherName := "test publisher one"
	articleTitle := "test article one"
	articleBody := "test article one body"

	if len(data) > 0  {
		categoryName = data["categoryName"]
		publisherName = data["publisherName"]
		articleTitle = data["articleTitle"]
		articleBody = data["articleBody"]
	}
	category := models.Category{
		Name: categoryName,
	}
	publisher := models.Publisher{
		Name: publisherName,
	}
	article := models.Article{
		Title:       articleTitle,
		Body:        articleBody,
		Publisher:   publisher,
		Category:    category,
	}

	return article
}

// CreateArticle helper for the repetitive action of creating an article
func CreateArticle(t *testing.T, data map[string]string) (response *httptest.ResponseRecorder) {
	body, _ := json.Marshal(GetArticleForCreate(data))
	request, err := http.NewRequest("POST", "/article", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.CreateArticle)
	handler.ServeHTTP(response, request)

	return response
}

// UpdateArticle helper for the updating an article
func UpdateArticle(t *testing.T) (response *httptest.ResponseRecorder) {
	body, _ := json.Marshal(GetArticleForUpdate())


	request, err := http.NewRequest("PUT", "/article", bytes.NewBuffer(body))
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.UpdateArticle)
	handler.ServeHTTP(response, request)

	return response
}

// GetArticle gets an  article given the id
func GetArticle(t *testing.T, id int) (response *httptest.ResponseRecorder) {
	request, err := http.NewRequest("GET", fmt.Sprintf("/article/%d", id), nil)
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetArticle)
	handler.ServeHTTP(response, request)

	return response
}

// DeleteArticle deletes an  article
func DeleteArticle(t *testing.T) (response *httptest.ResponseRecorder) {
	request, err := http.NewRequest("DELETE", "/article/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	response = httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.DeleteArticle)
	handler.ServeHTTP(response, request)

	return response
}
