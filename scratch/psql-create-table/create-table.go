package main

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/lib/pq"
)

func main() {

	db, err := sql.Open("postgres", "user=ranji dbname=pqgotest sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	stmt, err := db.Exec("CREATE TABLE testtable (id int PRIMARY KEY, name varchar)")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(stmt)
	fmt.Println("Created a table ", "testtable")

}
