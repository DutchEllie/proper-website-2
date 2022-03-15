package main

import (
	"fmt"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type htmlDoc struct {
	app.Compo

	Ihtml string
}

func newHTMLDoc() *htmlDoc {
	return &htmlDoc{}
}

func (h *htmlDoc) HTML(v string) *htmlDoc {
	h.Ihtml = fmt.Sprintf("<div>%s</div>", v)
	return h
}

func (h *htmlDoc) Render() app.UI {
	return app.Raw(h.Ihtml)
}

type remoteHTMLDoc struct {
	app.Compo

	Isrc string

	html htmlContent
}

func newRemoteHTMLDoc() *remoteHTMLDoc {
	return &remoteHTMLDoc{}
}

func (h *remoteHTMLDoc) Src(v string) *remoteHTMLDoc {
	h.Isrc = v
	return h
}

func (h *remoteHTMLDoc) OnMount(ctx app.Context) {
	h.load(ctx)
}

func (h *remoteHTMLDoc) OnNav(ctx app.Context) {
	h.load(ctx)
}

func (h *remoteHTMLDoc) load(ctx app.Context) {
	src := h.Isrc
	ctx.ObserveState(htmlState(src)).
		While(func() bool {
			return src == h.Isrc
		}).
		OnChange(func() {

		}).
		Value(&h.html)

	ctx.NewAction(getHTML, app.T("path", h.Isrc))
}

func (h *remoteHTMLDoc) Render() app.UI {
	return app.Div().
		Body(
			app.If(h.html.Status == loaded,
				newHTMLDoc().
					HTML(h.html.Data),
			).Else(),
		)
}
