package models

type Product struct {
	ID         int     `json:"id"`
	Nama       string  `json:"nama"`
	Harga      float64 `json:"harga"`
	KategoriID int     `json:"kategori_id"`
}
