package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "12345"
	dbname   = "dburun"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO urun (ad,marka,stok,kategori)
VALUES($1, $2, $3, $4)	
RETURNING id`

	var (
		ad       string
		marka    string
		stok     int
		kategori string
	)

	fmt.Println("Ürünün adını giriniz:")
	fmt.Scan(&ad)

	fmt.Println("Ürünün markasını giriniz:")
	fmt.Scan(&marka)

	fmt.Println("Ürünün stoğunu giriniz:")
	fmt.Scan(&stok)

	fmt.Println("Ürünün kategorisini giriniz:")
	fmt.Scan(&kategori)

	id := 0
	err = db.QueryRow(sqlStatement, ad, marka, stok, kategori).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("Yeni ID:", id)
}
