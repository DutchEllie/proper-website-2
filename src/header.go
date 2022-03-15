package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type header struct {
	app.Compo
}

func (h *header) Render() app.UI {
	return app.Div().Text("Internetica Galactica").Class("header")
}
