package handlers

import (
	"UTS/database"
	"UTS/models"
	"encoding/json"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	rows, _ := database.DB.Query(`
		SELECT produk.id, produk.nama, produk.harga, kategori.nama 
		FROM produk 
		JOIN kategori ON produk.kategori_id = kategori.id`)
	type ProductDetail struct {
		ID           int     `json:"id"`
		Nama         string  `json:"nama"`
		Harga        float64 `json:"harga"`
		NamaKategori string  `json:"nama_kategori"`
	}
	var products []ProductDetail
	for rows.Next() {
		var p ProductDetail
		rows.Scan(&p.ID, &p.Nama, &p.Harga, &p.NamaKategori)
		products = append(products, p)
	}
	json.NewEncoder(w).Encode(products)
}

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var prod models.Product
	json.NewDecoder(r.Body).Decode(&prod)
	_, _ = database.DB.Exec("INSERT INTO produk(nama, harga, kategori_id) VALUES(?, ?, ?)",
		prod.Nama, prod.Harga, prod.KategoriID)
	w.WriteHeader(http.StatusCreated)
}
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var prod models.Product
	json.NewDecoder(r.Body).Decode(&prod)
	_, err := database.DB.Exec("UPDATE produk SET nama=?, harga=?, kategori_id=? WHERE id=?", prod.Nama, prod.Harga, prod.KategoriID, prod.ID)
	if err != nil {
		http.Error(w, "Gagal update", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	_, err := database.DB.Exec("DELETE FROM produk WHERE id=?", id)
	if err != nil {
		http.Error(w, "Gagal delete", 500)
		return
	}
	w.WriteHeader(http.StatusOK)
}
