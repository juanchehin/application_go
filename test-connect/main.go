package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	_, err := sql.Open("pgx", "host=localhost post=5432 dbname=test_connect user=postgres password=postgres")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database %v\n", err))
	}
}
