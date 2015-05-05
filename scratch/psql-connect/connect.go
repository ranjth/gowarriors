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


	age := 32
	rows, err := db.Query("SELECT name from users WHERE age = $1", age)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string

	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Query Done")

	insertRecord(db, "atthai", 49)
}

func insertRecord (db *sql.DB, name string, age int) {

	stmt, err := db.Prepare("INSERT INTO users(name, age) VALUES($1, $2)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(name, age)
	if err != nil {
		log.Fatal(err)
	}

	_, err = res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("ID = %d, affected = %d", 999 , rowCnt)

	fmt.Println("Insert Done")

}
