package components

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
				app.A().Href("/friends").Text("Friends"),
			),
		),
	).Class("navbar")
}
