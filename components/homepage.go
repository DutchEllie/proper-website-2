package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Homepage struct {
	app.Compo
}

func NewHomepage() *Homepage {
	return &Homepage{}
}

func (p *Homepage) Render() app.UI {
	return app.Div().Body(
		&header{},
		app.Div().Body(
			app.Div().Body(
				app.Ul().Body(
					app.Li().Body(
						app.A().Href("/").Text("Home"),
					),
				),
			).Class("navbar"),
			app.Div().Text("Dit is test tekst voor in de div content").Class("content"),
		).Class("main"),
	)
}
