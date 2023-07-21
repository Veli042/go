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
INSERT INTO users (age, email, first_name, last_name,money)
VALUES ($1, $2, $3, $4, $5)
RETURNING id`
	var (
		age        int
		email      string
		first_name string
		last_name  string
		money      float32
	)
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

}
