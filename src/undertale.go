package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type UndertalePage struct {
	app.Compo
}

// TODO: Autoplay Megalovania

func NewUndertalePage() *UndertalePage {
	return &UndertalePage{}
}

func (u *UndertalePage) Render() app.UI {
	return newPage().
		Title("Undertale").
		LeftBar(
			newHTMLBlock().
				Class("left").
				Class("leftbarblock a").
				Src("/web/blocks/snippets/bannerpanel.html"),
		).
		Main(
			newHTMLBlock().
				Class("right").
				Class("contentblock").
				Src("/web/blocks/pages/undertale.html"),
		)
}
