package helpers

import (
	"articlemaker/store"
	"net/http"
)

// GetAllowedSearchParams returns the query parameters allowed for filtering articles
func GetAllowedSearchParams() (map[string]struct{}, map[string]struct{}) {
	columns := make(map[string]struct{})
	columns["created_at"] = struct{}{}
	columns["published_at"] = struct{}{}

	relations := make(map[string]struct{})
	relations["category"] = struct{}{}
	relations["publisher"] = struct{}{}

	return columns, relations
}

// GetCategoryAndPublisherIds checks if the given category and/or publisher exists
// returns their IDs if they exist and a bool false if one of them don't exist
func GetCategoryAndPublisherIds(r *http.Request) (categoryId int, publisherId int, notFound bool) {
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

// title, body, category, publisher, and published_at

//
func GetAllowedUpdateFields()  {

}
