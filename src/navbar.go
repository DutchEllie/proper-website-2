package main

import (
	"dutchellie.nl/DutchEllie/proper-website-2/ui"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type navbar struct {
	app.Compo
	updateAvailable bool

	OnClickButton func(page string)
}

func (n *navbar) OnAppUpdate(ctx app.Context) {
	n.updateAvailable = ctx.AppUpdateAvailable()
}

func (n *navbar) Render() app.UI {
	return ui.Menu().
		PaneWidth(250).
		Menu(
			newMenu(),
		).
		HamburgerMenu(
			newMenu(),
		)
}

func (n *navbar) onUpdateClick(ctx app.Context, e app.Event) {
	ctx.Reload()
}
