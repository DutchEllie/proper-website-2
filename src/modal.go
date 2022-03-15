package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// A generic modal to be used on the entire site
type Modal struct {
	app.Compo

	Title string
	Body  []app.UI // Body of the modal

	OnClose func()
}

func (m *Modal) Render() app.UI {
	return app.Div().
		Class("generic-modal").
		ID("genericModal").
		OnClick(func(ctx app.Context, e app.Event) {
			m.OnClose()
		}).
		Body(
			app.Div().
				Class("gb-modal-content").
				Body(
					app.Span().Class("close").Text("X").
						OnClick(func(ctx app.Context, e app.Event) {
							//modal := app.Window().GetElementByID("gbModal")
							//modal.Set("style", "none")
							m.OnClose()
						}),
					app.Div().
						Class("generic-modal-body").
						Body(m.Body...),
				),
		)
}
