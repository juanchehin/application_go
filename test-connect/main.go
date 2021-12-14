package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// Conectarse a la base de datos
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=postgres")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database %v\n", err))
	}

	defer conn.Close()

	log.Println("Connected to database!")

	// Testear la conexion
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!  %v\n", err)
		// log.Fatal(fmt.Sprintf(err))
	}

	log.Println("Pinged database!")
}
