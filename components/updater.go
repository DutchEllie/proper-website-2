package components

import "github.com/maxence-charriere/go-app/v9/pkg/app"

type updater struct {
	app.Compo

	updateAvailable bool
}

func (u *updater) onAppUpdate(ctx app.Context) {
	u.updateAvailable = ctx.AppUpdateAvailable()
}

func (u *updater) Render() app.UI {
	return app.Div().Body(
		app.If(u.updateAvailable,
			app.Div().Body(
				app.P().Text("An update for this website is available! Please click here to reload!"),
			).Styles(map[string]string{"position": "absolute", "width": "100px", "bottom": "10px", "right": "10px"}).OnClick(u.onUpdateClick),
		),
	)
}

func (u *updater) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
