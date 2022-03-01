package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type homePanel struct {
	app.Compo
}

func newHomePanel() *homePanel {
	return &homePanel{}
}

func (p *homePanel) Render() app.UI {
	return app.Div().Body(
		app.Raw(`<p>Welcome to my website, internet traveler!
This website is my creative outlet and a way of expressing myself.
As of now, it's probably the most impressive thing I've ever coded.
<br>
Please enjoy yourself and do sign the guestbook!!</p>`),
	).Class("content")
}
