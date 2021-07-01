package main

import (
	"fmt"
	"github.com/tbetcha/gotel/pkg/config"
	"github.com/tbetcha/gotel/pkg/handlers"
	"github.com/tbetcha/gotel/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Println(fmt.Sprintf("starting app on  port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
