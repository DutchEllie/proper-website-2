package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type menu struct {
	app.Compo

	Iclass          string
	updateAvailable bool

	IpaneWidth    int
	OnClickButton func(page string)
}

func newMenu() *menu {
	return &menu{}
}

func (m *menu) PaneWidth(px int) *menu {
	if px > 0 {
		m.IpaneWidth = px
	}
	return m
}

func (m *menu) Class(v string) *menu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) OnAppUpdate(ctx app.Context) {
	m.updateAvailable = ctx.AppUpdateAvailable()
}

func (m *menu) Render() app.UI {
	return app.Div().
		Class("block").
		//		Class("leftbarblock-nop").
		Class("navbar").
		Body(
			app.Ul().Body(
				newMenuLink().
					Link("/").
					Text("Home"),
				newMenuLink().
					Link("/about").
					Text("About"),
				newMenuLink().
					Link("/galaxies").
					Text("Galaxies"),
				newMenuLink().
					Link("/music").
					Text("Music"),
				// Disabled for now since there are none anyway
				app.Li().
					Body(
						app.A().Href("/blog").Text("Blog"),
					).Style("display", "none"),
			),
			app.If(m.updateAvailable,
				app.Div().Body(
					app.Img().
						Src("/web/static/images/hot1.gif").
						Class("update-img"),
					app.Span().
						Text("Update available! Click here to update!").
						Class("update-text"),
				).
					Class("update-div").
					OnClick(m.onUpdateClick),
			),
		)
}

func (m *menu) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}

type menuLink struct {
	app.Compo

	IText string
	ILink string
}

func newMenuLink() *menuLink {
	return &menuLink{}
}

func (m *menuLink) Text(v string) *menuLink {
	m.IText = v
	return m
}

func (m *menuLink) Link(v string) *menuLink {
	m.ILink = v
	return m
}

func (m *menuLink) Render() app.UI {
	return app.A().
		Class("menuitem-link").
		Href(m.ILink).
		Body(app.Div().
			Class("menuitem").
			Body(app.Span().
				Class("menuitem-text").
				Text(m.IText)),
		)
}
