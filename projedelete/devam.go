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
DELETE FROM users
WHERE id = $1;`
	var delete int
	fmt.Println("enter your want delete id:")
	fmt.Scan(&delete)
	_, err = db.Exec(sqlStatement, delete)
	if err != nil {
		panic(err)
	}

	fmt.Println(delete, "Line deleted successfully")
}
