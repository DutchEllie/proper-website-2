package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type AboutPage struct {
	app.Compo
}

func NewAboutPage() *AboutPage {
	return &AboutPage{}
}

func (a *AboutPage) Render() app.UI {
	return newPage().
		Title("About me").
		LeftBar(
			&bannerPanel{},
		).
		Main(
			newHTMLBlock().
				Class("right").
				Src("/web/blocks/about.html"),
		)
}
