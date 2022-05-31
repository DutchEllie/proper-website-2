package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type navbar struct {
	app.Compo
	updateAvailable bool

	OnClickButton func(page string)
}

func (n *navbar) OnAppUpdate(ctx app.Context) {
	n.updateAvailable = ctx.AppUpdateAvailable()
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
		app.If(n.updateAvailable,
			app.Div().Body(
				app.Img().
					Src("/web/static/images/hot1.gif").
					Class("update-img"),
				app.Span().
					Text("Update available! Click here to update!").
					Class("update-text"),
			).
				Class("update-div").
				OnClick(n.onUpdateClick),
		),
	).Class("navbar")
}

func (n *navbar) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
