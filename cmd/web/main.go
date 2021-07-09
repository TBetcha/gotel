package main

import (
	"encoding/gob"
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/tbetcha/gotel/internal/config"
	"github.com/tbetcha/gotel/internal/driver"
	"github.com/tbetcha/gotel/internal/handlers"
	"github.com/tbetcha/gotel/internal/helpers"
	"github.com/tbetcha/gotel/internal/models"
	"github.com/tbetcha/gotel/internal/render"
	"log"
	"net/http"
	"os"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	db,err := run()
	if err != nil{
		log.Fatal(err)
	}
	defer db.SQL.Close()

	fmt.Println(fmt.Sprintf("starting app on  port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() (*driver.DB, error) {
	gob.Register(models.Reservation{})
	gob.Register(models.User{})
	gob.Register(models.Room{})
	gob.Register(models.Restriction{})
	// change this to true when in prod
	app.InProduction = false

	infoLog = log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime)
	app.InfoLog = infoLog

	errorLog = log.New(os.Stdout, "ERROR:\t",log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	// connect to Db
	log.Println("connecting to db")
	db, err := driver.ConnectSQL("host=localhost port=5432 dbname=hotels user=tmb password=")
	if err != nil{
		log.Fatal("cannot connect to db, oh nooooo")
	}
	log.Println("connected to db bro")
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
		return nil,err
	}

	app.TemplateCache = tc

	app.UseCache = false
	repo := handlers.NewRepo(&app, db)
	handlers.NewHandlers(repo)
	render.NewRenderer(&app)
	helpers.NewHelpers(&app)

	return db, nil
}