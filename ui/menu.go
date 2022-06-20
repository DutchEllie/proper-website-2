package ui

import (
	"strconv"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IMenu interface {
	app.UI

	ID(v string) IMenu
	Class(v string) IMenu
	PaneWidth(px int) IMenu
	HamburgerButton(v app.UI) IMenu
	HamburgerMenu(v ...app.UI) IMenu
	Menu(v ...app.UI) IMenu
}

func Menu() IMenu {
	return &menu{}
}

type menu struct {
	app.Compo

	Iid              string
	Iclass           string
	Ipanewidth       int
	IhamburgerButton app.UI
	IhamburgerMenu   []app.UI
	Imenu            []app.UI

	hideMenu          bool
	showHamburgerMenu bool
	width             int
}

func (m *menu) ID(v string) IMenu {
	m.Iid = v
	return m
}

func (m *menu) Class(v string) IMenu {
	m.Iclass = app.AppendClass(m.Iclass, v)
	return m
}

func (m *menu) PaneWidth(px int) IMenu {
	if px > 0 {
		m.Ipanewidth = px
	}
	return m
}

func (m *menu) HamburgerButton(v app.UI) IMenu {
	b := app.FilterUIElems(v)
	if len(b) != 0 {
		m.IhamburgerButton = b[0]
	}
	return m
}

func (m *menu) HamburgerMenu(v ...app.UI) IMenu {
	m.IhamburgerMenu = app.FilterUIElems(v...)
	return m
}

func (m *menu) Menu(v ...app.UI) IMenu {
	m.Imenu = app.FilterUIElems(v...)
	return m
}

func (m *menu) OnPreRender(ctx app.Context) {
	m.refresh(ctx)
}

func (m *menu) OnMount(ctx app.Context) {
	m.refresh(ctx)
}

func (m *menu) OnResize(ctx app.Context) {
	m.refresh(ctx)
}

func (m *menu) OnUpdate(ctx app.Context) {
	m.refresh(ctx)
}

func (m *menu) Render() app.UI {
	visible := func(v bool) string {
		if v {
			return "block"
		}
		return "none"
	}

	return app.Div().
		ID(m.Iid).
		Class(m.Iclass).
		Body(
			app.Div().
				//Style("display", "flex").
				Style("width", "100%").
				Style("height", "100%").
				Style("overflow", "hidden").
				Body(
					app.Div().
						Style("position", "relative").
						Style("display", visible(!m.hideMenu)).
						Style("flex-shrink", "0").
						Style("flex-basis", pxToString(m.Ipanewidth)).
						Style("overflow", "hidden").
						Body(m.Imenu...),
				),
			app.Div().
				Style("display", visible(m.hideMenu && len(m.IhamburgerMenu) != 0)).
				Style("position", "absolute").
				Style("top", "0").
				Style("left", "0").
				Style("cursor", "pointer").
				OnClick(m.onHamburgerButtonClick).
				Body(
					app.If(m.IhamburgerButton == nil,
						app.Div().
							Class("goapp-shell-hamburger-button-default").
							Text("☰"),
					),
				),
			app.Div().
				Style("display", visible(m.hideMenu && m.showHamburgerMenu)).
				Style("position", "absolute").
				Style("top", "0").
				Style("left", "0").
				Style("width", "100%").
				Style("height", "100%").
				Style("overflow", "hidden").
				OnClick(m.hideHamburgerMenu).
				Body(m.IhamburgerMenu...),
		)
}

func (m *menu) refresh(ctx app.Context) {
	w, _ := app.Window().Size()
	hideMenu := true
	if w >= 914 {
		hideMenu = false
	}

	if hideMenu != m.hideMenu ||
		w != m.width {
		m.hideMenu = hideMenu
		m.width = w

		ctx.Defer(func(app.Context) {
			m.ResizeContent()
		})
	}
}

func pxToString(px int) string {
	return strconv.Itoa(px) + "px"
}

func (m *menu) onHamburgerButtonClick(ctx app.Context, e app.Event) {
	m.showHamburgerMenu = true
}

func (m *menu) hideHamburgerMenu(ctx app.Context, e app.Event) {
	m.showHamburgerMenu = false
}
