// Package helpers contains some useful functions for API actions
package helpers

import (
	"articlemaker/store"
	"net/http"
	"strings"
	"time"
)

var (
	// FieldsMapper helps to map json input to article update fields
	FieldsMapper = map[string]string{
		"id":           "Id",
		"title":        "Title",
		"body":         "Body",
		"publisher":    "Publisher",
		"category":     "Category",
		"published_at": "PublishedAt",
	}
)

// GetAllowedParams returns the query parameters allowed for filtering and updating articles
func GetAllowedParams() (map[string]struct{}, map[string]struct{}, map[string]struct{}) {
	searchColumns := make(map[string]struct{})
	searchColumns["created_at"] = struct{}{}
	searchColumns["published_at"] = struct{}{}

	updateColumns := make(map[string]struct{})
	updateColumns["title"] = struct{}{}
	updateColumns["body"] = struct{}{}
	updateColumns["published_at"] = struct{}{}

	relations := make(map[string]struct{})
	relations["category"] = struct{}{}
	relations["publisher"] = struct{}{}

	return searchColumns, relations, updateColumns
}

// GetRelationsIds checks if the given category and/or publisher exists
// returns their IDs if they exist and a bool false if one of them don't exist
func GetRelationsIds(r *http.Request) (categoryId int, publisherId int, notFound bool) {
	params := r.URL.Query()
	db := store.GetConnection()
	defer db.Close()

	if _, category := params["category"]; category {
		categoryRow := db.Table("categories").Where("name = ?", params["category"]).Select("id").Row()
		categoryRow.Scan(&categoryId)
		if categoryId == 0 {
			notFound = true
		}
	}

	if _, publisher := params["publisher"]; publisher {
		publisherRow := db.Table("publishers").Where("name = ?", params["publisher"]).Select("id").Row()
		publisherRow.Scan(&publisherId)
		if publisherId == 0 {
			notFound = true
		}
	}

	return categoryId, publisherId, notFound
}

// GetTimeFromString converts a string to a time.Time instance
func GetTimeFromString(dateString string) (date time.Time, malformed bool) {
	tokens := strings.Split(dateString, " ")

	if len(tokens) < 2 {
		return time.Time{}, true
	}

	dateTokens := strings.Split(tokens[0], "-")
	timeTokens := strings.Split(tokens[1], ":")

	if len(dateTokens) < 3 || len(timeTokens) < 3 {
		return time.Time{}, true
	}

	dateString = strings.Replace(dateString, " ", "T", 1)
	dateString += "+00:00" // todo get the timezone diff dynamically
	newTime, err := time.Parse(time.RFC3339, dateString)

	if err != nil {
		return time.Time{}, true
	}

	return newTime, false
}

// GetCategoryOrPublisherIdByName gets category or publisher id given the name
func GetCategoryOrPublisherIdByName(name string, table string) (id int, err error) {
	db := store.GetConnection()
	defer db.Close()

	row := db.Table(table).Where("name = ?", name).Select("id").Row()
	err = row.Scan(&id)

	return id, err
}
