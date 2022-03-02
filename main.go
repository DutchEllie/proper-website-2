package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"git.home.dutchellie.nl/DutchEllie/proper-website-2/components"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type application struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
}

func main() {
	homepage := components.NewHomepage()
	aboutpage := components.NewAboutPage()
	app.Route("/", homepage)
	app.Route("/about", aboutpage)

	// This is executed on the client side only.
	// It handles client side stuff
	// It exits immediately on the server side
	app.RunWhenOnBrowser()

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

	http.Handle("/", &app.Handler{
		Name:        "Internetica Galactica",
		Description: "A 1990's style PWA!",
		Styles: []string{
			"/web/static/style.css",
			"/web/static/adreena.css",
			"/web/static/anisha.css",
			"/web/static/havakana.css",
		},
		CacheableResources: []string{},
	})
	http.HandleFunc("/api/comment", apiapp.Comment)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
