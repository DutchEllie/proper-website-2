package main

import (
	"compress/gzip"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"

	"dutchellie.nl/DutchEllie/proper-website-2/api"
	"github.com/gorilla/handlers"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//type application struct {
//	client     *mongo.Client
//	database   *mongo.Database
//	collection *mongo.Collection
//}

var jar http.CookieJar
var client http.Client

func main() {
	// Create cookiejar
	var err error
	jar, err = cookiejar.New(nil)
	if err != nil {
		log.Fatalf("Error creating cookiejar: %s\n", err.Error())
	}

	client = http.Client{
		Jar: jar,
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
		Scripts: []string{
			"https://cdn.tailwindcss.com",
			"/web/static/tailwind.config.js",
		},
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
			"https://music-website.s3.nl-ams.scw.cloud/diamond-city-lights-lazulight.opus",
			"https://music-website.s3.nl-ams.scw.cloud/tsunami-finana.opus",
			"https://music-website.s3.nl-ams.scw.cloud/%E7%A5%9E%E3%81%A3%E3%81%BD%E3%81%84%E3%81%AA.m4a",
			"https://music-website.s3.nl-ams.scw.cloud/Servant%20of%20Evil%20with%20English%20Sub%20-%20%E6%82%AA%E3%83%8E%E5%8F%AC%E4%BD%BF%20-%20Kagamine%20Len%20-%20HQ.m4a",
			"https://music-website.s3.nl-ams.scw.cloud/%E3%83%94%E3%83%8E%E3%82%AD%E3%82%AA%E3%83%94%E3%83%BC%20-%20%E3%81%8D%E3%81%BF%E3%82%82%E6%82%AA%E3%81%84%E4%BA%BA%E3%81%A7%E3%82%88%E3%81%8B%E3%81%A3%E3%81%9F%20feat.%20%E5%88%9D%E9%9F%B3%E3%83%9F%E3%82%AF%20_%20I%27m%20glad%20you%27re%20evil%20too.m4a",
		},
	}

	app.GenerateStaticWebsite("./staticsite", handler)
	compressed := handlers.CompressHandlerLevel(handler, gzip.BestSpeed)

	// Create spyware module
	spywareapi, err := api.NewApiApp()
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", compressed)
	http.HandleFunc("/api/visit", spywareapi.Visit)

	//	router.HandleFunc("/api/visit", spywareapi.Visit)
	if os.Getenv("GEN_STATIC_SITE") == "true" {
		return
	}

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
