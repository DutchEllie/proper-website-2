package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type bannerPanel struct {
	app.Compo
}

func (b *bannerPanel) Render() app.UI {
	return app.Div().
		Class("leftbar").
		Body()
}
