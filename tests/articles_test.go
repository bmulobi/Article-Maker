// Package tests defines tests for the application
package tests

import (
	"articlemaker/handlers"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// happy path tests below todo: add unhappy path tests

// TestCreateArticle tests creating single article
func TestCreateArticle(t *testing.T) {
	SetUp()
	defer TearDown()

	rr := CreateArticle(t, map[string]string{})
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestGetArticle tests getting single article
func TestGetArticle(t *testing.T) {
	SetUp()
	defer TearDown()
	_ = CreateArticle(t, map[string]string{})

	response := GetArticle(t, 1)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	body := response.Body.String()
	if !strings.Contains(body, "\"id\":1,") {
		t.Errorf("Expected the response body to have the id=1 for the only article in the table")
	}
}

// TestGetArticles tests getting multiple articles
func TestGetArticles(t *testing.T) {
	SetUp()
	defer TearDown()

	_ = CreateArticle(t, map[string]string{})
	articleTwo := map[string]string{
		"categoryName": "test category two",
		"publisherName": "test publisher two",
		"articleTitle": "test article two",
		"articleBody": "test article two body",
	}
	_ = CreateArticle(t, articleTwo)

	request, err := http.NewRequest("GET", "/article", nil)
	if err != nil {
		t.Fatal(err)
	}
	response := httptest.NewRecorder()
	handler := http.HandlerFunc(handlers.GetArticles)
	handler.ServeHTTP(response, request)

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

// TestUpdateArticle tests updating an article
func TestUpdateArticle(t *testing.T) {
	SetUp()
	defer TearDown()
	_ = CreateArticle(t, map[string]string{})
	response := UpdateArticle(t)
	body := response.Body.String()

	if status := response.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	if !strings.Contains(body, `"title":"Updated title"`) {
		t.Errorf(`Expected the response body to have "title":"updated article title"`)
	}
}

// TestDeleteArticle tests deleting an article
func TestDeleteArticle(t *testing.T) {
	SetUp()
	defer TearDown()
	_ = CreateArticle(t, map[string]string{})
	_ = DeleteArticle(t)
	response := GetArticle(t,1)

	if status := response.Code; status != http.StatusNotFound {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusNotFound)
	}
}
