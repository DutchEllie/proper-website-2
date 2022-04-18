package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type GalaxiesPage struct {
	app.Compo

	gnuOver        bool
	mousex, mousey int
}

func NewGalaxiesPage() *GalaxiesPage {
	return &GalaxiesPage{}
}

func (f *GalaxiesPage) Render() app.UI {
	return newPage().
		Title("Galaxies").
		LeftBar(
			newHTMLBlock().
				Class("left").
				Class("leftbarblock").
				Src("/web/blocks/snippets/bannerpanel.html"),
		).
		Main(
			/*
				newHTMLBlock().
					Class("right").
					Class("contentblock").
					Src("/web/blocks/pages/galaxies.html"),
			*/
			newUIBlock().
				Class("right").
				Class("contentblock").
				UI(
					app.Div().
						Body(
							app.P().
								Class("p-h1").
								Text("Galaxies"),
							app.P().
								Class("content-text").
								Text(`Here you can find some really really really cool pages that I found on the internet.
Some of these are blogs or even blogposts I found, but the ones on top are special!
They're the websites of friends of mine! Please visit them, because they worked really hard
on their websites as well!`),
							app.Div().
								Body(
									app.P().
										Class("p-h2").
										Class("mt-20").
										Class("mb-10").
										Class("bold").
										Text("My friends!"),
									app.Ul().
										Body(
											app.Li().
												Body(
													newLinkWithStatus().
														Link("https://forestofunix.xyz").
														LinkText("Forest of Unix").
														Text("A website made by Sebastiaan. A massive Linux fanboy, runs Gentoo on his ThinkPad. Absolutely based."),
												),
											app.Li().
												Body(
													newLinkWithStatus().
														Link("https://nymphali.neocities.org").
														LinkText("Nymphali").
														Text("The website made by ■■■■■■, whoops Nymphali. They have an awesome minimalist website that's just lovely."),
												),
											app.Li().
												Body(
													newLinkWithStatus().
														Link("https://kristypixel.neocities.org").
														LinkText("Kristypixel").
														Text("Website made by Kristy. Very cute website, I love it! Keep up the awesome work!"),
												),
										),
								),
							app.Div().
								Body(
									app.P().
										Class("p-h2").
										Class("mt-20").
										Class("mb-10").
										Class("bold").
										Text("Neat webspaces"),
									app.P().
										Class("m-t5").
										Style("margin-left", "10px").
										Text("Just very neat websites I found and causes I support. Not necessarily by people I know. I just wanted to share them here!"),
									app.Ul().
										Body(
											app.Li().
												Body(
													newLinkWithStatus().
														Link("https://evillious.ylimegirl.com").
														LinkText("Evillious Chronicles fan guide").
														Text("A VERY cool website made by Ylimegirl! They wrote a whole website dedicated to Evillious Chronicles, which is a super good Japanese light novel and vocaloid series!! Definitely look it up!"),
												),
											app.Li().
												Body(
													app.Div().
														Body(
															app.A().
																Href("https://www.gnu.org").
																Class("p-h3").
																Class("m-t5").
																Text("The GNU Project").
																OnMouseOver(func(ctx app.Context, e app.Event) {
																	f.gnuOver = true
																}).
																OnMouseOut(func(ctx app.Context, e app.Event) {
																	f.gnuOver = false
																}).
																OnMouseMove(func(ctx app.Context, e app.Event) {
																	f.mousex, f.mousey = app.Window().CursorPosition()
																}),
															app.If(f.gnuOver,
																app.Img().
																	Src("/web/static/images/gnu-head-sm.png").
																	Style("position", "fixed").
																	Style("overflow", "hidden").
																	Style("top", fmt.Sprintf("%dpx", f.mousey+20)).
																	Style("left", fmt.Sprintf("%dpx", f.mousex+20)).
																	Width(129).
																	Height(122).
																	Alt("GNU"),
															),
															app.Div().
																Style("display", "flex").
																Style("gap", "5px").
																Body(
																	app.P().
																		Class("m-t5").
																		Text("The official website of the GNU project. They advocate for free/libre software. This is not to be confused with 'open source' software. I highly recommend you read about them and their efforts."),
																),
														),
												),
										),
								),
						),
				),
		)
}

/*
func (f *GalaxiesPage) onMouseOverGnu(ctx app.Context, e app.Event) {

}*/

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
	status    bool
}

func newLinkWithStatus() *linkWithStatus {
	return &linkWithStatus{}
}

func (f *linkWithStatus) Link(s string) *linkWithStatus {
	f.Ilink = s
	return f
}

func (f *linkWithStatus) LinkText(s string) *linkWithStatus {
	f.ILinkText = s
	return f
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
						Text(f.ILinkText),
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
