package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type navbar struct {
	app.Compo

	OnClickButton func(page string)
}

func (n *navbar) Render() app.UI {
	return app.Div().Body(
		app.Ul().Body(
			app.Li().Body(
				app.A().Href("/").Text("Home"),
			),
			app.Li().Body(
				app.A().Href("/about").Text("About"),
			),
			app.Li().Body(
				app.A().Href("/galaxies").Text("Galaxies"),
			),
			app.Li().Body(
				app.A().Href("/music").Text("Music"),
			),
			// Disabled for now since there are none anyway
			app.Li().Body(
				app.A().Href("/blog").Text("Blog"),
			).Style("display", "none"),
		),
	).Class("navbar")
}
