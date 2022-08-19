package main

import (
	"github.com/pablocampo/webserver/pkg/config"
	"github.com/pablocampo/webserver/pkg/handlers"
	"github.com/pablocampo/webserver/pkg/render"
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	err = http.ListenAndServe(portNumber, nil)

	if err != nil {
		fmt.Println(fmt.Sprintf("cannot open port: %s error detail: (%s)",portNumber,  err))
	}
}
