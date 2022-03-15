package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type GalaxiesPage struct {
	app.Compo
}

func NewGalaxiesPage() *GalaxiesPage {
	return &GalaxiesPage{}
}

func (f *GalaxiesPage) Render() app.UI {
	return newPage().
		Title("Galaxies").
		LeftBar(
			&bannerPanel{},
		).
		Main(
			newHTMLBlock().
				Class("right").
				Src("/web/blocks/pages/galaxies.html"),
		)
}
