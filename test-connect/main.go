package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	// ======================================
	// ==== Conectarse a la base de datos ===
	// ======================================
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=postgres")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect to database %v\n", err))
	}

	defer conn.Close()

	log.Println("Connected to database!")

	// ======================================
	// === Testear la conexion ====
	// ======================================
	err = conn.Ping()
	if err != nil {
		log.Fatal("Cannot ping database!  %v\n", err)
		// log.Fatal(fmt.Sprintf(err))
	}

	log.Println("Pinged database!")

	// ======================================
	// === Acceder a las filas de las tablas ===
	// ======================================
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// ======================================
	// === Insertar una fila (row) ===
	// ======================================
	query := `insert into users (first_name,last_name) values ($1,$2)`
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Insert a row!")

	// ======================================
	// === Acceder a las filas de las tablas ===
	// ======================================
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// ======================================
	// === Actualizo una fila  ===
	// ======================================
	stmt := `update users set first_name = $1 where id=$2`
	_, err = conn.Exec(stmt, "Jackie", 2)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated a row!")

	// ======================================
	// === Acceder a las filas de las tablas ===
	// ======================================
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// ======================================
	// === Obtengo una fila por id ===
	// ======================================
	query = `select id,first_name,last_name from users where id = $1`

	var firstName, lastName string
	var id int

	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Query rows return : ", id, firstName, lastName)

	// ======================================
	// === Actualizo una fila  ===
	// ======================================
	query = `delete from users where id=$1`

	_, err = conn.Exec(query, 5)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Delete a row!")

	// ======================================
	// === Acceder a las filas de las tablas ===
	// ======================================
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id,first_name,last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is : ", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("------------------")

	return nil
}
