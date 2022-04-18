package main

import (
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
														Text("A website made by Sebastiaan. A massive Linux fanboy, runs Gentoo on his ThinkPad. Absolutely based. His website is written in Lisp, that's why it's often offline. That was the inspiration for the online/offline status text."),
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
													newLinkWithStatus().
														Link("https://www.gnu.org").
														LinkBody(
															newTextWithTooltip().
																Text("The GNU Project").
																Tooltip(
																	app.Img().
																		Src("/web/static/images/gnu-head-sm.png").
																		Width(129).
																		Height(122).
																		Alt("GNU"),
																),
														).
														Text("The official website of the GNU project. They advocate for free/libre software. This is not to be confused with 'open source' software. I highly recommend you read about them and their efforts."),
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
