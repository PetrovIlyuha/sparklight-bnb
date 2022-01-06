package main

import (
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/config"
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/handlers"
	"github.com/PetrovIlyuha/sparklight-bnb/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

const PORT = ":3000"

var app config.AppConfig
var session *scs.SessionManager

func main() {

	//change to true when in production
	app.InProduction = false

	//adding session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	// end adding session


	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can't create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	log.Printf("Server is running on port: %s", PORT)
	srv := &http.Server{
		Addr: PORT,
		Handler: routes(&app),
	}
	err = srv.ListenAndServe()
	log.Fatal(err)
}
