package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type navbar struct {
	app.Compo
}

func newNavbar() *navbar {
	return &navbar{}
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
		),
	).Class("navbar")
}
