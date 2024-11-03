package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"github.com/sangketkit01/go-web/pkg/config"
	"github.com/sangketkit01/go-web/pkg/handlers"
	"github.com/sangketkit01/go-web/pkg/render"
	"log"
	"net/http"
	"time"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalln("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))

	server := &http.Server{
		Addr:    portNumber,
		Handler: route(&app),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatalln(err)
	}
}
