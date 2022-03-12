package components

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
		app.Raw(`<p>This website is my creative outlet and a way of expressing myself.
As of now, it's probably the most impressive thing I've ever coded.
<br><br>
Please enjoy yourself and do sign the guestbook!!!</p>`),
		app.If(p.updateAvailable,
			app.Div().Body(
				app.P().Text("An update is available! Reload to update!!"),
			)),
		app.Div().Body(
			app.P().Text("Please sign my guestbook!").Style("font-size", "0.8em"),
			app.Img().Src("/web/static/images/email3.gif").Style("width", "40px").Style("position", "absolute").Style("bottom", "0px").Style("right", "0px"),
		).Style("position", "absolute").Style("bottom", "5px").Style("right", "5px").
			OnClick(func(ctx app.Context, e app.Event) {
				e.PreventDefault()
				p.onShowClick()
			}),
	).Class("content")
}

func (p *homePanel) OnAppUpdate(ctx app.Context) {
	p.updateAvailable = true
}
