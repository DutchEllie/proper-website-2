package main

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type updater struct {
	app.Compo

	updateAvailable bool
}

func (u *updater) OnAppUpdate(ctx app.Context) {
	u.updateAvailable = ctx.AppUpdateAvailable()
}

func (u *updater) Render() app.UI {
	return app.Div().Body(
		app.If(u.updateAvailable,
			app.Div().
				Class("update-box").
				Body(
					app.Img().
						Class("pulsing").
						Height(50).
						Src("/web/static/images/hot1.gif"),
					app.P().
						Class("update-message").
						Text("An update is available! Click here to reload!"),
				).
				OnClick(func(ctx app.Context, e app.Event) {
					u.onUpdateClick(ctx, e)
				}),
		),
	)
}

func (u *updater) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
