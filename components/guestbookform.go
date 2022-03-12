package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type guestbookForm struct {
	app.Compo

	name    string
	message string

	OnSubmit func(
		name string,
		message string,
	) // Handler to implement which calls the api
}

func (g *guestbookForm) Render() app.UI {
	return app.Div().Body(
		app.Form().Body(
			app.Input().
				Type("text").
				Name("name").
				Placeholder("Name").
				Required(true).
				OnChange(g.ValueTo(&g.name)).
				Value(g.name),
			app.Input().
				Type("text").
				Name("message").
				Placeholder("Message").
				Required(true).
				OnChange(g.ValueTo(&g.message)).
				Value(g.message),
			app.Input().
				Type("submit").
				Name("submit"),
		).ID("form").
			OnSubmit(func(ctx app.Context, e app.Event) {
				// This was to prevent the page from reloading
				e.PreventDefault()
				if g.name == "" || g.message == "" {
					fmt.Printf("Error: one or more field(s) are empty. For now unhandled\n")
				}
				g.OnSubmit(g.name, g.message)
				g.clear()
			}),
	).Class("content")
}

func (g *guestbookForm) clear() {
	g.name = ""
	g.message = ""
}
