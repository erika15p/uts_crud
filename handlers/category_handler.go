package handlers

import (
	"UTS/database"
	"UTS/models"
	"encoding/json"
	"net/http"
)

func GetCategories(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query("SELECT id, nama FROM kategori")
	var categories []models.Category
	for rows.Next() {
		var cat models.Category
		rows.Scan(&cat.ID, &cat.Nama)
		categories = append(categories, cat)
	}
	json.NewEncoder(w).Encode(categories)
}

func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var cat models.Category
	json.NewDecoder(r.Body).Decode(&cat)
	_, _ = database.DB.Exec("INSERT INTO kategori(nama) VALUES(?)", cat.Nama)
	w.WriteHeader(http.StatusCreated)
}
