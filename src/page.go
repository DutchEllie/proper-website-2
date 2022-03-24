package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

// Page is a generic page. By default it has a header, navbar and a default leftbar
type page struct {
	app.Compo

	Ititle string
	/*Description
	Blah blah
	etc*/

	IbackgroundClass string
	Ibackground      []app.UI
	IleftBar         []app.UI
	Imain            []app.UI

	// TODO: Possibly add "updateavailable" here, so it shows up on every page
}

func newPage() *page {
	return &page{}
}

func (p *page) Background(v ...app.UI) *page {
	p.Ibackground = app.FilterUIElems(v...)
	return p
}

func (p *page) BackgroundClass(t string) *page {
	p.IbackgroundClass = app.AppendClass(p.IbackgroundClass, t)
	return p
}

func (p *page) Title(t string) *page {
	p.Ititle = t
	return p
}

func (p *page) LeftBar(v ...app.UI) *page {
	p.IleftBar = app.FilterUIElems(v...)
	return p
}

func (p *page) Main(v ...app.UI) *page {
	p.Imain = app.FilterUIElems(v...)
	return p
}

func (p *page) Render() app.UI {
	if p.IbackgroundClass == "" {
		p.IbackgroundClass = app.AppendClass(p.IbackgroundClass, "background")
	}
	return app.Div().
		Class(p.IbackgroundClass).
		Body(
			app.Range(p.Ibackground).Slice(func(i int) app.UI {
				return p.Ibackground[i]
			}),
			app.Div().
				Class("main").
				Body(
					// Header and navbar
					&header{},
					app.Div().
						Class("left").
						Body(
							&navbar{},
							app.Range(p.IleftBar).Slice(func(i int) app.UI {
								return p.IleftBar[i]
							}),
						),
					app.Div().
						Class("right").
						Body(
							app.Range(p.Imain).Slice(func(i int) app.UI {
								return p.Imain[i]
							}),
						),
				),
		)
}
