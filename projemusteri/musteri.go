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
INSERT INTO musteri (age, email, first_name, last_name, money)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	var (
		age        int
		email      string
		first_name string
		last_name  string
		money      float32
		id1        int
		process    string
		update1    int
		key        string
	)
	for {
		fmt.Println("for new customer (1),for existing customer(2) key in:")
		fmt.Scan(&key)
		if key == "1" {
			fmt.Println("Enter your age:")
			fmt.Scan(&age)

			fmt.Println("Enter your email:")
			fmt.Scan(&email)

			fmt.Println("Enter your first_name:")
			fmt.Scan(&first_name)

			fmt.Println("Enter your last_name:")
			fmt.Scan(&last_name)

			fmt.Println("Enter your money:")
			fmt.Scan(&money)

			id := 0
			err = db.QueryRow(sqlStatement, age, email, first_name, last_name, money).Scan(&id)
			if err != nil {
				panic(err)
			}

			fmt.Println("New record ID is:", id)
		} else if key == "2" {
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
DELETE FROM musteri
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
UPDATE musteri
SET first_name =$2 , last_name = $3 , email = $4 , money = $5
WHERE id=$1
RETURNING id;`
	var (
		first_name1 string
		last_name1  string
		id2         int
		email       string
		id          int
		money       int
	)

	fmt.Println("enter the id you want to change:")
	fmt.Scan(&id2)

	fmt.Println("enter the new first name:")
	fmt.Scan(&first_name1)

	fmt.Println("enter the new last name:")
	fmt.Scan(&last_name1)

	fmt.Println("enter the new e-mail:")
	fmt.Scan(&email)

	fmt.Println("enter the new money:")
	fmt.Scan(&money)

	err = db.QueryRow(sqlStatement, id2, first_name1, last_name1, email, money).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, "updated successfully ..")
	return update1
}
