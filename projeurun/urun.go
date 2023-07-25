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
INSERT INTO urun (name,product,stock,category)
VALUES($1, $2, $3, $4)	
RETURNING id`

	var (
		name     string
		product  string
		stock    int
		category string
		process  string
		id1      int
		update1  int
	)
	for {

		fmt.Println("for new product (1), for existing product (2) key in:")
		fmt.Scan(&process)

		if process == "1" {
			fmt.Println("enter the product name:")
			fmt.Scan(&name)

			fmt.Println("enter the product brand:")
			fmt.Scan(&product)

			fmt.Println("enter the product stock:")
			fmt.Scan(&stock)

			fmt.Println("enter the product category:")
			fmt.Scan(&category)

			id := 0
			err = db.QueryRow(sqlStatement, name, product, stock, category).Scan(&id)
			if err != nil {
				panic(err)
			}
			fmt.Println("new record id is:", id)
		} else if process == "2" {
			break
		}
	}

	for {
		fmt.Println("for deletion (-),for updation (+),for terminate (0) key in:")
		fmt.Scan(&process)

		if process == "-" {
			delete1 := delete(id1)
			fmt.Println(delete1)
		} else if process == "+" {
			update2 := update(update1)
			fmt.Println(update2)
		} else if process == "0" {
			break
		}
	}
}

func delete(id1 int) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `
DELETE FROM urun
WHERE id = $1;`
	var delete int
	fmt.Println("enter your want delete id:")
	fmt.Scan(&delete)
	_, err = db.Exec(sqlStatement, delete)
	if err != nil {
		panic(err)
	}

	fmt.Println(delete, "Line deleted successfully")

	return id1
}

func update(update1 int) int {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
UPDATE urun
SET name =$2 , product = $3 , stock = $4 , category = $5
WHERE id=$1
RETURNING id;`
	var (
		name     string
		product  string
		id2      int
		stock    string
		id       int
		category int
	)

	fmt.Println("enter the id you want to change:")
	fmt.Scan(&id2)

	fmt.Println("enter the new name:")
	fmt.Scan(&name)

	fmt.Println("enter the new product name:")
	fmt.Scan(&product)

	fmt.Println("enter the new stock:")
	fmt.Scan(&stock)

	fmt.Println("enter the new category:")
	fmt.Scan(&category)

	err = db.QueryRow(sqlStatement, id2, name, product, stock, category).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, "updated successfully ..")
	return update1
}
