package components

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type guestbookForm struct {
	app.Compo

	name    string
	message string

	gbModalOpen bool
	OnSubmit    func(
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
					return
				}
				if len(g.name) > 40 || len(g.message) > 360 {
					fmt.Printf("Error: Your message is too long fucker\n")
					g.gbModalOpen = true
					return
				}
				g.OnSubmit(g.name, g.message)
				g.clear()
			}),
		app.If(
			g.gbModalOpen,
			&guestbookAlertModal{
				OnClose: func() {
					g.gbModalOpen = false
					g.Update()
				},
			},
		),
	).Class("content")
}

func (g *guestbookForm) clear() {
	g.name = ""
	g.message = ""
}

type guestbookAlertModal struct {
	app.Compo

	PreviousAttempts int
	OnClose          func() // For when we close the modal
}

func (g *guestbookAlertModal) Render() app.UI {
	return app.Div().
		Class("gb-modal").
		ID("gbModal").
		OnClick(func(ctx app.Context, e app.Event) {
			g.OnClose()
		}).
		Body(
			app.Div().
				Class("gb-modal-content").
				Body(
					app.Span().Class("close").Text("X").
						OnClick(func(ctx app.Context, e app.Event) {
							//modal := app.Window().GetElementByID("gbModal")
							//modal.Set("style", "none")
							g.OnClose()
						}),
					app.P().Text("Your name must be <= 40 and your message must be <= 360 characters"),
				),
		)
}
