package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type homePanel struct {
	app.Compo

	onShowClick     func()
	updateAvailable bool
}

func newHomePanel() *homePanel {
	return &homePanel{}
}

func (p *homePanel) Render() app.UI {
	return app.Div().Body(
		app.P().Text("Welcome, internet surfer!").Class("p-h1"),
		app.Div().Body(
			app.P().Text("Please sign my guestbook!").Class("small"),
			app.Img().Src("/web/static/images/email3.gif").Style("width", "40px").Style("position", "absolute").Style("bottom", "0px").Style("right", "0px"),
		).Style("position", "absolute").Style("top", "10px").Style("right", "5px").
			OnClick(func(ctx app.Context, e app.Event) {
				e.PreventDefault()
				p.onShowClick()
			}),
		app.Img().
			Style("float", "right").
			Style("margin-bottom", "10px").
			Height(230).
			Src("/web/static/images/rin-len1.webp"),
		app.Raw(
			`
		<p class="content-text">
		Welcome to my webspace! Whether you stumbled across this page by accident
		or were linked here, you're more than welcome! This is my personal project that I like
		to work on! I was inspired by a couple friends of mine, please do check their webspaces
		out as well under "Galaxies" on the left side there!
		If you like this page, there is a lot more, so have a look around! You can also leave a
		nice message for me in the guestbook! There is no registration (unlike the rest of the "modern"
		internet) so nothing of that sort!
		That said, this website is my creative outlet and a way to introduce myself, so be kind please!
		Also its code is entirely open-source and can be found 
		<a href="https://dutchellie.nl/DutchEllie/proper-website-2">here</a> so if you like that sort 
		of stuff, be my guest it's cool!</p>
		`),
		app.If(p.updateAvailable,
			app.Div().Body(
				app.P().
					Class("content-text").
					Text("An update is available! Reload to update!"),
			)),
	).Class("content")
}

func (p *homePanel) OnAppUpdate(ctx app.Context) {
	p.updateAvailable = true
}
