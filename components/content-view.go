package components

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type contentView struct {
	app.Compo

	panels []app.UI
}

func newContentView(panels ...app.UI) *contentView {
	return &contentView{panels: panels}
}

func (c *contentView) Render() app.UI {
	return app.Div().Body(
		app.Range(c.panels).Slice(func(i int) app.UI {
			return c.panels[i]
		}),
	)
}
