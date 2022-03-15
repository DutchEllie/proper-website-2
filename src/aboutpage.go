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
			newHTMLBlock().
				Class("left").
				Class("leftbarblock").
				Src("/web/blocks/snippets/bannerpanel.html"),
		).
		Main(
			newHTMLBlock().
				Class("right").
				Class("contentblock").
				Src("/web/blocks/pages/about.html"),
		)
}
