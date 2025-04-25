package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() {
	var err error
	DB, err = sql.Open("mysql", "root:@tcp(localhost)/uts_crud")
	if err != nil {
		log.Fatal("Gagal Koneksi ke Database:", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatal("Gagal Konek ke Database:", err)
	}
	fmt.Println("Berhasil Konek ke Database")
}
