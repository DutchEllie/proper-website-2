package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type header struct {
	app.Compo
}

func (h *header) Render() app.UI {
	return app.Div().
		Class("header").
		Body(
			app.Text("Internetica Galactica"),
			//&updater{},
		)
}
