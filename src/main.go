package main

import (
	"compress/gzip"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//type application struct {
//	client     *mongo.Client
//	database   *mongo.Database
//	collection *mongo.Collection
//}

var (
	ApiURL string
)

func main() {
	ApiURL = os.Getenv("APIURL")
	if ApiURL == "" {
		log.Fatalln("Unable to get API URL from environment variables!")
	}
	homepage := NewHomepage()
	aboutpage := NewAboutPage()
	galaxiespage := NewGalaxiesPage()
	undertalePage := NewUndertalePage()
	musicPage := NewMusicPage()
	app.Route("/", homepage)
	app.Route("/about", aboutpage)
	app.Route("/galaxies", galaxiespage)
	app.Route("/undertale", undertalePage)
	app.Route("/blog", NewBlogPage())
	app.Route("/music", musicPage)

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
			"/web/static/images/gnu-head-sm.png",
			// Pages
			"/web/blocks/pages/about.html",
			"/web/blocks/pages/intro.html",
			"/web/blocks/snippets/bannerpanel.html",
			// Music
			"https://music-website.s3.nl-ams.scw.cloud/Tokusya-Seizon%20Wonder-la-der%21%21.mp3",
			"https://music-website.s3.nl-ams.scw.cloud/kegarenaki-barajuuji.mp3",
			"https://music-website.s3.nl-ams.scw.cloud/error-towa.mp3",
		},
	}

	app.GenerateStaticWebsite("./staticsite", handler)
	compressed := handlers.CompressHandlerLevel(handler, gzip.BestSpeed)
	http.Handle("/", compressed)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
