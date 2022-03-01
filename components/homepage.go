package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Homepage struct {
	app.Compo

	content *contentView
}

func NewHomepage() *Homepage {
	p1 := newHomePanel()
	c := newContentView(p1)
	return &Homepage{content: c}
}

func (p *Homepage) Render() app.UI {
	return app.Div().Body(
		&header{},
		app.Div().Body(
			newNavbar(),
			newHomePanel(),
			newHomePanel(),
			newHomePanel(),
		).Class("main"),
	)
}
