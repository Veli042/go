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
UPDATE users
SET first_name =$2 , last_name = $3
WHERE id=$1
RETURNING id,email;`
	var (
		first_name1 string
		last_name1  string
		id1         int
		email       string
		id          int
	)

	fmt.Println("enter your first name:")
	fmt.Scan(&first_name1)

	fmt.Println("enter your last name:")
	fmt.Scan(&last_name1)

	fmt.Println("enter your id:")
	fmt.Scan(&id1)

	err = db.QueryRow(sqlStatement, id1, first_name1, last_name1).Scan(&id, &email)
	if err != nil {
		panic(err)
	}
	fmt.Println(id, email)

}
