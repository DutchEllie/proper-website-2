package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type FriendsPage struct {
	app.Compo
}

func NewFriendsPage() *FriendsPage {
	return &FriendsPage{}
}

func (f *FriendsPage) Render() app.UI {
	return app.Div().Body(
		&header{},
		&navbar{},
		&friendsPanel{},
	).Class("main")
}

type friendsPanel struct {
	app.Compo
}

func (f *friendsPanel) Render() app.UI {
	return app.Div().Body(
		app.P().
			Text(`My friends!`).
			Class("p-h1"),
		app.Ul().Body(
			app.Li().Body(
				app.IFrame().
					Src("forestofunix.xyz").
					Class("friend-iframe"),
			),
		),
	).Class("content")
}
