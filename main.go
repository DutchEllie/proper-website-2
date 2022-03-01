package main

import (
	"log"
	"net/http"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func main() {
	homepage := components.NewHomepage()
	app.Route("/", homepage)

	// This is executed on the client side only.
	// It handles client side stuff
	// It exits immediately on the server side
	app.RunWhenOnBrowser()

	http.Handle("/", &app.Handler{
		Name:        "Internetica Galactica",
		Description: "A 1990's style PWA!",
		Styles: []string{
			"/web/static/style.css",
			"/web/static/anisha.css",
		},
		CacheableResources: []string{
			"/web/static/style.css",
			"/web/static/anisha.css",
			"/web/static/images/background_star.gif",
		},
	})

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
