package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

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

	hideRightContent bool
	// TODO: Possibly add "updateavailable" here, so it shows up on every page
}

func newPage() *page {
	return &page{
		hideRightContent: false,
	}
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

func (p *page) OnMount(ctx app.Context) {
	ctx.Handle("right-hide", p.hideRight)
	ctx.Handle("right-show", p.showRight)

	// Send a visit request to the spyware API to track people
	ctx.Async(func() {
		resp, err := client.Get("/api/visit")
		if err != nil {
			app.Logf("Error while creating vist request %s\n", err.Error())
		}
		defer resp.Body.Close()

		c := resp.Cookies()
		fmt.Printf("c: %v\n", c)
	})

}

func (p *page) Render() app.UI {
	if p.IbackgroundClass == "" {
		p.IbackgroundClass = app.AppendClass(p.IbackgroundClass, "background")
	}
	return app.Div().
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
				Style("display", visible(p.hideRightContent)).
				Class("right").
				ID("right").
				Body(
					app.Range(p.Imain).Slice(func(i int) app.UI {
						return p.Imain[i]
					}),
				),
		)
}

func visible(v bool) string {
	if v {
		return "none"
	}
	return "block"
}

func (p *page) hideRight(ctx app.Context, a app.Action) {
	p.hideRightContent = true
}

func (p *page) showRight(ctx app.Context, a app.Action) {
	p.hideRightContent = false
}
