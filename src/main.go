package main

import (
	"log"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//type application struct {
//	client     *mongo.Client
//	database   *mongo.Database
//	collection *mongo.Collection
//}

const (
	apiurl = "https://quenten.nl:8007/"
)

func main() {
	homepage := NewHomepage()
	aboutpage := NewAboutPage()
	galaxiespage := NewGalaxiesPage()
	app.Route("/", homepage)
	app.Route("/about", aboutpage)
	app.Route("/galaxies", galaxiespage)

	app.Handle(getHTML, handleGetHTML)

	// This is executed on the client side only.
	// It handles client side stuff
	// It exits immediately on the server side
	app.RunWhenOnBrowser()

	icon := &app.Icon{
		Default: "/web/static/images/icon-small.png",
		Large:   "/web/static/images/icon.png",
	}
	handler := &app.Handler{
		Name:            "Internetica Galactica",
		Icon:            *icon,
		BackgroundColor: "#362730",
		ThemeColor:      "#362730",
		LoadingLabel:    "Internetica Galactica",
		Title:           "Internetica Galactica",
		Description:     "A 1990's style PWA!",
		Author:          "Quenten",
		Keywords: []string{
			"Based website",
			"Cool website",
			"PWA",
			"Programming",
			"Go", "Golang",
			"Webassembly", "WASM",
			"DutchEllie", "Quenten",
		},
		Styles: []string{
			"/web/static/style.css",
			"/web/static/adreena.css",
			"/web/static/anisha.css",
			"/web/static/havakana.css",
		},
		CacheableResources: []string{},
	}

	app.GenerateStaticWebsite("./staticsite", handler)
	http.Handle("/", handler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
