package main

import (
	// "encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/tsawler/bookings/internal/config"
	"github.com/tsawler/bookings/internal/handlers"
	"github.com/tsawler/bookings/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {
	db, err := run()
	if err != nil {
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	// _ = http.ListenAndServe(portNumber, nil)
	fmt.Println("Pasa 2")
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	fmt.Println("Pasa 3")
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run(*driver.DB, error) {
	fmt.Println("Hello World!")

	// gob.Register(models.Reservation{})
	// cambiar esto a true en produccion
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// Conexion a la BD
	log.Println("Conectando a la BD...")
	db, err = driver.ConnectSQL("host=localhost port=5432 dbname=bookings user=postgres password=postgres")
	if err != nil {
		log.Fatal("Error en conexion a la BD")
	}

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("no se pudo crear template cache en main.go", err)
		fmt.Println("Pasa 0")
		return nil, err
	}
	fmt.Println("Pasa 1")

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// render.NewTemplate(&app)
	render.NewRenderer(&app)
	// helpers.NewHelpers(&app)

	return db, nil
}
