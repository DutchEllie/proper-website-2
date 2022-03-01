package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type AboutPage struct {
	app.Compo
}

func NewAboutPage() *AboutPage {
	return &AboutPage{}
}

func (a *AboutPage) Render() app.UI {
	return app.Div().Body(
		&header{},
		app.Div().Body(
			newNavbar(),
			&aboutPanel{},
		).Class("main"),
	)
}

type aboutPanel struct {
	app.Compo
}

func (a *aboutPanel) Render() app.UI {
	return app.Div().Body(
		app.Raw(`<p>I am a 21 year old computer science student, living and studying in The Netherlands.</p>`),
	).Class("content")
}
