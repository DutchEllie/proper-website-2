package main

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type htmlBlock struct {
	app.Compo

	Iclass string
	Isrc   string // HTML document source
	Iid    string

	// TODO: implement invisibility for other background functions
}

func newHTMLBlock() *htmlBlock {
	return &htmlBlock{}
}

func (b *htmlBlock) ID(v string) *htmlBlock {
	b.Iid = v
	return b
}

func (b *htmlBlock) Class(v string) *htmlBlock {
	b.Iclass = app.AppendClass(b.Iclass, v)
	return b
}

func (b *htmlBlock) Src(v string) *htmlBlock {
	b.Isrc = v
	return b
}

func (b *htmlBlock) Render() app.UI {
	return app.Div().
		Class("block").
		Class(b.Iclass).
		Body(
			newRemoteHTMLDoc().
				Src(b.Isrc),
		)
}

// ==================
// UI element block
// ==================

type uiBlock struct {
	app.Compo

	Iclass string
	Iui    []app.UI
	Iid    string
}

func newUIBlock() *uiBlock {
	return &uiBlock{}
}

func (b *uiBlock) ID(v string) *uiBlock {
	b.Iid = v
	return b
}

func (b *uiBlock) Class(v string) *uiBlock {
	b.Iclass = app.AppendClass(b.Iclass, v)
	return b
}

func (b *uiBlock) UI(v ...app.UI) *uiBlock {
	b.Iui = app.FilterUIElems(v...)
	return b
}

func (b *uiBlock) Render() app.UI {
	return app.Div().
		Class("block").
		Class(b.Iclass).
		Body(
			app.Range(b.Iui).Slice(func(i int) app.UI {
				return b.Iui[i]
			}),
		)
}
