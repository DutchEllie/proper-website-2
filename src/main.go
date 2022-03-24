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

func main() {
	homepage := NewHomepage()
	aboutpage := NewAboutPage()
	galaxiespage := NewGalaxiesPage()
	undertalePage := NewUndertalePage()
	emptyPage := NewEmptyPage()
	app.Route("/", homepage)
	app.Route("/about", aboutpage)
	app.Route("/galaxies", galaxiespage)
	app.Route("/undertale", undertalePage)
	app.Route("/empty", emptyPage)

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
			"/web/static/form.css",
		},
		CacheableResources: []string{
			// Images
			"/web/static/images/email3.gif",
			"/web/static/images/rin-len1.webp",
			"/web/static/images/background_star.gif",
			"/web/static/images/kanata-1.gif",
			"/web/static/images/rin-1.gif",
			"/web/static/images/rin-2.gif",
			// Pages
			"/web/blocks/pages/about.html",
			"/web/blocks/pages/galaxies.html",
			"/web/blocks/pages/intro.html",
			"/web/blocks/snippets/bannerpanel.html",
		},
	}

	app.GenerateStaticWebsite("./staticsite", handler)
	http.Handle("/", handler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}