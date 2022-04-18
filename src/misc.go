package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

//#################################
//##                            ###
//## Link with status component ###
//##                            ###
//#################################

type linkWithStatus struct {
	app.Compo

	Ilink     string
	IText     string
	ILinkText string
	ILinkBody app.UI
	status    bool
}

func newLinkWithStatus() *linkWithStatus {
	return &linkWithStatus{}
}

func (f *linkWithStatus) Link(s string) *linkWithStatus {
	f.Ilink = s
	return f
}

func (f *linkWithStatus) LinkBody(v app.UI) *linkWithStatus {
	f.ILinkBody = v
	return f
}

func (f *linkWithStatus) LinkText(s string) *linkWithStatus {
	return f.LinkBody(app.Text(s))
}

func (f *linkWithStatus) Text(s string) *linkWithStatus {
	f.IText = s
	return f
}

func (f *linkWithStatus) checkStatus(ctx app.Context) {
	ctx.Async(func() {
		data := struct {
			Url string `json:"url"`
		}{
			Url: f.Ilink,
		}
		jsondata, err := json.Marshal(data)
		if err != nil {
			app.Log(err)
			return
		}

		req, err := http.NewRequest("POST", ApiURL+"/checkonline", bytes.NewBuffer(jsondata))
		if err != nil {
			app.Log(err)
			return
		}
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			app.Log(err)
			return
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			app.Log(err)
			return
		}

		jsonresp := struct {
			Status bool `json:"status"`
		}{}
		err = json.Unmarshal(body, &jsonresp)
		if err != nil {
			app.Log(err)
			return
		}

		ctx.Dispatch(func(ctx app.Context) {
			f.status = jsonresp.Status
		})
	})
}

func (f *linkWithStatus) OnNav(ctx app.Context) {
	f.checkStatus(ctx)
}

func (f *linkWithStatus) Render() app.UI {
	return app.Div().
		Body(
			app.Div().
				Style("display", "flex").
				Style("gap", "20px").
				Body(
					app.A().
						Href(f.Ilink).
						Class("p-h3").
						Class("m-t5").
						Body(f.ILinkBody),
					app.If(f.status,
						app.P().
							Style("color", "green").
							Style("width", "fit-content").
							Style("margin", "5px 0px 0px 0px").
							Text("Online"),
					).Else(
						app.P().
							Style("color", "red").
							Style("width", "fit-content").
							Style("margin", "5px 0px 0px 0px").
							Text("Offline"),
					),
				),
			app.P().
				Class("m-t5").
				Text(f.IText),
		)
}

//#################################
//##                            ###
//## Text with tooltip          ###
//##                            ###
//#################################

type textWithTooltip struct {
	app.Compo

	ITooltip   app.UI
	IClass     string
	ITextClass string
	IText      string

	activated      bool
	mousex, mousey int
}

func newTextWithTooltip() *textWithTooltip {
	return &textWithTooltip{}
}

func (f *textWithTooltip) Class(v string) *textWithTooltip {
	f.IClass = app.AppendClass(f.IClass, v)
	return f
}

func (f *textWithTooltip) TextClass(v string) *textWithTooltip {
	f.ITextClass = app.AppendClass(f.ITextClass, v)
	return f
}

func (f *textWithTooltip) Text(v string) *textWithTooltip {
	f.IText = v
	return f
}

func (f *textWithTooltip) Tooltip(v app.UI) *textWithTooltip {
	f.ITooltip = v
	return f
}

func (f *textWithTooltip) Render() app.UI {
	return app.Div().
		Class(f.IClass).
		Body(
			app.Span().
				Class(f.ITextClass).
				Text(f.IText).
				OnMouseOver(func(ctx app.Context, e app.Event) {
					f.activated = true
				}).
				OnMouseMove(func(ctx app.Context, e app.Event) {
					f.mousex, f.mousey = app.Window().CursorPosition()
				}).
				OnMouseOut(func(ctx app.Context, e app.Event) {
					f.activated = false
				}),
			app.If(f.activated,
				app.Div().
					Style("position", "fixed").
					Style("overflow", "hidden").
					Style("top", fmt.Sprintf("%dpx", f.mousey+20)).
					Style("left", fmt.Sprintf("%dpx", f.mousex+20)).
					Body(
						f.ITooltip,
					),
			),
		)
}

//#################################
//##                            ###
//## Tooltip                    ###
//##                            ###
//#################################

type toolTip struct {
	app.Compo
}
