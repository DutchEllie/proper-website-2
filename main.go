package main

import (
	"log"
	"net/http"

	"dutchellie.nl/DutchEllie/proper-website-2/components"
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
	homepage := components.NewHomepage()
	aboutpage := components.NewAboutPage()
	galaxiespage := components.NewGalaxiesPage()
	app.Route("/", homepage)
	app.Route("/about", aboutpage)
	app.Route("/galaxies", galaxiespage)

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
	/*
		uri := "mongodb+srv://guestbook-database:5WUDzpvBKBBiiMCy@cluster0.wtt64.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"

		client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			fmt.Println(err)
			return
		}
		defer func() {
			if err = client.Disconnect(context.TODO()); err != nil {
				panic(err)
			}
		}()

		// Ping the primary
		if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
			panic(err)
		}
		db := client.Database("guestbook")
		coll := db.Collection("comments")

		apiapp := &application{
			client:     client,
			database:   db,
			collection: coll,
		}
	*/
	http.Handle("/", handler)
	//http.HandleFunc("/api/comment", apiapp.Comment)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
