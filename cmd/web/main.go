package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jsur/go-web-helloworld/pkg/config"
	"github.com/jsur/go-web-helloworld/pkg/handlers"
	"github.com/jsur/go-web-helloworld/pkg/render"
)

const port = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	// pass pointer to app to NewRepo
	// to enable handlers pkg access to app config
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// pass pointer to app
	// this way render package gets access to app-wide AppConfig
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", port))

	serve := &http.Server{
		Addr:    port,
		Handler: Routes(&app),
	}

	err = serve.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
